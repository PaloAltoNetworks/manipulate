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

package manipmongo

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"go.aporeto.io/manipulate"
	"go.uber.org/zap"
)

// NewMongoManipulator returns a new TransactionalManipulator backed by MongoDB
func NewMongoManipulator(connectionString string, dbName string, user string, password string, authsource string, poolLimit int, CAPool *x509.CertPool, clientCerts []tls.Certificate) manipulate.TransactionalManipulator {

	fmt.Println("DEPRECATED: manipmongo.NewMongoManipulator is deprecated in favor of manipmongo.New")

	m, err := New(
		connectionString,
		dbName,
		OptionCredentials(user, password, authsource),
		OptionConnectionPoolLimit(poolLimit),
		OptionTLS(&tls.Config{
			RootCAs:      CAPool,
			Certificates: clientCerts,
		}),
	)

	if err != nil {
		zap.L().Fatal("Unable to connect to mongo",
			zap.String("uri", connectionString),
			zap.String("db", dbName),
			zap.String("username", user),
			zap.Error(err),
		)
	}

	return m
}
