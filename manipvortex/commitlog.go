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
	"encoding/json"
	"os"
)

// newLogWriter creates a new log for transactions. At this point there is no
// file rotation. If it runs out of storage it will die.
func newLogWriter(ctx context.Context, filename string, size int) (chan *Transaction, error) {

	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	events := make(chan *Transaction, size)

	go func(f *os.File) {
		// #nosec G307
		defer f.Close() // nolint errcheck

		for {

			select {

			case d := <-events:
				data, err := json.Marshal(d)
				if err != nil {
					continue
				}

				_, err = f.Write(data)
				if err != nil {
					return
				}

			case <-ctx.Done():
				close(events)
				return
			}
		}
	}(f)

	return events, nil

}
