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
	"context"
	"fmt"
)

// Retry is deprecated and only calls manipulateFunc for backward compatibility.
func Retry(ctx context.Context, manipulateFunc func() error, onRetryFunc func(int, error) error) error {
	fmt.Println("DEPRECATED: manipulate.Retry is deprecated. Retry mechanism is now part of Manipulator implementations. You can safely remove this wrapper.")
	return manipulateFunc()
}
