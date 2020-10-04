package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {

	var (
		bpDevURL = flag.String("url", "", "bp-dev url")
		cert     = flag.String("cert", "/app/bp-dev.crt", "path to SSL certificate")
		key      = flag.String("key", "/app/bp-dev.key", "path to SSL certificate key")
		port     = flag.Int("port", 9000, "port to listen on")
	)

	flag.Parse()

	bpDev, _ := url.Parse(*bpDevURL)

	director := func(req *http.Request) {
		req.URL.Scheme = bpDev.Scheme
		req.URL.Host = bpDev.Host
		req.Host = strings.Split(req.Host, ":")[0]
	}

	proxy := &httputil.ReverseProxy{Director: director}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", *port), *cert, *key, nil))
}
