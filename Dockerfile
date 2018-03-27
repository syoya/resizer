FROM golang:1.10.0-alpine3.7

WORKDIR /go/src/github.com/syoya/resizer

COPY . .

RUN apk --update add \
      bash \
      git && \
      go get -u github.com/golang/dep/... && \
      go install .

CMD resizer
