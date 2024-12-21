package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/handlers"
)

func SetupRouter(ctx context.Context) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/register", func(c *gin.Context) {
		resp := handlers.Register(c)

		c.JSON(resp.Status, resp.Message)
	})

	// Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := db[user]
	// 	if ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

	return r
}
