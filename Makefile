.PHONY: build deps test clean

build:
	go build ./...

deps:
	go get github.com/mjibson/esc
	go get github.com/cihub/seelog
	go get github.com/codegangsta/negroni
	go get github.com/facebookgo/inject
	go get github.com/gorilla/mux
	go get github.com/robfig/cron

test:
	go test ./...
