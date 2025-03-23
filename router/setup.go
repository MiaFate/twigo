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
	r.GET("/products", func(c *gin.Context) {
		resp := handlers.GetProducts(c)
		c.PureJSON(resp.Status, resp.Data)
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
	r.StaticFS("/images", http.Dir("public/images"))

	r.Use(middlewares.JwtMiddleware())
	r.GET("/profile", func(c *gin.Context) {
		resp := handlers.GetProfile(c)
		c.PureJSON(resp.Status, resp.Data)
	})

	r.GET("/users", func(c *gin.Context) {
		resp := handlers.GetUsers(c, c.MustGet("claim").(*models.Claim))
		c.PureJSON(http.StatusOK, resp.Data)
	})

	r.PUT("/user", func(c *gin.Context) {
		resp := handlers.UpdateUser(c, c.MustGet("claim").(*models.Claim))
		c.PureJSON(resp.Status, resp)
	})

	r.POST("/post", func(c *gin.Context) {
		resp := handlers.AddPost(c, c.MustGet("claim").(*models.Claim))
		c.PureJSON(resp.Status, resp)
	})
	// r.POST("/product", func(c *gin.Context) {
	// 	resp := handlers.AddProduct(c)
	// 	c.PureJSON(resp.Status, resp)
	// })
	r.POST("/products", func(c *gin.Context) {
		resp := handlers.AddProductsBulk(c)
		if resp != nil {
			c.PureJSON(200, resp)
		}
		c.PureJSON(400, "error")
	})

	r.GET("/posts", func(c *gin.Context) {
		resp := handlers.GetPosts(c)
		c.PureJSON(resp.Status, resp.Data)
	})
	r.GET("/friendsposts", func(c *gin.Context) {
		resp := handlers.GetFriendsPosts(c, c.MustGet("claim").(*models.Claim))
		c.PureJSON(resp.Status, resp)
	})

	r.DELETE("/post", func(c *gin.Context) {
		resp := handlers.DeletePost(c, c.MustGet("claim").(*models.Claim))
		c.PureJSON(resp.Status, resp.Message)
	})

	//r.MaxMultipartMemory = 8 << 20 // 8 MiB
	//r.Static("/", "./public")
	upload := r.Group("/upload")
	{
		upload.POST("/avatar", func(c *gin.Context) {
			resp := handlers.UploadImage(c, "A", c.MustGet("claim").(*models.Claim))
			c.PureJSON(resp.Status, resp)
		})
		upload.POST("/banner", func(c *gin.Context) {
			resp := handlers.UploadImage(c, "B", c.MustGet("claim").(*models.Claim))
			c.PureJSON(resp.Status, resp)
		})
	}

	r.POST("/addfriend", func(c *gin.Context) {
		resp := handlers.AddRelationship(c, c.MustGet("claim").(*models.Claim))
		c.PureJSON(resp.Status, resp)

	})

	r.DELETE("/delfriend", func(c *gin.Context) {
		resp := handlers.DeleteRelationship(c, c.MustGet("claim").(*models.Claim))
		c.PureJSON(resp.Status, resp)
	})

	r.GET("/checkfriend", func(c *gin.Context) {
		resp := handlers.GetRelationship(c, c.MustGet("claim").(*models.Claim))
		c.PureJSON(resp.Status, resp)
	})

	// pending
	// r.GET("/avatar", func(c *gin.Context) {
	// 	resp := handlers.GetImage(c, "A", c.MustGet("claim").(*models.Claim))
	// 	c.PureJSON(resp.Status, resp)
	// })

	// r.GET("/banner", func(c *gin.Context) {
	// 	resp := handlers.GetImage(c, "B", c.MustGet("claim").(*models.Claim))
	// 	c.PureJSON(resp.Status, resp)
	// })

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
