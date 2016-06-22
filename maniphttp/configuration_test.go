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
			So(func() { NewTLSConfiguration("", "", "", true) }, ShouldNotPanic)
		})
	})

	Convey("Given I create a new TLSConfiguration with certificate info", t, func() {

		Convey("Then it should not panic", func() {
			So(func() { NewTLSConfiguration("1", "2", "3", true) }, ShouldNotPanic)
		})
	})

	Convey("When I create a new TLSConfiguration with some missing certificate info", t, func() {

		Convey("Then it should panic if only cacert is given", func() {
			So(func() { NewTLSConfiguration("1", "", "", true) }, ShouldPanic)
		})

		Convey("Then it should panic if only cert is given", func() {
			So(func() { NewTLSConfiguration("", "1", "", true) }, ShouldPanic)
		})

		Convey("Then it should panic if only key is given", func() {
			So(func() { NewTLSConfiguration("", "", "1", true) }, ShouldPanic)
		})

		Convey("Then it should panic if only cacert and cert are given", func() {
			So(func() { NewTLSConfiguration("1", "1", "", true) }, ShouldPanic)
		})

		Convey("Then it should panic if only cacert and key are given", func() {
			So(func() { NewTLSConfiguration("1", "", "1", true) }, ShouldPanic)
		})

		Convey("Then it should panic if only cert and key are given", func() {
			So(func() { NewTLSConfiguration("", "1", "1", true) }, ShouldPanic)
		})
	})

}

func TestTLSConfiguration_makeHTTPClient(t *testing.T) {

	Convey("Given I create a new TLSConfiguration with no certificate info", t, func() {

		config := NewTLSConfiguration("", "", "", true)

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

	Convey("Given I create a new TLSConfiguration with all bad certificate info", t, func() {

		config := NewTLSConfiguration("nothing.pem", "nothing.pem", "nothing.pem", true)

		Convey("When I make a http client", func() {
			_, err := config.makeHTTPClient()

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I create a new TLSConfiguration with certificate and bad CA info", t, func() {

		config := NewTLSConfiguration("bad.pem", "fixtures/cert.pem", "fixtures/key.pem", true)

		Convey("When I make a http client", func() {
			_, err := config.makeHTTPClient()

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I create a new TLSConfiguration with all good certificate info", t, func() {

		config := NewTLSConfiguration("fixtures/ca.pem", "fixtures/cert.pem", "fixtures/key.pem", true)

		Convey("When I make a http client", func() {
			_, err := config.makeHTTPClient()

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
