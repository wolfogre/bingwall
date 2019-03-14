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
	access  = flag.String("access", "", "Access key")
	secret  = flag.String("secret", "", "Secret key")
	bucket  = flag.String("bucket", "", "Bucket")
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

	crawler.RunDaemon()

	if err := web.Run(); err != nil {
		log.Panic(err)
	}
}
