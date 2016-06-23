// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package maniphttp

import (
	"crypto/tls"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTLSConfiguration_NewTLSConfiguration(t *testing.T) {

	Convey("Given I create a new TLSConfiguration with no certificate info", t, func() {

		Convey("Then it should not panic", func() {
			So(func() { NewTLSConfiguration("", "", true) }, ShouldNotPanic)
		})
	})

	Convey("Given I create a new TLSConfiguration with certificate info", t, func() {

		Convey("Then it should not panic", func() {
			So(func() { NewTLSConfiguration("1", "password", true) }, ShouldNotPanic)
		})
	})
}

func TestTLSConfiguration_makeHTTPClient(t *testing.T) {

	Convey("Given I create a new TLSConfiguration with no p12 info", t, func() {

		config := NewTLSConfiguration("", "", true)

		Convey("When I make a http client", func() {
			client, err := config.makeHTTPClient()

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the client should not have any tls information", func() {
				So(client.Transport, ShouldResemble, &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}})
			})
		})
	})

	Convey("Given I create a new TLSConfiguration with bad p12 path", t, func() {

		config := NewTLSConfiguration("bad.p12", "", true)

		Convey("When I make a http client", func() {
			_, err := config.makeHTTPClient()

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I create a new TLSConfiguration with bad p12 content", t, func() {

		config := NewTLSConfiguration("./Makefile", "", true)

		Convey("When I make a http client", func() {
			_, err := config.makeHTTPClient()

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I create a new TLSConfiguration with all good p12 and password", t, func() {

		config := NewTLSConfiguration("fixtures/cert.p12", "password", true)

		Convey("When I make a http client", func() {
			_, err := config.makeHTTPClient()

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I create a new TLSConfiguration with all good p12 info and bad password", t, func() {

		config := NewTLSConfiguration("fixtures/cert.p12", "bad", true)

		Convey("When I make a http client", func() {
			_, err := config.makeHTTPClient()

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
