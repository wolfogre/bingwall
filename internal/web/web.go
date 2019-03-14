package web

import (
	"time"

	"bingwall/internal/crawler"

	"github.com/gin-gonic/gin"
)

func Run() error {
	r := gin.Default()
	r.GET("/_status", func(c *gin.Context) {
		if crawler.LatestDay != crawler.Today() {
			if time.Now().Hour() == 0 {
				c.JSON(200, gin.H{
					"latest_day": crawler.LatestDay,
					"forgiven": true,
				})
			} else {
				c.JSON(500, gin.H{
					"latest_day": crawler.LatestDay,
					"forgiven": false,
				})
			}
		} else {
			c.JSON(200, gin.H{
				"latest_day": crawler.LatestDay,
				"forgiven": false,
			})
		}
	})
	return r.Run(":80")
}