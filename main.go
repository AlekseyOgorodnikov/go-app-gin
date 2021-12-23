package main

import (
	"go-gin-app/controller"
	"go-gin-app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.NewVideo()
	videoController controller.VideoController = controller.NewController(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})

	server.Run(":8080")
}
