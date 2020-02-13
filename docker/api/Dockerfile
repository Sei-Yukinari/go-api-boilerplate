# create API image

# base image
FROM golang:1.13.1-alpine

# Copy Gopkg.toml in advance so that `dev ensure` can be executed
COPY ./src/api /go/src/api/

# `dep ensure` excute directory
WORKDIR /go/src/api/

# package install
RUN set -x \
  && apk update \
  && apk add --no-cache git \
  # Web Framework
  && go get -u github.com/codegangsta/gin \
  # Package Manager
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure \
  # Hot Reload
  && go get -u github.com/oxequa/realize \
  # Debugger
  && go get -u github.com/derekparker/delve/cmd/dlv \
  && go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv \
  # create main.go
  && go build -gcflags "-N -l" -o /server main.go

# exec
CMD ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "exec", "/server"]
