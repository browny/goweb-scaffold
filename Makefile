.PHONY: asset build snapshot-deps install deps run test cover
.DEFAULT_GOAL := help

asset: ## Build asset
	rm -rf config/asset.go
	esc -o config/asset.go -pkg config -ignore=".go" -ignore="DS_Store" config/

build: ## Compile all packages
	go build $(shell go list ./... | grep -v /vendor/)

save-deps: ## Save deps (Godeps)
	rm -rf Godeps/
	rm -rf vendor/
	godep save goweb-scaffold

restore-deps: ## Restore deps (Godeps)
	godep restore goweb-scaffold
	go get github.com/mjibson/esc
	go get github.com/stretchr/testify

install: ## Install the app
	go install goweb-scaffold

run: ## Run goweb service as container
	bash scripts/docker-run.sh

test: asset build ## Run all tests
	go test -v $(shell go list ./... | grep -v vendor | grep /)

cover: asset build
	sh scripts/coverage

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
