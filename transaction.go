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

import "github.com/gofrs/uuid"

// TransactionID is the type used to define a transcation ID of a store
type TransactionID string

// NewTransactionID returns a new transaction ID.
func NewTransactionID() TransactionID {

	return TransactionID(uuid.Must(uuid.NewV4()).String())
}
