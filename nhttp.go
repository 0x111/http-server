package main

import (
	"flag"
	"fmt"
	"github.com/NYTimes/gziphandler"
	"log"
	"net/http"
)

const appVersion = "v0.1"

func main() {
	mux := http.NewServeMux()

	var host string
	flag.StringVar(&host, "host", "localhost", "hostname to listen on, default: localhost")
	var port string
	flag.StringVar(&port, "port", "8080", "port to listen on, default: 8080")
	var path string
	flag.StringVar(&path, "path", ".", "provide the path which we need to serve, default: ./")
	var version bool
	flag.BoolVar(&version, "version", false, "prints version")
	var useGzip bool
	flag.BoolVar(&useGzip, "gzip", false, "use gzip compression, default: false")
	flag.Parse()

	// no host provided, default to localhost
	if host == "" {
		host = "localhost"
	}

	// no port provided, default to 8080
	if port == "" {
		port = "8080"
	}

	// print version information
	if version {
		fmt.Printf("nhttp %s\n", appVersion)
		return
	}

	fs := http.FileServer(http.Dir(path))

	// if gzip enabled
	if useGzip {
		fs = gziphandler.GzipHandler(fs)
	}

	mux.Handle("/", fs)

	fmt.Printf("Listening on %s:%s and serving path %s\n", host, port, path)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), mux))
}
