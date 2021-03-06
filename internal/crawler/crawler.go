package crawler

import (
	"log"
	"regexp"
	"runtime"
	"strings"
	"time"

	"bingwall/internal/db"
	"bingwall/internal/entity"
	"bingwall/internal/storage"
	"github.com/wolfogre/lamp/pkg/lampc"
)

const (
	rootUrl = "https://cn.bing.com"
	apiPath = "/HPImageArchive.aspx?format=js&idx=0&n=8&mkt=zh-cn"
)

var (
	LatestDay string
)

func Today() string {
	return time.Now().Format(entity.DayFormat)
}

func Start() {
	for {
		log.Println("wake up")

		run()

		nextDay := time.Now().AddDate(0, 0, 1)
		nextDay = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, nextDay.Location())

		if err := lampc.Prolong("bingwall", nextDay.Add(time.Hour)); err != nil {
			log.Printf("Prolong: %v", err)
		}

		log.Printf("sleep %v to %v\n", nextDay.Sub(time.Now()), nextDay)
		runtime.GC()
		time.Sleep(nextDay.Sub(time.Now()))
	}
}

func Run() {
	run()
	nextDay := time.Now().AddDate(0, 0, 1)
	nextDay = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, nextDay.Location())

	if err := lampc.Prolong("bingwall", nextDay.Add(time.Hour)); err != nil {
		log.Printf("Prolong: %v", err)
	}
}

func run() {
	delay := time.Second

RETRY:
	log.Printf("delay %v\n", delay)
	time.Sleep(delay)
	delay *= 2
	if delay > time.Hour {
		delay = time.Hour
	}

	infos, err := getImageInfos()
	if err != nil {
		log.Printf("get image info failed: %v\n", err)
		goto RETRY
	}

	for _, v := range infos.Images {
		image := entity.History{
			Id:   v.EndDate,
			Info: v.Copyright,
			Time: time.Now(),
		}
		// "urlbase": "/th?id=OHR.LaPertusa_ZH-CN7227946330"
		image.Name = regexp.MustCompile("[A-Za-z0-9]+_ZH-CN[0-9]+").FindString(v.UrlBase)
		if image.Name == "" {
			log.Printf("can't get name from %v\n", v.UrlBase)
			goto RETRY
		}
		image.Filename = image.Name + "_UHD.jpg"

		exists, err := db.ExistHistory(image.Id)
		if err != nil {
			log.Printf("check image existing failed: %v\n", err)
			goto RETRY
		}
		if exists {
			if strings.Compare(LatestDay, image.Id) < 0 {
				LatestDay = image.Id
				log.Printf("updated latest day to %v\n", LatestDay)
			}
			continue
		}

		log.Printf("find new date: %v\n", image.Id)

		fileUrl := v.UrlBase + "_UHD.jpg"
		if !strings.HasPrefix(fileUrl, "http") {
			fileUrl = rootUrl + fileUrl
		}
		content, err := downloadImage(fileUrl)
		if err != nil {
			log.Printf("download image failed: %v\n", err)
			goto RETRY
		}
		log.Printf("downloaded image from %v\n", fileUrl)

		if err := storage.UploadToQiniu(image.Filename, content); err != nil {
			log.Printf("upload image failed: %v\n", err)
			goto RETRY
		}
		log.Printf("uploaded image to qiniu: %v\n", image.Filename)

		if err := db.InsertHistory(image); err != nil {
			log.Printf("insert history failed: %v\n", err)
			goto RETRY
		}
		log.Printf("inserted history: %+v\n", image)

		if strings.Compare(LatestDay, image.Id) < 0 {
			LatestDay = image.Id
			log.Printf("updated latest day to %v\n", LatestDay)
		}
	}

	if LatestDay != Today() {
		log.Printf("latest day isn't today\n")
		goto RETRY
	}
}
