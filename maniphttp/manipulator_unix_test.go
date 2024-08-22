//go:build linux
// +build linux

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

package maniphttp

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"syscall"
	"testing"
	"time"

	// nolint:revive // Allow dot imports for readability in tests
	. "github.com/smartystreets/goconvey/convey"
	internalsyscall "go.aporeto.io/manipulate/maniphttp/internal/syscall"
	"golang.org/x/sys/unix"
)

func TestHTTP_TCPUserTimeout(t *testing.T) {
	Convey("When I create a simple manipulator with custom transport, with TCP option", t, func() {
		dialer := (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
			Control:   internalsyscall.MakeDialerControlFunc(30 * time.Second),
		}).DialContext

		transport := &http.Transport{
			DialContext: dialer,
		}
		transport.TLSClientConfig = &tls.Config{}

		mm, _ := New(
			context.Background(),
			"http://url.com/",
			OptionHTTPTransport(transport),
			OptionTCPUserTimeout(40*time.Second),
		)

		m := mm.(*httpManipulator)

		Convey("Then the tls config is correct", func() {
			So(m.tlsConfig, ShouldEqual, transport.TLSClientConfig)
		})
		Convey("Then the dialer is correct", func() {
			l, err := net.Listen("tcp", ":0")
			So(err, ShouldBeNil)

			opt := -1
			dctx := m.client.Transport.(*http.Transport).DialContext
			So(dctx, ShouldNotBeNil)
			conn, err := dctx(context.TODO(), "tcp", l.Addr().String())
			So(err, ShouldBeNil)

			tcpConn, ok := conn.(*net.TCPConn)
			So(ok, ShouldBeTrue)

			rawConn, err := tcpConn.SyscallConn()
			So(err, ShouldBeNil)

			err = rawConn.Control(func(fd uintptr) {
				opt, err = syscall.GetsockoptInt(int(fd), syscall.IPPROTO_TCP, unix.TCP_USER_TIMEOUT)
			})
			So(err, ShouldBeNil)
			So(opt, ShouldEqual, 30*time.Second/time.Millisecond)
			l.Close() // nolint
		})
	})
	Convey("When I create a simple manipulator with default transport, with TCP_USER_TIMEOUT", t, func() {
		mm, _ := New(
			context.Background(),
			"http://url.com/",
			OptionTCPUserTimeout(40*time.Second),
		)

		m := mm.(*httpManipulator)

		Convey("Then the dialer is correct", func() {
			l, err := net.Listen("tcp", ":0")
			So(err, ShouldBeNil)

			opt := -1
			dctx := m.client.Transport.(*http.Transport).DialContext
			So(dctx, ShouldNotBeNil)
			conn, err := dctx(context.TODO(), "tcp", l.Addr().String())
			So(err, ShouldBeNil)

			tcpConn, ok := conn.(*net.TCPConn)
			So(ok, ShouldBeTrue)

			rawConn, err := tcpConn.SyscallConn()
			So(err, ShouldBeNil)

			err = rawConn.Control(func(fd uintptr) {
				opt, err = syscall.GetsockoptInt(int(fd), syscall.IPPROTO_TCP, unix.TCP_USER_TIMEOUT)
			})
			So(err, ShouldBeNil)
			So(opt, ShouldEqual, 40*time.Second/time.Millisecond)

			l.Close() // nolint
		})
	})
	Convey("When I create a simple manipulator with default transport, without TCP_USER_TIMEOUT", t, func() {
		mm, _ := New(
			context.Background(),
			"http://url.com/",
		)

		m := mm.(*httpManipulator)

		Convey("Then the dialer is correct", func() {
			l, err := net.Listen("tcp4", ":0")
			So(err, ShouldBeNil)

			opt := -1
			dctx := m.client.Transport.(*http.Transport).DialContext
			So(dctx, ShouldNotBeNil)
			conn, err := dctx(context.TODO(), "tcp4", l.Addr().String())
			So(err, ShouldBeNil)

			tcpConn, ok := conn.(*net.TCPConn)
			So(ok, ShouldBeTrue)

			rawConn, err := tcpConn.SyscallConn()
			So(err, ShouldBeNil)

			err = rawConn.Control(func(fd uintptr) {
				opt, err = syscall.GetsockoptInt(int(fd), syscall.IPPROTO_TCP, unix.TCP_USER_TIMEOUT)
			})
			So(err, ShouldBeNil)
			So(opt, ShouldEqual, 0)

			l.Close() // nolint
		})
	})
}
