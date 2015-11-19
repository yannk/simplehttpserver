Super simple tool to serve a filesystem directory over HTTP (similar to Python's eponymous module. Useful when Python is not installed on your system).

Go's standard library does all the work, we just need a little extra
config:
- port/address to use
- root directory to serve

```
$ go get -u github.com/yannk/simplehttpserver
$ simplehttpserver
```

or,

$ simplehttpserver -listen=127.0.0.1:8001 /usr/share/doc
