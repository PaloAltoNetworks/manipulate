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
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
)

func TestManipulate_ContextOption(t *testing.T) {

	mctx := NewContext(context.Background())

	Convey("Calling ContextOptionFilter should work", t, func() {
		ContextOptionFilter(elemental.NewFilterComposer().WithKey("a").Equals("b").Done())(mctx.(*mcontext))
		So(mctx.Filter().String(), ShouldEqual, `a == "b"`)
	})

	Convey("Calling ContextOptionNamespace should work", t, func() {
		ContextOptionNamespace("/hello")(mctx.(*mcontext))
		So(mctx.Namespace(), ShouldEqual, `/hello`)
	})

	Convey("Calling ContextOptionRecursive should work", t, func() {
		ContextOptionRecursive(true)(mctx.(*mcontext))
		So(mctx.Recursive(), ShouldEqual, true)
	})

	Convey("Calling ContextOptionOverride should work", t, func() {
		ContextOptionOverride(true)(mctx.(*mcontext))
		So(mctx.Override(), ShouldEqual, true)
	})

	Convey("Calling ContextOptionVersion should work", t, func() {
		ContextOptionVersion(12)(mctx.(*mcontext))
		So(mctx.Version(), ShouldEqual, 12)
	})

	Convey("Calling ContextOptionPage should work", t, func() {
		ContextOptionPage(1, 2)(mctx.(*mcontext))
		So(mctx.Page(), ShouldEqual, 1)
		So(mctx.PageSize(), ShouldEqual, 2)
	})

	Convey("Calling ContextOptionTracking should work", t, func() {
		ContextOptionTracking("a", "b")(mctx.(*mcontext))
		So(mctx.ExternalTrackingID(), ShouldEqual, "a")
		So(mctx.ExternalTrackingType(), ShouldEqual, "b")
	})

	Convey("Calling ContextOptionOrder should work", t, func() {
		ContextOptionOrder("a", "b")(mctx.(*mcontext))
		So(mctx.Order(), ShouldResemble, []string{"a", "b"})
	})

	Convey("Calling ContextOptionParameters should work", t, func() {
		ContextOptionParameters(url.Values{"a": []string{"b"}})(mctx.(*mcontext))
		So(mctx.Parameters(), ShouldResemble, url.Values{"a": []string{"b"}})
	})

	Convey("Calling ContextOptionFinalizer should work", t, func() {
		f := func(elemental.Identifiable) error { return nil }
		ContextOptionFinalizer(f)(mctx.(*mcontext))
		So(mctx.Finalizer(), ShouldEqual, f)
	})

	Convey("Calling ContextOptionFinalizer should work", t, func() {
		tid := NewTransactionID()
		ContextOptionTransationID(tid)(mctx.(*mcontext))
		So(mctx.TransactionID(), ShouldEqual, tid)
	})

	Convey("Calling ContextOptionParent should work", t, func() {
		i := testmodel.NewList()
		ContextOptionParent(i)(mctx.(*mcontext))
		So(mctx.Parent(), ShouldEqual, i)
	})

	Convey("Calling ContextOptionFields should work", t, func() {
		ContextOptionFields([]string{"a", "b"})(mctx.(*mcontext))
		So(mctx.Fields(), ShouldResemble, []string{"a", "b"})
	})

	Convey("Calling ContextOptionWriteConsistency should work", t, func() {
		ContextOptionWriteConsistency(WriteConsistencyStrong)(mctx.(*mcontext))
		So(mctx.WriteConsistency(), ShouldEqual, WriteConsistencyStrong)
	})

	Convey("Calling ContextOptionReadConsistency should work", t, func() {
		ContextOptionReadConsistency(ReadConsistencyStrong)(mctx.(*mcontext))
		So(mctx.ReadConsistency(), ShouldEqual, ReadConsistencyStrong)
	})

	Convey("Calling ContextOptionCredentials should work", t, func() {
		ContextOptionCredentials("username", "password")(mctx.(*mcontext))
		u, p := mctx.Credentials()
		So(u, ShouldResemble, "username")
		So(p, ShouldResemble, "password")
	})

	Convey("Calling ContextOptionToken should work", t, func() {
		ContextOptionToken("token")(mctx.(*mcontext))
		u, p := mctx.Credentials()
		So(u, ShouldResemble, "Bearer")
		So(p, ShouldResemble, "token")
	})

	Convey("Calling ContextOptionClientIP should work", t, func() {
		ContextOptionClientIP("10.1.1.1")(mctx.(*mcontext))
		So(mctx.ClientIP(), ShouldEqual, "10.1.1.1")
	})
}
