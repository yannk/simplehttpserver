package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var flagListen = flag.String("listen", ":8080", "socket address to listen to")

func main() {
	flag.Parse()

	var dir string
	if len(flag.Args()) > 0 {
		dir = flag.Args()[0]
	} else {
		var err error
		dir, err = os.Getwd()
		if err != nil {
			panic("can't get current working directory")
		}
	}
	log.Printf("starting serving %s on %s", dir, *flagListen)
	panic(http.ListenAndServe(*flagListen, http.FileServer(http.Dir(dir))))
}
