#!/bin/bash
# Usage: sh scripts/test.sh
 
echo 'Run test in parallel'
go test -p 4 ./config ./rest ./utility
