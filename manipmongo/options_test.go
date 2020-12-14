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
	"context"
	"crypto/tls"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

type fakeSharder struct{}

func (*fakeSharder) Shard(manipulate.TransactionalManipulator, manipulate.Context, elemental.Identifiable) error {
	return nil
}
func (*fakeSharder) OnShardedWrite(manipulate.TransactionalManipulator, manipulate.Context, elemental.Operation, elemental.Identifiable) error {
	return nil
}
func (*fakeSharder) FilterOne(manipulate.TransactionalManipulator, manipulate.Context, elemental.Identifiable) (bson.D, error) {
	return nil, nil
}
func (*fakeSharder) FilterMany(manipulate.TransactionalManipulator, manipulate.Context, elemental.Identity) (bson.D, error) {
	return nil, nil
}

func Test_newConfig(t *testing.T) {

	Convey("Given call newConfig", t, func() {

		c := newConfig()

		Convey("Then I should get the default config", func() {
			So(c.username, ShouldEqual, "")
			So(c.password, ShouldEqual, "")
			So(c.tlsConfig, ShouldEqual, nil)
			So(c.poolLimit, ShouldEqual, 4096)
			So(c.connectTimeout, ShouldEqual, 10*time.Second)
			So(c.socketTimeout, ShouldEqual, 60*time.Second)
			So(c.readConsistency, ShouldEqual, manipulate.ReadConsistencyDefault)
			So(c.writeConsistency, ShouldEqual, manipulate.WriteConsistencyDefault)
		})
	})
}

func Test_Options(t *testing.T) {

	Convey("Calling OptionCredentials should work", t, func() {
		c := newConfig()
		OptionCredentials("user", "password", "authdb")(c)
		So(c.username, ShouldEqual, "user")
		So(c.password, ShouldEqual, "password")
		So(c.authsource, ShouldEqual, "authdb")
	})

	Convey("Calling OptionTLS should work", t, func() {
		c := newConfig()
		t := &tls.Config{}
		OptionTLS(t)(c)
		So(c.tlsConfig, ShouldEqual, t)
	})

	Convey("Calling OptionConnectionPoolLimit should work", t, func() {
		c := newConfig()
		OptionConnectionPoolLimit(12)(c)
		So(c.poolLimit, ShouldEqual, 12)
	})

	Convey("Calling OptionConnectionTimeout should work", t, func() {
		c := newConfig()
		OptionConnectionTimeout(12 * time.Second)(c)
		So(c.connectTimeout, ShouldEqual, 12*time.Second)
	})

	Convey("Calling OptionSocketTimeout should work", t, func() {
		c := newConfig()
		OptionSocketTimeout(12 * time.Second)(c)
		So(c.socketTimeout, ShouldEqual, 12*time.Second)
	})

	Convey("Calling OptionDefaultReadConsistencyMode should work", t, func() {
		c := newConfig()
		OptionDefaultReadConsistencyMode(manipulate.ReadConsistencyNearest)(c)
		So(c.readConsistency, ShouldEqual, manipulate.ReadConsistencyNearest)
	})

	Convey("Calling OptionDefaultWriteConsistencyMode should work", t, func() {
		c := newConfig()
		OptionDefaultWriteConsistencyMode(manipulate.WriteConsistencyStrong)(c)
		So(c.writeConsistency, ShouldEqual, manipulate.WriteConsistencyStrong)
	})

	Convey("Calling OptionSharder should work", t, func() {
		c := newConfig()
		s := &fakeSharder{}
		OptionSharder(s)(c)
		So(c.sharder, ShouldEqual, s)
	})

	Convey("Calling OptionDefaultRetryFunc should work", t, func() {
		f := func(manipulate.RetryInfo) error { return nil }
		c := newConfig()
		OptionDefaultRetryFunc(f)(c)
		So(c.defaultRetryFunc, ShouldEqual, f)
	})

	Convey("Calling OptionForceReadFilter should work", t, func() {
		f := bson.D{}
		c := newConfig()
		OptionForceReadFilter(f)(c)
		So(c.forcedReadFilter, ShouldResemble, f)
	})

	Convey("Calling OptionAttributeEncrypter should work", t, func() {
		enc, _ := elemental.NewAESAttributeEncrypter("0123456789ABCDEF")
		c := newConfig()
		OptionAttributeEncrypter(enc)(c)
		So(c.attributeEncrypter, ShouldEqual, enc)
	})

	Convey("Calling OptionExplain should work", t, func() {
		m := map[elemental.Identity]map[elemental.Operation]struct{}{}
		c := newConfig()
		OptionExplain(m)(c)
		So(c.explain, ShouldEqual, m)
	})

	Convey("Calling OptionTranslateKeysFromSpec should work", t, func() {
		c := newConfig()
		specs := map[elemental.Identity]map[string]elemental.AttributeSpecification{
			elemental.MakeIdentity("test", "test"): {
				"someAttribute": elemental.AttributeSpecification{},
			},
		}
		OptionTranslateKeysFromSpec(specs)(c)
		So(c.specs, ShouldEqual, specs)
	})

	Convey("Calling OptionTranslateKeysFromSpec should panic if provided nil specs", t, func() {
		c := newConfig()
		So(func() { OptionTranslateKeysFromSpec(nil)(c) }, ShouldPanic)
	})

	Convey("Calling OptionTranslateKeysFromSpec should panic if provided a nil spec for an identity", t, func() {
		c := newConfig()
		specs := map[elemental.Identity]map[string]elemental.AttributeSpecification{
			elemental.MakeIdentity("test", "test"): nil,
		}
		So(func() { OptionTranslateKeysFromSpec(specs)(c) }, ShouldPanic)
	})
}

func Test_ContextOptions(t *testing.T) {

	Convey("Calling ContextOptionUpsert should work", t, func() {
		b := bson.M{
			"$setOnInsert": bson.M{"hello": "world"},
		}
		mctx := manipulate.NewContext(context.Background())
		ContextOptionUpsert(b)(mctx)
		So(mctx.(opaquer).Opaque()[opaqueKeyUpsert], ShouldEqual, b)
	})

	Convey("Calling ContextOptionUpsert with $set should panic", t, func() {
		b := bson.M{"$set": true}
		So(func() { ContextOptionUpsert(b)(nil) }, ShouldPanicWith, "cannot use $set in upsert operations")
	})

	Convey("Calling ContextOptionUpsert with $setOnInsert with _id should panic", t, func() {
		b := bson.M{"$setOnInsert": bson.M{"_id": 1}}
		So(func() { ContextOptionUpsert(b)(nil) }, ShouldPanicWith, "cannot use $setOnInsert on _id in upsert operations")
	})
}
