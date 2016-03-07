#!/bin/bash

echo "Re generate dependencies"
rm -rf config/asset.go
esc -o config/asset.go -pkg config config/

rm -rf Godeps/
rm -rf vendor/
godep save goweb-scaffold

echo "Build all ..."
GOMAXPROCS=4 GOGC=400 go install goweb-scaffold
