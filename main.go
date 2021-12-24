package main

import (
	"fmt"
	"go-gin-app/controller"
	"go-gin-app/middlewares"
	"go-gin-app/service"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.NewVideo()
	videoController controller.VideoController = controller.NewController(videoService)
)

func setupLogOutput() {
	f, err := os.Create("./logs/gin.log")
	if err != nil {
		log.Printf("Open Log File Failed: %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	fmt.Println("Init Gin log set ")
}

func main() {

	setupLogOutput()
	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid!"})
			}

		})

	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}
	port, err := os.Getenv("PORT")

	if err != nil {
		port = "3000"
	}

	server.Run(":" + port)
}
