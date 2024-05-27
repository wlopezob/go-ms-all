package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wlopezob/go-firstendpoint/model"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/albumns", model.GetAlbumns)
	r.POST("/albumns", model.PostAlbums)
	r.GET("/albumns/:id", model.GetAlbumById)
	r.Run("0.0.0.0:6000")
}
