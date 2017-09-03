package main

import (
	"flag"
	"net/http"
)

var (
	access = flag.String("access", "", "Access key")
	secret = flag.String("secret", "", "Secret key")
	bucket = flag.String("bucket", "", "Bucket")
)

func main() {
	flag.Parse()
	if *access == "" || *secret == ""  || *bucket == "" {
		flag.PrintDefaults()
		return
	}
	handler := &Handler{}
	go handler.Crawl()
	http.ListenAndServe(":80", handler)
}


