package entity

const (
	DayFormat = "20060102"
)

/*
curl "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-cn&uhd=1" | jq

{
  "images": [
    {
      "startdate": "20200606",
      "fullstartdate": "202006061600",
      "enddate": "20200607",
      "url": "/th?id=OHR.LaPertusa_ZH-CN7227946330_UHD.jpg&rf=LaDigue_UHD.jpg&pid=hp&w=1920&h=1080&rs=1&c=4",
      "urlbase": "/th?id=OHR.LaPertusa_ZH-CN7227946330",
      "copyright": "La Pertusa教堂，西班牙莱里达 (© bbsferrari/Getty Images)",
      "copyrightlink": "https://www.bing.com/search?q=La+Pertusa%E6%95%99%E5%A0%82&form=hpcapt&mkt=zh-cn",
      "title": "",
      "quiz": "/search?q=Bing+homepage+quiz&filters=WQOskey:%22HPQuiz_20200606_LaPertusa%22&FORM=HPQUIZ",
      "wp": true,
      "hsh": "8e7fc0d19c6744cf16569626d29a787f",
      "drk": 1,
      "top": 1,
      "bot": 1,
      "hs": []
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
