module go.aporeto.io/manipulate

go 1.16

require (
	go.aporeto.io/elemental v1.100.1-0.20211117023454-0c8e5dca0782
	go.aporeto.io/wsc v1.36.1-0.20211206153718-20df727097f4
)

require (
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/golang/mock v1.6.0
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/go-memdb v1.2.1
	github.com/mitchellh/copystructure v1.2.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.6.4
	go.uber.org/zap v1.19.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007
)

replace github.com/gorilla/websocket v1.4.2 => github.com/philipatl/websocket v1.4.3-0.20211206152948-d16969baa130
