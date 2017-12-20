package main

import (
	"net/http"
	"time"
	"log"
	"io/ioutil"
	"encoding/json"
	"context"
	"path/filepath"

	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/api.v7/auth/qbox"
	"gopkg.in/mgo.v2/bson"
)

type Handler struct {
	Finished string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/robots.txt" {
		w.Write([]byte("User-agent: *\nDisallow: /"))
		return
	}

	if r.RequestURI == "/status" {
		now := time.Now().Local()
		nowFormat := now.Format("20060102")
		if now.Hour() == 0 && now.Minute() < 5 { // 凌晨留足5分钟供爬虫工作
			w.Write([]byte(nowFormat))
			return
		}
		if h.Finished == nowFormat {
			w.Write([]byte(h.Finished))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(h.Finished))
			return
		}
	}
	http.Error(w, "400 Not Found", http.StatusNotFound)
}

func (h *Handler) Crawl() {
	for {
		imageInfo := GetImage()
		DownloadImage("https://cn.bing.com" + imageInfo.Images[0].Url)
		SaveMongo(imageInfo)
		h.Finished = time.Now().Local().Format("20060102")
		nextDay := time.Now().Local().Add(24 * time.Hour)
		nextDay = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, nextDay.Location())
		time.Sleep(nextDay.Sub(time.Now()))
	}
}

type JsonResponse struct {
	Images [] struct{
		Enddate string `json:"enddate"`
		Url     string `json:"url"`
		Urlbase string `json:"urlbase"`
		Copyright string `json:"copyright"`
	} `json:"images"`
}

func GetImage() JsonResponse {
RETRY:
	time.Sleep(time.Second)
	res, err := http.Get("http://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")
	if err != nil {
		log.Println(err)
		goto RETRY
	}
	if res.StatusCode != http.StatusOK {
		log.Println(res.Status)
		goto RETRY
	}
	buffer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		goto RETRY
	}
	jr := JsonResponse{}
	if err := json.Unmarshal(buffer, &jr); err != nil {
		log.Println(err, string(buffer))
		goto RETRY
	}
	if len(jr.Images) != 1 || jr.Images[0].Enddate != time.Now().Local().Format("20060102") || jr.Images[0].Url == "" {
		log.Println(time.Now().Local().Format("20060102"))
		log.Println(string(buffer))
		goto RETRY
	}

	return jr
}

func DownloadImage(url string) {
RETRY:
	time.Sleep(time.Second)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		goto RETRY
	}
	if res.StatusCode != http.StatusOK {
		log.Println(res.Status)
		goto RETRY
	}

	policy := &storage.PutPolicy{
		Scope: *bucket,
	}
	token := policy.UploadToken(qbox.NewMac(*access, *secret))
	uploader := storage.NewFormUploader(&storage.Config{
		Zone: &storage.ZoneHuadong,
		UseHTTPS: false,
		UseCdnDomains: false,
	})

	if err := uploader.Put(context.Background(), nil, token, filepath.Base(url), res.Body, int64(res.ContentLength), nil); err != nil {
		log.Println(err)
		goto RETRY
	}
}

func SaveMongo(response JsonResponse) {
	session := Session.Copy()
	defer session.Close()

	retry := true
RETRY:
	if retry {
		session.Refresh()
		time.Sleep(time.Second)
	}
	retry = true

	updateMap := bson.M{
		"$set": bson.M{
			"name": filepath.Base(response.Images[0].Urlbase),
			"url": "https://cn.bing.com" + response.Images[0].Url,
			"info": response.Images[0].Copyright,
			"time": time.Now(),
		},
	}

	_, err := session.DB(MONGO_DB).C(MONGO_C).UpsertId(response.Images[0].Enddate, updateMap)
	if err != nil {
		log.Printf("mongo upsert id failed: %v\n", err)
		goto RETRY
	}
}