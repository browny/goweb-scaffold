FROM golang:1.6.0

RUN go get github.com/tools/godep

ADD . /go/src/goweb-scaffold
WORKDIR /go/src/goweb-scaffold
RUN godep restore goweb-scaffold
RUN go install goweb-scaffold

ENTRYPOINT ["/go/bin/goweb-scaffold"]
