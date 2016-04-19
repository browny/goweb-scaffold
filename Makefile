.PHONY: asset build snapshot-deps install deps run test
.DEFAULT_GOAL := help

asset: ## Build asset
	rm -rf config/asset.go
	esc -o config/asset.go -pkg config -ignore=".go" -ignore="DS_Store" config/

build: ## Compile all packages
	go build ./...

snapshot-deps: ## Snapshot dependencies (Godeps)
	rm -rf Godeps/
	rm -rf vendor/
	godep save goweb-scaffold

install: ## Install the app
	go install goweb-scaffold

deps: ## Install required dependencies
	go get github.com/cihub/seelog
	go get github.com/codegangsta/negroni
	go get github.com/facebookgo/inject
	go get github.com/gorilla/mux
	go get github.com/mjibson/esc
	go get github.com/robfig/cron
	go get github.com/stretchr/testify
	go get github.com/tools/godep

run: ## Run goweb service as container
	bash scripts/docker-run.sh

test: asset build ## Run all tests
	bash scripts/test.sh

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
