package web

import (
	"github.com/gin-gonic/gin"
)

func download(c *gin.Context) {
	c.JSON(200, gin.H{
		"ok": true,
	})
}
