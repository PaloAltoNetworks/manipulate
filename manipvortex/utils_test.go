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

package manipvortex

import (
	"context"
	"testing"

	"go.aporeto.io/manipulate"
)

func Test_isStrongReadConsistency(t *testing.T) {
	type args struct {
		mctx               manipulate.Context
		processor          *Processor
		defaultConsistency manipulate.ReadConsistency
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"mctx default, proc default, default default",
			args{
				manipulate.NewContext(context.Background()),
				&Processor{ReadConsistency: manipulate.ReadConsistencyDefault},
				manipulate.ReadConsistencyDefault,
			},
			false,
		},
		{
			"mctx default, proc default, default strong",
			args{
				manipulate.NewContext(context.Background()),
				&Processor{ReadConsistency: manipulate.ReadConsistencyDefault},
				manipulate.ReadConsistencyStrong,
			},
			true,
		},
		{
			"mctx default, proc strong, default default",
			args{
				manipulate.NewContext(context.Background()),
				&Processor{ReadConsistency: manipulate.ReadConsistencyStrong},
				manipulate.ReadConsistencyDefault,
			},
			true,
		},
		{
			"mctx strong, proc default, default default",
			args{
				manipulate.NewContext(context.Background(), manipulate.ContextOptionReadConsistency(manipulate.ReadConsistencyStrong)),
				&Processor{ReadConsistency: manipulate.ReadConsistencyDefault},
				manipulate.ReadConsistencyDefault,
			},
			true,
		},
		{
			"mctx strong, proc strong, default default",
			args{
				manipulate.NewContext(context.Background(), manipulate.ContextOptionReadConsistency(manipulate.ReadConsistencyStrong)),
				&Processor{ReadConsistency: manipulate.ReadConsistencyStrong},
				manipulate.ReadConsistencyDefault,
			},
			true,
		},
		{
			"mctx strong, proc strong, default strong",
			args{
				manipulate.NewContext(context.Background(), manipulate.ContextOptionReadConsistency(manipulate.ReadConsistencyStrong)),
				&Processor{ReadConsistency: manipulate.ReadConsistencyStrong},
				manipulate.ReadConsistencyStrong,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStrongReadConsistency(tt.args.mctx, tt.args.processor, tt.args.defaultConsistency); got != tt.want {
				t.Errorf("isStrongReadConsistency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStrongWriteConsistency(t *testing.T) {
	type args struct {
		mctx               manipulate.Context
		processor          *Processor
		defaultConsistency manipulate.WriteConsistency
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"mctx default, proc default, default default",
			args{
				manipulate.NewContext(context.Background()),
				&Processor{WriteConsistency: manipulate.WriteConsistencyDefault},
				manipulate.WriteConsistencyDefault,
			},
			false,
		},
		{
			"mctx default, proc default, default strong",
			args{
				manipulate.NewContext(context.Background()),
				&Processor{WriteConsistency: manipulate.WriteConsistencyDefault},
				manipulate.WriteConsistencyStrong,
			},
			true,
		},
		{
			"mctx default, proc strong, default default",
			args{
				manipulate.NewContext(context.Background()),
				&Processor{WriteConsistency: manipulate.WriteConsistencyStrong},
				manipulate.WriteConsistencyDefault,
			},
			true,
		},
		{
			"mctx strong, proc default, default default",
			args{
				manipulate.NewContext(context.Background(), manipulate.ContextOptionWriteConsistency(manipulate.WriteConsistencyStrong)),
				&Processor{WriteConsistency: manipulate.WriteConsistencyDefault},
				manipulate.WriteConsistencyDefault,
			},
			true,
		},
		{
			"mctx strong, proc strong, default default",
			args{
				manipulate.NewContext(context.Background(), manipulate.ContextOptionWriteConsistency(manipulate.WriteConsistencyStrong)),
				&Processor{WriteConsistency: manipulate.WriteConsistencyStrong},
				manipulate.WriteConsistencyDefault,
			},
			true,
		},
		{
			"mctx strong, proc strong, default strong",
			args{
				manipulate.NewContext(context.Background(), manipulate.ContextOptionWriteConsistency(manipulate.WriteConsistencyStrong)),
				&Processor{WriteConsistency: manipulate.WriteConsistencyStrong},
				manipulate.WriteConsistencyStrong,
			},
			true,
		},

		{
			"mctx default, proc default, default default",
			args{
				manipulate.NewContext(context.Background()),
				&Processor{WriteConsistency: manipulate.WriteConsistencyDefault},
				manipulate.WriteConsistencyDefault,
			},
			false,
		},
		{
			"mctx default, proc default, default strong",
			args{
				manipulate.NewContext(context.Background()),
				&Processor{WriteConsistency: manipulate.WriteConsistencyDefault},
				manipulate.WriteConsistencyStrongest,
			},
			true,
		},
		{
			"mctx default, proc strong, default default",
			args{
				manipulate.NewContext(context.Background()),
				&Processor{WriteConsistency: manipulate.WriteConsistencyStrongest},
				manipulate.WriteConsistencyDefault,
			},
			true,
		},
		{
			"mctx strong, proc default, default default",
			args{
				manipulate.NewContext(context.Background(), manipulate.ContextOptionWriteConsistency(manipulate.WriteConsistencyStrongest)),
				&Processor{WriteConsistency: manipulate.WriteConsistencyDefault},
				manipulate.WriteConsistencyDefault,
			},
			true,
		},
		{
			"mctx strong, proc strong, default default",
			args{
				manipulate.NewContext(context.Background(), manipulate.ContextOptionWriteConsistency(manipulate.WriteConsistencyStrongest)),
				&Processor{WriteConsistency: manipulate.WriteConsistencyStrongest},
				manipulate.WriteConsistencyDefault,
			},
			true,
		},
		{
			"mctx strong, proc strong, default strong",
			args{
				manipulate.NewContext(context.Background(), manipulate.ContextOptionWriteConsistency(manipulate.WriteConsistencyStrongest)),
				&Processor{WriteConsistency: manipulate.WriteConsistencyStrongest},
				manipulate.WriteConsistencyStrongest,
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStrongWriteConsistency(tt.args.mctx, tt.args.processor, tt.args.defaultConsistency); got != tt.want {
				t.Errorf("isStrongWriteConsistency() = %v, want %v", got, tt.want)
			}
		})
	}
}
