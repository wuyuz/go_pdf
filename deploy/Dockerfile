FROM golang:1.13.6

MAINTAINER Alex "1417506149@qq.com"

WORKDIR $GOPATH/src/github.com/pdfUtil

ADD . $GOPATH/src/github.com/pdfUtil

RUN go build .

EXPOSE 8002

ENTRYPOINT  ["./pdfUtil"]