## goweb-scaffold

This is a scaffold project for go web application with some convenient components as below.

- Dependency injection with [facebookgo-inject](github.com/facebookgo/inject)
- Portable dependencies with [godep](https://github.com/tools/godep)
- Logging with [seelog](github.com/cihub/seelog)
- HTTP middleware with [negroni](https://github.com/codegangsta/negroni)
- Testing with [testify](https://github.com/stretchr/testify)


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

1. Inside the container, cd into `/go/src/goweb-scaffold`, then install dependencies ( **this takes about 3min**; if you don't want to do this every time; use `docker commit` to save your changes )

  ``` sh
  cd /go/src/goweb-scaffold;
  go get github.com/stretchr/testify;
  godep restore goweb-scaffold  
  ```

1. Here your go

  ``` sh
  // test inside container
  sh scripts/test.sh

  // run app inside container
  go run main.go
  ```
