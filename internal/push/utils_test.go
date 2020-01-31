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

package push

import (
	"testing"
	"time"
)

func Test_makeURL(t *testing.T) {
	type args struct {
		u         string
		namespace string
		password  string
		recursive bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"with password, recursive",
			args{
				"https://toto",
				"/ns",
				"password",
				true,
			},
			"wss://toto?namespace=%2Fns&token=password&mode=all",
		},
		{
			"with password, no recursive",
			args{
				"https://toto",
				"/ns",
				"password",
				false,
			},
			"wss://toto?namespace=%2Fns&token=password",
		},
		{
			"without password, recursive",
			args{
				"https://toto",
				"/ns",
				"",
				true,
			},
			"wss://toto?namespace=%2Fns&mode=all",
		},
		{
			"without password, no recursive",
			args{
				"https://toto",
				"/ns",
				"",
				false,
			},
			"wss://toto?namespace=%2Fns",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeURL(tt.args.u, tt.args.namespace, tt.args.password, tt.args.recursive); got != tt.want {
				t.Errorf("makeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextBackoff(t *testing.T) {
	type args struct {
		try int
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			"try 1",
			args{0},
			0,
		},
		{
			"try 2",
			args{1},
			3 * time.Millisecond,
		},
		{
			"try 3",
			args{3},
			63 * time.Millisecond,
		},
		{
			"try 4",
			args{4},
			255 * time.Millisecond,
		},
		{
			"try 5",
			args{5},
			1023 * time.Millisecond,
		},
		{
			"try 6",
			args{6},
			4095 * time.Millisecond,
		},
		{
			"try 7",
			args{7},
			8000 * time.Millisecond,
		},
		{
			"try 8",
			args{8},
			8000 * time.Millisecond,
		},
		{
			"try 1000",
			args{1000},
			8000 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextBackoff(tt.args.try); got != tt.want {
				t.Errorf("nextBackoff() = %v, want %v", got, tt.want)
			}
		})
	}
}
