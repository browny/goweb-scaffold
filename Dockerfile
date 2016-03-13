FROM golang:1.6.0

ADD . /go/src/goweb-scaffold
WORKDIR /go/src/goweb-scaffold
RUN make deps
RUN go install goweb-scaffold

ENTRYPOINT ["/go/bin/goweb-scaffold"]
