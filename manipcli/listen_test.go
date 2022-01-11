package manipcli

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniphttp"
)

func Test_generateListenCommand(t *testing.T) {

	Convey("Given a server that can listen", t, func() {

		event := &elemental.Event{
			Identity: testmodel.TaskIdentity.Name,
			Type:     elemental.EventCreate,
		}

		data, err := elemental.Encode(elemental.EncodingTypeJSON, event)
		if err != nil {
			fmt.Println("can't encode task", err)
			return
		}

		ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			var upgrader = websocket.Upgrader{}
			ws, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				fmt.Println("cannot upgrade", err)
				return
			}

			go func() {
				defer ws.Close() // nolint
				for {
					fmt.Println("writeMessage...")
					if err := ws.WriteMessage(websocket.TextMessage, data); err != nil {
						fmt.Println("cannot write message", err)
						return
					}
					<-time.After(1 * time.Second)
				}
			}()

		}))
		defer ts.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		tlsConfig := ts.Client().Transport.(*http.Transport).TLSClientConfig
		m, err := maniphttp.New(ctx, ts.URL, maniphttp.OptionTLSConfig(tlsConfig))
		So(err, ShouldEqual, nil)

		cmd, err := generateListenCommand(testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute", func() {

			cmd.Flags().Set(flagAPI, ts.URL)           // nolint
			cmd.Flags().Set(flagNamespace, "/test")    // nolint
			cmd.Flags().Set(flagEncoding, "json")      // nolint
			cmd.Flags().Set(flagToken, "token1234")    // nolint
			cmd.Flags().Set(flagAPISkipVerify, "true") // nolint
			cmd.Flags().Set("identity", "task")        // nolint

			output := bytes.NewBufferString("")
			cmd.SetOut(output)

			execContext, execCancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer execCancel()

			err := cmd.ExecuteContext(execContext)

			Convey("Then I should get a generated command", func() {
				So(err, ShouldEqual, nil)
				So(output.String(), ShouldEqual, `{
  "encoding": "",
  "entity": null,
  "identity": "task",
  "timestamp": null,
  "type": "create"
}`)
			})
		})
	})

	Convey("Given a server that fails", t, func() {

		ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()

		cmd, err := generateListenCommand(testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return nil, fmt.Errorf("boom")
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute", func() {

			cmd.Flags().Set(flagAPI, ts.URL)           // nolint
			cmd.Flags().Set(flagNamespace, "/test")    // nolint
			cmd.Flags().Set(flagEncoding, "json")      // nolint
			cmd.Flags().Set(flagToken, "token1234")    // nolint
			cmd.Flags().Set(flagAPISkipVerify, "true") // nolint
			cmd.Flags().Set("identity", "task")        // nolint

			output := bytes.NewBufferString("")
			cmd.SetOut(output)

			execContext, execCancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer execCancel()

			err := cmd.ExecuteContext(execContext)

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, "unable to make manipulator: boom")
			})
		})
	})

	Convey("Given a command with wrong identity", t, func() {

		m, err := maniphttp.New(context.Background(), "https://blabla.bla")
		So(err, ShouldEqual, nil)

		cmd, err := generateListenCommand(testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute", func() {

			cmd.Flags().Set(flagAPI, "https://mytest.chris") // nolint
			cmd.Flags().Set(flagNamespace, "/test")          // nolint
			cmd.Flags().Set(flagEncoding, "json")            // nolint
			cmd.Flags().Set(flagToken, "token1234")          // nolint
			cmd.Flags().Set(flagAPISkipVerify, "true")       // nolint
			cmd.Flags().Set("identity", "bip")               // nolint

			output := bytes.NewBufferString("")
			cmd.SetOut(output)

			execContext, execCancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer execCancel()

			err := cmd.ExecuteContext(execContext)

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, "unknown identity bip")
			})
		})
	})
}
