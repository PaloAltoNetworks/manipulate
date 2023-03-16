package manipbolt

import (
	"reflect"
	"testing"

	"github.com/asdine/storm/q"
	"go.aporeto.io/elemental"
)

func Test_compileFilter(t *testing.T) {
	type args struct {
		f *elemental.Filter
	}
	tests := []struct {
		name    string
		args    args
		want    q.Matcher
		wantErr bool
	}{
		{
			name: "invalid filter",
			args: args{
				elemental.NewFilterComposer().WithKey("pid").NotEquals("5d83e7eedb40280001887565").Done(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty filter",
			args: args{
				elemental.NewFilterComposer().Done(),
			},
			want:    q.And(),
			wantErr: false,
		},
		{
			name: "simple filter",
			args: args{
				elemental.NewFilterComposer().WithKey("pid").Equals("5d83e7eedb40280001887565").Done(),
			},
			want: q.And(
				containsOrEqual("pid", "5d83e7eedb40280001887565"),
			),
			wantErr: false,
		},
		{
			name: "two key filter",
			args: args{
				elemental.NewFilterComposer().WithKey("Name").Equals("Dragon").WithKey("Name").Equals("Eragon").Done(),
			},
			want: q.And(
				containsOrEqual("Name", "Dragon"),
				containsOrEqual("Name", "Eragon"),
			),
			wantErr: false,
		},
		{
			name: "simple regex",
			args: args{
				elemental.NewFilterComposer().
					WithKey("x").Matches("$abc^", ".*").
					Done(),
			},
			want: q.And(
				q.Or(
					regexMatcher("x", "$abc^"),
					regexMatcher("x", ".*"),
				),
			),
			wantErr: false,
		},
		{
			name: "regex with unsupported type",
			args: args{
				elemental.NewFilterComposer().
					WithKey("x").Matches("$abc^", ".*", true).
					Done(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nested and filter",
			args: args{
				elemental.NewFilterComposer().
					WithKey("namespace").Equals("coucou").
					And(
						elemental.NewFilterComposer().
							WithKey("name").Equals("toto").
							WithKey("surname").Equals("titi").
							Done(),
						elemental.NewFilterComposer().
							WithKey("yes").Equals(true).
							Done(),
						elemental.NewFilterComposer().
							WithKey("num").Equals(2).
							Done(),
					).Done(),
			},
			want: q.And(
				containsOrEqual("namespace", "coucou"),
				q.And(
					q.And(
						containsOrEqual("name", "toto"),
						containsOrEqual("surname", "titi"),
					),
					q.And(
						containsOrEqual("yes", true),
					),
					q.And(
						containsOrEqual("num", 2),
					),
				),
			),
			wantErr: false,
		},
		{
			name: "nested or filter",
			args: args{
				elemental.NewFilterComposer().
					WithKey("namespace").Equals("coucou").
					Or(
						elemental.NewFilterComposer().
							WithKey("name").Equals("toto").
							WithKey("surname").Equals("titi").
							Done(),
						elemental.NewFilterComposer().
							WithKey("yes").Equals(true).
							Done(),
						elemental.NewFilterComposer().
							WithKey("num").Equals(2).
							Done(),
					).Done(),
			},
			want: q.And(
				containsOrEqual("namespace", "coucou"),
				q.Or(
					q.And(
						containsOrEqual("name", "toto"),
						containsOrEqual("surname", "titi"),
					),
					q.And(
						containsOrEqual("yes", true),
					),
					q.And(
						containsOrEqual("num", 2),
					),
				),
			),
			wantErr: false,
		},
		{
			name: "nested and/or filter",
			args: args{
				elemental.NewFilterComposer().
					WithKey("namespace").Equals("coucou").
					Or(
						elemental.NewFilterComposer().
							WithKey("name").Equals("toto").
							WithKey("surname").Equals("titi").
							Done(),
					).And(
					elemental.NewFilterComposer().
						WithKey("yes").Equals(true).
						Done(),
					elemental.NewFilterComposer().
						WithKey("num").Equals(2).
						Done(),
				).Done(),
			},
			want: q.And(
				containsOrEqual("namespace", "coucou"),
				q.Or(
					q.And(
						containsOrEqual("name", "toto"),
						containsOrEqual("surname", "titi"),
					),
				),
				q.And(
					q.And(
						containsOrEqual("yes", true),
					),
					q.And(
						containsOrEqual("num", 2),
					),
				),
			),
			wantErr: false,
		},
		{
			name: "complex nested filters with error invalid sub filter",
			args: args{
				elemental.NewFilterComposer().
					WithKey("namespace").Equals("coucou").
					And(
						elemental.NewFilterComposer().
							WithKey("name").Equals("toto").
							WithKey("surname").Equals("titi").
							Done(),
						elemental.NewFilterComposer().
							WithKey("color").Equals("blue").
							Or(
								elemental.NewFilterComposer().
									WithKey("size").Equals("big").
									Done(),
								elemental.NewFilterComposer().
									WithKey("size").Equals("medium").
									Done(),
								elemental.NewFilterComposer().
									WithKey("list").NotIn("a", "b", "c").
									Done(),
							).
							Done(),
					).
					Done(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "complex nested filters",
			args: args{
				elemental.NewFilterComposer().
					WithKey("namespace").Equals("coucou").
					And(
						elemental.NewFilterComposer().
							WithKey("name").Equals("toto").
							WithKey("surname").Equals("titi").
							Done(),
						elemental.NewFilterComposer().
							WithKey("color").Equals("blue").
							WithKey("prefix").Matches("^nope").
							Or(
								elemental.NewFilterComposer().
									WithKey("size").Equals("big").
									Done(),
								elemental.NewFilterComposer().
									WithKey("size").Equals("medium").
									Done(),
								elemental.NewFilterComposer().
									WithKey("list").Contains("a", "b", "c").
									Done(),
							).
							Done(),
					).
					Done(),
			},
			want: q.And(
				containsOrEqual("namespace", "coucou"),
				q.And(
					q.And(
						containsOrEqual("name", "toto"),
						containsOrEqual("surname", "titi"),
					),
					q.And(
						containsOrEqual("color", "blue"),
						q.Or(
							regexMatcher("prefix", "^nope"),
						),
						q.Or(
							q.And(
								containsOrEqual("size", "big"),
							),
							q.And(
								containsOrEqual("size", "medium"),
							),
							q.And(
								q.Or(
									containsOrEqual("list", "a"),
									containsOrEqual("list", "b"),
									containsOrEqual("list", "c"),
								),
							),
						),
					),
				),
			),
			wantErr: false,
		},
		{
			name: "found bug",
			args: args{
				elemental.NewFilterComposer().Or(
					elemental.NewFilterComposer().Or(
						elemental.NewFilterComposer().
							WithKey("namespace").Equals("/").
							WithKey("propagate").Equals(true).Done(),
						elemental.NewFilterComposer().WithKey("namespace").Equals("/apomux").Done(),
					).WithKey("normalizedTags").Contains("network=dns").Done(),
				).WithKey("archived").Equals(false).Done(),
			},
			want: q.And(
				q.Or(
					q.And(
						q.Or(
							q.And(
								containsOrEqual("namespace", "/"),
								containsOrEqual("propagate", true),
							),
							q.And(
								containsOrEqual("namespace", "/apomux"),
							),
						),
						q.Or(
							containsOrEqual("normalizedTags", "network=dns"),
						),
					),
				),
				containsOrEqual("archived", false),
			),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compileFilter(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("compileFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compileFilter() = %#+v, want %#+v", got, tt.want)
			}
		})
	}
}
