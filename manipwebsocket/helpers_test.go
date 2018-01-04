package manipwebsocket

import (
	"sync"
	"testing"

	"github.com/gorilla/websocket"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHelpers_IsConnected(t *testing.T) {

	Convey("Given I have a connected manipulator", t, func() {

		m := &websocketManipulator{
			connectedLock: &sync.Mutex{},
			renewLock:     &sync.Mutex{},
			wsLock:        &sync.Mutex{},
			connected:     true,
			ws:            &websocket.Conn{},
		}

		Convey("When I call IsConnected", func() {
			connected := IsConnected(m)

			Convey("Then connected should be true", func() {
				So(connected, ShouldBeTrue)
			})
		})
	})

	Convey("Given I have a connected manipulator with no websocket initialized", t, func() {

		m := &websocketManipulator{
			connectedLock: &sync.Mutex{},
			renewLock:     &sync.Mutex{},
			wsLock:        &sync.Mutex{},
			connected:     true,
		}

		Convey("When I call IsConnected", func() {
			connected := IsConnected(m)

			Convey("Then connected should be true", func() {
				So(connected, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have a disconnected manipulator with", t, func() {

		m := &websocketManipulator{
			connectedLock: &sync.Mutex{},
			renewLock:     &sync.Mutex{},
			wsLock:        &sync.Mutex{},
			connected:     false,
		}

		Convey("When I call IsConnected", func() {
			connected := IsConnected(m)

			Convey("Then connected should be true", func() {
				So(connected, ShouldBeFalse)
			})
		})
	})
}
