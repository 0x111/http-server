package main

import (
	"flag"
	"fmt"
	"github.com/NYTimes/gziphandler"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const appVersion = "v0.1.1"

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
// @src: https://github.com/gorilla/mux
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
// @src: https://github.com/gorilla/mux
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	mux := http.NewServeMux()

	var host string
	var port string
	var path string
	var version bool
	var useGzip bool
	var spa bool
	var staticPath string
	var indexPath string

	flag.StringVar(&host, "host", "localhost", "hostname to listen on, default: localhost")
	flag.StringVar(&port, "port", "8080", "port to listen on, default: 8080")
	flag.StringVar(&path, "path", ".", "provide the path which we need to serve, default: ./")
	flag.BoolVar(&useGzip, "gzip", false, "use gzip compression, default: false")
	flag.BoolVar(&spa, "spa", false, "set to true, to host Single Page Applications, default: false")
	flag.StringVar(&indexPath, "indexPath", "index.html", "path to the index.html file, default: index.html")
	flag.StringVar(&staticPath, "staticPath", "", "path to the index.html file, default: empty")

	flag.BoolVar(&version, "version", false, "prints version")
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

	if spa {
		// in case the SPA is set, the path will be overwritten with staticPath
		path = staticPath
		fs = spaHandler{staticPath: staticPath, indexPath: indexPath}
	}

	mux.Handle("/", fs)

	fmt.Printf("Listening on %s:%s and serving path %s\n", host, port, path)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), mux))
}
