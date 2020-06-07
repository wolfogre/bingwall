package web

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"bingwall/internal/db"
	"bingwall/internal/entity"
	"bingwall/internal/storage"
)

var (
	downloadM = &sync.Mutex{}
)

func download(c *gin.Context) {
	downloadM.Lock()
	defer downloadM.Unlock() // control speed

	date := c.Query("date")
	if _, err := time.ParseInLocation(entity.DayFormat, date, time.Local); err != nil {
		c.AbortWithStatus(400)
		return
	}

	history, err := db.FindHistory(date)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
		return
	}
	if history.Filename == "" {
		history.Filename = history.Name + "_UHD.jpg"
	}

	c.Header("Content-type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fmt.Sprintf("%s_%s", date, history.Filename))

	if c.Request.Method == http.MethodGet {
		if c.Request.Method == http.MethodGet {
			content, err := storage.DowloadFromQiniu(history.Filename)
			if err != nil {
				c.AbortWithStatusJSON(500, gin.H{
					"error": err,
				})
				return
			}
			c.Writer.Write(content)
		}
	}
}
