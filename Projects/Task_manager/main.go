package main

import (
	"Task_manager/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/checkServer", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "server is running"})
	})
	server.GET("/Task", handlers.GetTask)
	server.POST("/Task", handlers.CreateTask)
	server.PUT("/Task/:id", handlers.UpdateTask)
	server.Run(":8080")
}
