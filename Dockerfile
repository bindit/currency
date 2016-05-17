FROM golang:1.4

RUN go get github.com/gorilla/mux

ADD . /go/src/currency
WORKDIR /go/src/currency

EXPOSE 8080