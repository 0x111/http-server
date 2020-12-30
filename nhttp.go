package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const appVersion = "v0.1"

func main() {
	var host string
	flag.StringVar(&host, "host", "localhost:8080", "hostname to listen on, example: localhost:8080")
	var path string
	flag.StringVar(&path, "path", ".", "provide the path which we need to serve, example: ./")
	version := flag.Bool("version", false, "prints version")
	flag.Parse()

	if *version {
		fmt.Printf("nhttp %s\n", appVersion)
		return
	}

	fmt.Printf("Listening on %s and serving path %s\n", host, path)

	http.Handle("/", http.FileServer(http.Dir(path)))
	log.Fatal(http.ListenAndServe(host, nil))
}
