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
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.NewController(videoService)

	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
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

	// server.Use(gin.Recovery(), middlewares.Logger(),
	// 	middlewares.BasicAuth(), gindump.Dump())

	server.Use(gin.Recovery(), middlewares.Logger())

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)

		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
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
	port := os.Getenv("PORT")

	server.Run(":" + port)
}
