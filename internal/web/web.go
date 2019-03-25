package web

import (
	"github.com/gin-gonic/gin"
)

const (
	routerStatus   = "/_status"
	routerRobots   = "/robots.txt"
	routerFavicon  = "/favicon.ico"
	routerDownload = "/download"
)

func Run() error {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.LoggerWithWriter(gin.DefaultWriter, routerStatus, routerRobots, routerFavicon))

	engine.GET(routerStatus, status)
	engine.GET(routerDownload, download)
	engine.HEAD(routerDownload, download)

	return engine.Run(":80")
}
