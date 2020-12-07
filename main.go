package main

import (
	"flag"
	"log"

	"bingwall/internal/crawler"
	"bingwall/internal/db"
	"bingwall/internal/storage"
	"bingwall/internal/version"
)

var (
	access = flag.String("access", "", "access key")
	secret = flag.String("secret", "", "secret key")
	bucket = flag.String("bucket", "", "bucket")
	domain = flag.String("domain", "", "download domain")
	mongo  = flag.String("mongo", "", "mongo url")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("version", version.Version())

	flag.Parse()
	if *access == "" || *secret == "" || *bucket == "" || *domain == "" || *mongo == "" {
		flag.PrintDefaults()
		log.Fatal()
	}

	if err := db.Init(*mongo); err != nil {
		log.Panic(err)
	}
	storage.InitQiniu(*domain, *bucket, *access, *secret)

	crawler.Run()
}
