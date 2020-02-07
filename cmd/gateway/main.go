package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/apex/gateway"
	"github.com/carlmjohnson/feed2json"
)

func main() {
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := ""
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		http.Handle("/", http.FileServer(http.Dir("./public")))
	}

	http.HandleFunc("/api/echo", echo)
	http.Handle("/api/feed", feed2json.Handler(
		feed2json.StaticURLInjector("https://news.ycombinator.com/rss"), nil, nil, nil, addCache))
	log.Fatal(listener(portStr, nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "public, max-age=60")
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Printf("could not dump request: %v", err)
		return
	}
	w.Write(b)
}

func addCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=300")
		h.ServeHTTP(w, r)
	})
}
