package web

import (
	"bingwall/internal/crawler"
	"github.com/gin-gonic/gin"
	"time"
)

func status(c *gin.Context) {
	if crawler.LatestDay != crawler.Today() {
		if time.Now().Hour() == 0 {
			c.JSON(200, gin.H{
				"ok": true,
				"latest_day": crawler.LatestDay,
				"forgiven": true,
			})
		} else {
			c.JSON(500, gin.H{
				"ok": false,
				"latest_day": crawler.LatestDay,
				"forgiven": false,
			})
		}
	} else {
		c.JSON(200, gin.H{
			"ok": true,
			"latest_day": crawler.LatestDay,
			"forgiven": false,
		})
	}
}