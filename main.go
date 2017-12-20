package main

import (
	"flag"
	"net/http"

	"gopkg.in/mgo.v2"
	"log"
)

const (
	MONGO_DB = "bingwall"
	MONGO_C = "history"
)
var (
	access = flag.String("access", "", "Access key")
	secret = flag.String("secret", "", "Secret key")
	bucket = flag.String("bucket", "", "Bucket")
	mongo = flag.String("mongo", "", "mongo url")
	Session *mgo.Session
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Parse()
	if *access == "" || *secret == ""  || *bucket == "" || *mongo == "" {
		flag.PrintDefaults()
		log.Fatal()
	}

	var err error
	Session, err = mgo.Dial(*mongo)
	if err != nil {
		log.Fatal(err)
	}

	handler := &Handler{}
	go handler.Crawl()
	http.ListenAndServe(":80", handler)
}


