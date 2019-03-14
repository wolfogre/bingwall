package entity

import (
	"regexp"
	"time"
)

/*
https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-cn

{
    "images": [
        {
            "startdate": "20190313",
            "fullstartdate": "201903131600",
            "enddate": "20190314",
            "url": "/th?id=OHR.AgriculturalPi_ZH-CN9754138523_1920x1080.jpg&rf=NorthMale_1920x1080.jpg&pid=hp",
            "urlbase": "/th?id=OHR.AgriculturalPi_ZH-CN9754138523",
            "copyright": "圆形农田，科罗拉多州摩根县 (© Jim Wark/Getty Images)",
            "copyrightlink": "http://www.bing.com/search?q=%E5%9C%86%E5%BD%A2%E5%86%9C%E7%94%B0&form=hpcapt&mkt=zh-cn",
            "title": "",
            "quiz": "/search?q=Bing+homepage+quiz&filters=WQOskey:%22HPQuiz_20190313_AgriculturalPi%22&FORM=HPQUIZ",
            "wp": true,
            "hsh": "84d6f5070a773f3aa6e2b5b5a53a3435",
            "drk": 1,
            "top": 1,
            "bot": 1,
            "hs": [ ]
        }
    ],
    "tooltips": {
        "loading": "正在加载...",
        "previous": "上一个图像",
        "next": "下一个图像",
        "walle": "此图片不能下载用作壁纸。",
        "walls": "下载今日美图。仅限用作桌面壁纸。"
    }
}
*/

type Api struct {
	Images []ApiImage `json:"images"`
}

type ApiImage struct {
	StartDate     string `json:"startdate"`
	FullStartDate string `json:"fullstartdate"`
	EndDate       string `json:"enddate"`
	Url           string `json:"url"`
	UrlBase       string `json:"urlbase"`
	Copyright     string `json:"copyright"`
	CopyrightLink string `json:"copyrightlink"`
	Title         string `json:"title"`
	Quiz          string `json:"quiz"`
	Wp            bool   `json:"wp"`
	Hsh           string `json:"hsh"`
	Drk           int    `json:"drk"`
	Top           int    `json:"top"`
	Bot           int    `json:"bot"`
}

func (i ApiImage) ToHistory() History {
	ret := History{
		Id: i.EndDate,
		Info: i.Copyright,
		Time: time.Now(),
	}
	// AgriculturalPi_ZH-CN9754138523
	nameReg := regexp.MustCompile("[A-Za-z0-9]+_ZH-CN[0-9]+")
	// AgriculturalPi_ZH-CN9754138523_1920x1080.jpg
	filenameReg := regexp.MustCompile("[A-Za-z0-9]+_ZH-CN[0-9]+_1920x1080\\.jpg")
	ret.Name = nameReg.FindString()
}