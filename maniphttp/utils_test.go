package maniphttp

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"
	"syscall"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMethodNewCertificateAndPoolWithLoadX509KeyPairError(t *testing.T) {
	Convey("Given I give wrong a wrong certificate path", t, func() {
		cert, pool, err := NewCertificateAndPool("Resources/coucou", "Resources/reer", "Resources/path")

		Convey("Then I should get an error", func() {
			So(cert, ShouldBeNil)
			So(pool, ShouldBeNil)

			e := &os.PathError{}
			e.Op = "open"
			e.Path = "Resources/coucou"
			e.Err = syscall.Errno(2)

			So(err, ShouldResemble, e)
		})
	})
}

func TestMethodNewCertificateAndPoolWithCAPathError(t *testing.T) {
	Convey("Given I give wrong a wrong ca path", t, func() {
		cert, pool, err := NewCertificateAndPool("Resources/test-cert.pem", "Resources/test-key.pem", "Resources/path")

		Convey("Then I should get an error", func() {
			So(cert, ShouldBeNil)
			So(pool, ShouldBeNil)

			e := &os.PathError{}
			e.Op = "open"
			e.Path = "Resources/path"
			e.Err = syscall.Errno(2)

			So(err, ShouldResemble, e)
		})
	})
}

func TestMethodNewCertificateAndPool(t *testing.T) {
	Convey("Given I give wrong a good path for NewCertificateAndPool", t, func() {
		cert, pool, err := NewCertificateAndPool("Resources/test-cert.pem", "Resources/test-key.pem", "Resources/ca-test.pem")

		c, _ := tls.LoadX509KeyPair("Resources/test-cert.pem", "Resources/test-key.pem")
		caCert, _ := ioutil.ReadFile("Resources/ca-test.pem")

		p := x509.NewCertPool()
		p.AppendCertsFromPEM(caCert)

		Convey("Then I should not get an error", func() {
			So(err, ShouldBeNil)
			So(*cert, ShouldResemble, c)
			So(p, ShouldResemble, pool)

		})
	})
}
