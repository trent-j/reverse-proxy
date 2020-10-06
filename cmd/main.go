package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {

	bpDevURL := os.Getenv("BP_DEV_URL")
	cert := os.Getenv("GHES_CERT")
	key := os.Getenv("GHES_KEY")

	bpDev, _ := url.Parse(bpDevURL)

	director := func(req *http.Request) {
		req.URL.Scheme = bpDev.Scheme
		req.URL.Host = bpDev.Host
		req.Host = strings.Split(req.Host, ":")[0]
	}

	proxy := &httputil.ReverseProxy{Director: director}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Printf("Proxy listening on port: %d", 9000)

	log.Fatal(http.ListenAndServeTLS(":9000", cert, key, nil))
}
