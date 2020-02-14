# create API image

# base image
FROM golang:1.13.1-alpine

# `dep ensure` excute directory
WORKDIR /go/src/api/

# Copy Gopkg.toml in advance so that `dev ensure` can be executed
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only=true
# package install
RUN set -x \
  && apk update \
  && apk add --no-cache git \
  # Web Framework
  && go get -u github.com/codegangsta/gin \
  # Package Manager
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure -v -vendor-only=true \
  # Hot Reload
  && go get -u github.com/oxequa/realize \
  # Debugger
  && go get -u github.com/derekparker/delve/cmd/dlv \
  && go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv \
  # create entry point
  && go build -gcflags "-N -l" -o /server main.go

# exec
CMD ["./server"]
