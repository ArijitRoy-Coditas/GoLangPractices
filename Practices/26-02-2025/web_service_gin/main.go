package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/album", handleGetAlbum)
	router.GET("/album:id", handleGetAlbumByID)
	router.POST("/album",handlePostAlbum)
	
	router.Run("localhost:8080")
}