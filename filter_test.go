// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package manipulate

import (
	"reflect"
	"testing"
	"time"

	"go.aporeto.io/elemental"
)

func TestNewNamespaceFilter(t *testing.T) {
	type args struct {
		namespace string
		recursive bool
	}
	tests := []struct {
		name string
		args args
		want *elemental.Filter
	}{
		{
			"non recursive non root",
			args{
				"/a/b",
				false,
			},
			elemental.NewFilterComposer().WithKey("namespace").Equals("/a/b").Done(),
		},
		{
			"recursive non root",
			args{
				"/a/b",
				true,
			},
			elemental.NewFilterComposer().Or(
				elemental.NewFilterComposer().WithKey("namespace").Equals("/a/b").Done(),
				elemental.NewFilterComposer().WithKey("namespace").Matches("^/a/b/").Done(),
			).Done(),
		},
		{
			"non recursive root",
			args{
				"/",
				false,
			},
			elemental.NewFilterComposer().WithKey("namespace").Equals("/").Done(),
		},
		{
			"non recursive empty",
			args{
				"",
				false,
			},
			elemental.NewFilterComposer().WithKey("namespace").Equals("/").Done(),
		},
		{
			"recursive empty",
			args{
				"",
				true,
			},
			elemental.NewFilterComposer().WithKey("namespace").Matches("^/").Done(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNamespaceFilter(tt.args.namespace, tt.args.recursive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NamespaceFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFiltersFromQueryParameters(t *testing.T) {
	type args struct {
		parameters elemental.Parameters
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1      []*elemental.Filter
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		{
			"nil parameters",
			func(t *testing.T) args {
				return args{
					parameters: nil,
				}
			},
			[]*elemental.Filter{},
			false,
			nil,
		},
		{
			"empty parameters",
			func(t *testing.T) args {
				return args{
					parameters: elemental.Parameters{},
				}
			},
			[]*elemental.Filter{},
			false,
			nil,
		},
		{
			"unrelated parameters",
			func(t *testing.T) args {
				return args{
					parameters: elemental.Parameters{
						"not-q": elemental.NewParameter(elemental.ParameterTypeDuration, 10*time.Second),
					},
				}
			},
			[]*elemental.Filter{},
			false,
			nil,
		},
		{
			"invalid parameter type",
			func(t *testing.T) args {
				return args{
					parameters: elemental.Parameters{
						"q": elemental.NewParameter(elemental.ParameterTypeDuration, 10*time.Second),
					},
				}
			},
			[]*elemental.Filter{},
			false,
			nil,
		},
		{
			"simple valid parameters",
			func(t *testing.T) args {
				return args{
					parameters: elemental.Parameters{
						"q": elemental.NewParameter(elemental.ParameterTypeString, "prop == value"),
					},
				}
			},
			[]*elemental.Filter{
				elemental.NewFilterComposer().WithKey("prop").Equals("value").Done(),
			},
			false,
			nil,
		},
		{
			"simple valid multiple parameter",
			func(t *testing.T) args {
				return args{
					parameters: elemental.Parameters{
						"q": elemental.NewParameter(elemental.ParameterTypeString, "prop == value", "prop2 == value2"),
					},
				}
			},
			[]*elemental.Filter{
				elemental.NewFilterComposer().WithKey("prop").Equals("value").Done(),
				elemental.NewFilterComposer().WithKey("prop2").Equals("value2").Done(),
			},
			false,
			nil,
		},
		{
			"invalid filter in parameters",
			func(t *testing.T) args {
				return args{
					parameters: elemental.Parameters{
						"q": elemental.NewParameter(elemental.ParameterTypeString, "prop = oh no! this is not a filter"),
					},
				}
			},
			nil,
			true,
			func(err error, t *testing.T) {
				if err.Error() != "unable to parse filter in query parameter: invalid operator. found = instead of (==, !=, <, <=, >, >=, contains, in, matches, exists)" {
					t.Fail()
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := NewFiltersFromQueryParameters(tArgs.parameters)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewFiltersFromQueryParameters got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("NewFiltersFromQueryParameters error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestNewPropagationFilterWithCustomProperty(t *testing.T) {
	type args struct {
		propagationPropName string
		namespacePropName   string
		namespace           string
		additionalFiltering *elemental.Filter
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 *elemental.Filter
	}{
		{
			"root namespace",
			func(t *testing.T) args {
				return args{
					"propagate",
					"namespace",
					"/",
					nil,
				}
			},
			nil,
		},
		{
			"empty namespace",
			func(t *testing.T) args {
				return args{
					"propagate",
					"namespace",
					"",
					nil,
				}
			},
			nil,
		},
		{
			"first level namespace",
			func(t *testing.T) args {
				return args{
					"propagate",
					"namespace",
					"/level1",
					nil,
				}
			},
			elemental.NewFilterComposer().WithKey("namespace").Equals("/").WithKey("propagate").Equals(true).Done(),
		},
		{
			"second level namespace",
			func(t *testing.T) args {
				return args{
					"propagate",
					"namespace",
					"/level1/level2",
					nil,
				}
			},
			elemental.NewFilterComposer().Or(
				elemental.NewFilterComposer().WithKey("namespace").Equals("/level1").WithKey("propagate").Equals(true).Done(),
				elemental.NewFilterComposer().WithKey("namespace").Equals("/").WithKey("propagate").Equals(true).Done(),
			).Done(),
		},
		{
			"custom properties",
			func(t *testing.T) args {
				return args{
					"p1",
					"p2",
					"/level1",
					nil,
				}
			},
			elemental.NewFilterComposer().WithKey("p2").Equals("/").WithKey("p1").Equals(true).Done(),
		},
		{
			"additional filters",
			func(t *testing.T) args {
				return args{
					"propagate",
					"namespace",
					"/level1/level2",
					elemental.NewFilterComposer().WithKey("x").Equals(true).Done(),
				}
			},
			elemental.NewFilterComposer().Or(
				elemental.NewFilterComposer().
					WithKey("namespace").Equals("/level1").
					WithKey("propagate").Equals(true).
					And(
						elemental.NewFilterComposer().WithKey("x").Equals(true).Done(),
					).
					Done(),
				elemental.NewFilterComposer().
					WithKey("namespace").Equals("/").
					WithKey("propagate").Equals(true).
					And(
						elemental.NewFilterComposer().WithKey("x").Equals(true).Done(),
					).
					Done(),
			).Done(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := NewPropagationFilterWithCustomProperty(tArgs.propagationPropName, tArgs.namespacePropName, tArgs.namespace, tArgs.additionalFiltering)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewPropagationFilter got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
