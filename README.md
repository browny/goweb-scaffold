## goweb-scaffold
[![GoDoc](https://godoc.org/github.com/browny/goweb-scaffold?status.svg)](http://godoc.org/github.com/browny/goweb-scaffold)
[![Travis Build Status](https://travis-ci.org/browny/goweb-scaffold.svg?branch=master)](https://travis-ci.org/browny/goweb-scaffold)


goweb-scaffold is a template project for go web application with some convenient components as below.

- Dependency injection with [facebookgo-inject](github.com/facebookgo/inject)
- Portable dependencies with [godep](https://github.com/tools/godep)
- Logging with [seelog](github.com/cihub/seelog)
- HTTP middleware with [negroni](https://github.com/codegangsta/negroni)
- Testing with [testify](https://github.com/stretchr/testify)

### Local Dev

1. Assume your environment with golang installed (For Mac: `brew install go`)
1. Install depedencies

  ``` sh
  make deps
  ```

1. Run http server

  ``` sh
  git clone git@github.com:browny/goweb-scaffold.git
  cd ./goweb-scaffold
  go run main.go
  ```

1. Run test

  ``` sh
  cd ./goweb-scaffold
  sh script/test.sh
  ```


### Use Docker to run it

1. Assume your environment equipped with Docker (ie. `docker ps` works)
1. Run below shell script, you will find 2 containers running for you (`nginx` is proxy, `goweb-xxx` is the scaffold app)

  ``` sh
  sh scripts/docker-run.sh

  // If above succeeds, you will get below
  $ docker ps
  CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                                      NAMES
  49d938a04188        nginx-proxy         "/app/docker-entrypoi"   38 minutes ago      Up 38 minutes       0.0.0.0:80->80/tcp, 0.0.0.0:443->443/tcp   nginx
  5745797c3caa        goweb-scaffold      "/go/bin/goweb-scaffo"   38 minutes ago      Up 38 minutes       0.0.0.0:28983->28983/tcp                   goweb-0310-1457577308
  ```


### Docker Dev

1. Make sure your system with Docker installed (check [how](https://docs.docker.com/engine/installation/))

1. Pull docker images (this takes seconds)

  ``` sh
  docker pull browny/go-docker-dev
  ```

1. Clone the source code

  ``` sh
  git clone git@github.com:browny/goweb-scaffold.git
  ```

1. Run the container under the source root

  ``` sh
  docker run --rm -it -v `pwd`:/go/src/goweb-scaffold -p 8000:8000 browny/go-docker-dev
  ```

1. Inside the container, cd into `/go/src/goweb-scaffold`, then install dependencies (few minutes)

  ``` sh
  cd /go/src/goweb-scaffold;
  make deps
  ```

1. Here your go

  ``` sh
  // test inside container
  sh scripts/test.sh

  // run app inside container
  go run main.go
  ```

### Start project from this scaffold

1. Change the package name
  - Linux
  
  	``` sh 
  	grep -rl  --exclude-dir=".git" goweb-scaffold . | xargs sed -i 's/goweb-scaffold/your-name/g'
  	```
  
  - Mac
  
  	``` sh 
  	grep -rl  --exclude-dir=".git" goweb-scaffold . | xargs sed -i '' -e 's/goweb-scaffold/your-name/g'
  	```

2. Reset the git upstream url

	``` sh
	git remote set-url origin <your project git repository url>
	```
