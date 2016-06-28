FROM golang:1.6

MAINTAINER Antoine Mercadal <antoine@aporeto.com>

ARG GITHUB_TOKEN

ADD . /go/src/github.com/aporeto-inc/manipulate

RUN cd /tmp && \
    wget https://github.com/Masterminds/glide/releases/download/0.10.2/glide-0.10.2-linux-amd64.tar.gz > /dev/null 2>&1 && \
    tar -xzf glide-*.tar.gz && \
    cp linux-amd64/glide /go/bin/ && rm -rf linux-amd64 glide-*-.tar.gz && \
    git config --global url."https://$GITHUB_TOKEN:x-oauth-basic@github.com/".insteadOf "https://github.com/"

WORKDIR /go/src/github.com/aporeto-inc/manipulate

CMD make apoinit && make test
