#!/bin/bash
# Usage:

ENV=$1

# The name of middleware container
s=$(date +'%m%d')
t=$(date +'%s')
PREFIX_NAME="goweb"
NAME="$PREFIX_NAME-$s-$t"
PORT=$RANDOM

OLD=$(docker ps | grep $PREFIX_NAME | awk "{print \$1}")

echo "Run goweb-scaffold container"
docker build -t goweb-scaffold .
docker run -d -e VIRTUAL_HOST=localhost \
	-p $PORT:$PORT \
	-v /var/log:/var/log \
	--name $NAME \
	--restart=always \
	goweb-scaffold -env=$ENV -port=$PORT

echo "Run nginx container"
if docker ps | grep -q nginx; then
  echo "nginx existed, do nothing"
else
  docker build -t nginx-proxy ./nginx-proxy
  docker run -d -p 80:80 -p 443:443 \
	  -v /etc/ssl/certs:/etc/nginx/certs \
	  -v /var/run/docker.sock:/tmp/docker.sock \
	  --name nginx \
	  --restart=always \
	  nginx-proxy
fi

# Clean
if [ -z "$OLD" ]; then
  echo "OLD not existed, do nothing"
else
  echo "Remove old"
  docker stop $OLD
  docker rm $OLD
fi

echo "Remove old unused images"
if docker images | grep -q "<none>"; then
  docker rmi -f $(docker images | grep "<none>" | awk "{print \$3}")
fi
