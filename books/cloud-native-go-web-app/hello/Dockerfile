FROM golang:1.15

ADD . /go/src/github.com/oleg/hello

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
ENTRYPOINT /go/bin/hello
