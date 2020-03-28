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
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"go.aporeto.io/elemental"
)

func Test_makeURL(t *testing.T) {
	type args struct {
		u             string
		namespace     string
		password      string
		recursive     bool
		supportErrors bool
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
				false,
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
				false,
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
				false,
			},
			"wss://toto?namespace=%2Fns",
		},
		{
			"with support errors",
			args{
				"https://toto",
				"/ns",
				"",
				false,
				true,
			},
			"wss://toto?namespace=%2Fns&enableErrors=true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeURL(tt.args.u, tt.args.namespace, tt.args.password, tt.args.recursive, tt.args.supportErrors); got != tt.want {
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

type brokenReader struct{}

func (r *brokenReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("boom")
}

func Test_decodeErrors(t *testing.T) {

	err1, err := elemental.Encode(elemental.EncodingTypeMSGPACK, []error{elemental.NewError("name", "desc", "subj", 3)})
	if err != nil {
		panic(err)
	}

	err2, err := elemental.Encode(elemental.EncodingTypeJSON, []error{elemental.NewError("name", "desc", "subj", 3)})
	if err != nil {
		panic(err)
	}

	type args struct {
		r        io.Reader
		encoding elemental.EncodingType
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{
			"msgpack with good content type",
			args{
				bytes.NewBuffer(err1),
				elemental.EncodingTypeMSGPACK,
			},
			`error 3 (subj): name: desc`,
		},
		{
			"msgpack with bad content type",
			args{
				bytes.NewBuffer(err1),
				elemental.EncodingTypeJSON,
			},
			`Unable to unmarshal data: unable to decode application/json: json decode error [pos 1]: only encoded map or array can be decoded into a slice (0):`,
		},
		{
			"json with good content type",
			args{
				bytes.NewBuffer(err2),
				elemental.EncodingTypeJSON,
			},
			`error 3 (subj): name: desc`,
		},
		{
			"json with bad content type",
			args{
				bytes.NewBuffer(err2),
				elemental.EncodingTypeMSGPACK,
			},
			`Unable to unmarshal data: unable to decode application/msgpack: msgpack decode error [pos 1]: only encoded map or array can be decoded into a slice (0): [{"code":3,"description":"desc","subject":"subj","title":"name"}]`,
		},
		{
			"broken buffer",
			args{
				&brokenReader{},
				elemental.EncodingTypeMSGPACK,
			},
			`Unable to unmarshal data: boom:`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := decodeErrors(tt.args.r, tt.args.encoding); !strings.HasPrefix(err.Error(), tt.wantErr) {
				t.Errorf("decodeErrors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
