package main

import (
	"GinPractice/controller"
	"GinPractice/middleware"
	"GinPractice/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	file, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, VideoController.Save(ctx))
	})
	server.Run(":8080")
}
