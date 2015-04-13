Super simple tool to serve a filesystem directory over HTTP.

Go's standard library does all the work, we just need a little extra
config:
- port/address to use
- root directory to serve

$ go get github.com/yannk/simplehttpserver
$ go install github.com/yannk/simplehttpserver
$ simplehttpserver

or,

$ simplehttpserver -listen=127.0.0.1:8001 /usr/share/doc
