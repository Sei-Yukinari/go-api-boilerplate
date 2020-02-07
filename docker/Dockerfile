# API Image 作成

# Base Image
FROM golang:1.13.0-alpine

# move project root
WORKDIR /go/src/app/

# install package
RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/codegangsta/gin \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep init || dep ensure

# load libraly
CMD gin -i run
