# Hello World app in Go
This is a simple 'hello world' app based on go standard library.

### Setup Go environment
 
* [go](http://golang.org/doc/install) installation 
* [GOPATH](https://code.google.com/p/go-wiki/wiki/GOPATH) setup
* [godep](https://github.com/tools/godep) dependency management

App directory:

    $GOPATH/src/github.com/cloudControl/go_hello_world_app 

### Run local

~~~ bash
$ go build .
$ ./go_hello_world_app
~~~

### Deploy app via custom buildpack

~~~ bash
$ cctrlapp APP_NAME create custom --buildpack https://github.com/cloudControl/buildpack-go
$ cctrlapp APP_NAME/default push --ship
~~~

