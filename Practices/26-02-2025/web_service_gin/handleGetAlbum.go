package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func handleGetAlbum(c *gin.Context, ){
	c.IndentedJSON(http.StatusOK, AlbumDetails())
}

func handleGetAlbumByID(c *gin.Context){

	id := c.Param("id")

	for _, a := range AlbumDetails() {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}