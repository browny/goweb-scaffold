FROM golang:1.6.2

ADD . /go/src/goweb-scaffold
WORKDIR /go/src/goweb-scaffold
RUN go get github.com/tools/godep
RUN make restore-deps
RUN go install goweb-scaffold

ENTRYPOINT ["/go/bin/goweb-scaffold"]
