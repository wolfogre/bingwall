package main

import (
	"bingwall/internal/crawler"
	"bingwall/internal/db"
	"bingwall/internal/storage"
	"bingwall/internal/version"
	"bingwall/internal/web"
	"flag"
	"log"
)

var (
	access  = flag.String("access", "", "access key")
	secret  = flag.String("secret", "", "secret key")
	bucket  = flag.String("bucket", "", "bucket")
	mongo   = flag.String("mongo", "", "mongo url")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("version", version.Version())

	flag.Parse()
	if *access == "" || *secret == "" || *bucket == "" || *mongo == "" {
		flag.PrintDefaults()
		log.Fatal()
	}

	if err := db.Init(*mongo); err != nil {
		log.Panic(err)
	}
	storage.InitQiniu(*bucket, *access, *secret)

	go crawler.Run()

	if err := web.Run(); err != nil {
		log.Panic(err)
	}
}
