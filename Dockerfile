FROM ubuntu:14.04
MAINTAINER Alexandre Wilhelm <alex@aporeto.com>

ARG     GITHUB_TOKEN

RUN     apt-get update && apt-get install -y software-properties-common
RUN     add-apt-repository ppa:masterminds/glide

# Install curl, make and git glide
RUN     apt-get update && apt-get install -y curl make git glide

RUN     git config --global url."https://$GITHUB_TOKEN:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# Install golang
RUN     curl -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz && \
        tar -xvf go1.6.linux-amd64.tar.gz && \
        mv go /usr/local

RUN    mkdir -p /aporeto/golang/src/github.com/aporeto-inc/manipulable && mkdir /aporeto/golang/bin

ENV    GOPATH   /aporeto/golang/
ENV    GOBIN    /aporeto/golang/bin
ENV    PATH     /aporeto/golang/bin:$PATH
ENV    PATH     /usr/local/go/bin:$PATH

RUN    go get github.com/smartystreets/goconvey/convey
RUN    go get github.com/aporeto-inc/kennebec/
RUN    go get github.com/golang/lint/golint
RUN    go get golang.org/x/tools/cmd/goimports

ADD . /aporeto/golang/src/github.com/aporeto-inc/manipulable

WORKDIR /aporeto/golang/src/github.com/aporeto-inc/manipulable

ENTRYPOINT     ["/aporeto/golang/src/github.com/aporeto-inc/manipulable/run.sh"]
