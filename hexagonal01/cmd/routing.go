package main

import (
	"github.com/gin-gonic/gin"
	ports "github.com/wlopezob/hexagonal01/cmd/services/dashboard-api/ports/drivers"
)

func CreateRouter(userAdapter ports.ForUser) *gin.Engine {
	router := gin.Default()
	router.GET("/getUser", func(c *gin.Context) {
		username := c.Query("username")
		user, err := userAdapter.GetUser(username)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})
	return router
}
