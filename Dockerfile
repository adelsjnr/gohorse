FROM golang:alpine
MAINTAINER jroliv@gmail.com

RUN mkdir -p /go/src /go/bin && chmod -R 777 /go

ENV GOPATH /go
ENV PATH /go/bin:$PATH

WORKDIR /go

ADD . /go

CMD ["go", "run", "main.go"]
