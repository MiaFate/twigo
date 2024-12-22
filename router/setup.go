package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/handlers"
	"github.com/miafate/twigo/middlewares"
	"github.com/miafate/twigo/models"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/health-check", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.POST("/register", func(c *gin.Context) {
		resp := handlers.Register(c)

		c.JSON(resp.Status, resp.Message)
	})

	r.POST("/login", func(c *gin.Context) {
		resp := handlers.Login(c)
		c.PureJSON(resp.Status, resp.Data)
	})

	r.Use(middlewares.JwtMiddleware())
	r.GET("/profile", func(c *gin.Context) {
		resp := handlers.GetProfile(c)
		c.PureJSON(resp.Status, resp.Data)
	})

	r.GET("/users", func(c *gin.Context) {
		resp := handlers.GetUsers(c)
		c.PureJSON(http.StatusOK, resp.Data)
	})

	r.PUT("/user/:id", func(c *gin.Context) {
		resp := handlers.UpdateUser(c, c.MustGet("claim").(*models.Claim))
		c.PureJSON(resp.Status, resp)
	})

	r.POST("/post", func(c *gin.Context) {
		resp := handlers.AddPost(c, c.MustGet("claim").(*models.Claim))
		c.PureJSON(resp.Status, resp)
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
