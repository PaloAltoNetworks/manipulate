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
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/url"
	"strings"
	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

func decodeErrors(r io.Reader, encoding elemental.EncodingType) error {

	es := []elemental.Error{}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return manipulate.ErrCannotUnmarshal{Err: fmt.Errorf("%w: %s", err, string(data))}
	}

	if err := elemental.Decode(encoding, data, &es); err != nil {
		return manipulate.ErrCannotUnmarshal{Err: fmt.Errorf("%w: %s", err, string(data))}
	}

	errs := elemental.NewErrors()
	for _, e := range es {
		errs = append(errs, e)
	}

	return errs
}

func makeURL(u string, namespace string, password string, recursive, supportErrorEvents bool) string {

	u = strings.Replace(u, "https://", "wss://", 1)

	args := []string{
		fmt.Sprintf("namespace=%s", url.QueryEscape(namespace)),
	}

	if password != "" {
		args = append(args, fmt.Sprintf("token=%s", password))
	}

	if recursive {
		args = append(args, "mode=all")
	}

	if supportErrorEvents {
		args = append(args, "enableErrors=true")
	}

	return fmt.Sprintf("%s?%s", u, strings.Join(args, "&"))
}

const maxBackoff = 8000

func nextBackoff(try int) time.Duration {

	return time.Duration(math.Min(math.Pow(4.0, float64(try))-1, maxBackoff)) * time.Millisecond
}
