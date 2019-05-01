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
	"io/ioutil"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate"
)

type testDataType struct {
	Date     time.Time
	Object   testmodel.List
	Method   elemental.Operation
	Deadline time.Time
}

func Test_newLogWriter(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	Convey("When I create a new log writer with a bad file, I should get an error", t, func() {
		_, err := newLogWriter(ctx, "", 100)
		So(err, ShouldNotBeNil)
	})

	Convey("When I create a new log writer with a good file", t, func() {
		c, err := newLogWriter(ctx, "test.log", 100)
		defer os.Remove("test.log") // nolint errcheck

		Convey("There should be no error and a valid channel", func() {
			So(err, ShouldBeNil)
			So(c, ShouldNotBeNil)
			So(cap(c), ShouldEqual, 100)
		})

		Convey("When I send an event in the channel, the data should be in the file", func() {
			now := time.Now()

			object := &testmodel.List{
				ID:   "1",
				Name: "Object",
			}

			e := &Transaction{
				Date:     now,
				mctx:     manipulate.NewContext(ctx),
				Object:   object,
				Method:   elemental.OperationCreate,
				Deadline: now.Add(10 * time.Second),
			}
			c <- e
			time.Sleep(500 * time.Millisecond)

			data, err := ioutil.ReadFile("test.log")
			So(err, ShouldBeNil)

			model := &testDataType{}

			err = json.Unmarshal(data, &model)
			So(err, ShouldBeNil)
			So(model.Method, ShouldResemble, e.Method)
			So(model.Object.ID, ShouldEqual, "1")
			So(model.Object.Name, ShouldEqual, "Object")
		})
	})
}
