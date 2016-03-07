#!/bin/bash
# Usage:

# The name of middleware container
s=$(date +'%m%d')
t=$(date +'%s')
NAME="goweb-$s-$t"
PORT=$RANDOM

echo "Run goweb-scaffold container"
docker build -t goweb-scaffold .
docker run -d -e VIRTUAL_HOST=localhost -p $PORT:$PORT -v /var/log:/var/log --name $NAME --restart=always goweb-scaffold -port=$PORT
