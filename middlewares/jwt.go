package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/jwt"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{
				"message": "token not found",
			})
			c.Abort()
			return
		}
		claim, isOk, msg, err := jwt.ProcesoToken(tokenString, os.Getenv("JWTSIGN"))
		if !isOk {
			if err != nil {
				c.JSON(401, gin.H{
					"message": msg,
					"error":   err.Error(),
				})
				c.Abort()
			} else {
				c.JSON(401, gin.H{
					"message": msg,
				})
				c.Abort()
			}
		}
		c.Set("claim", claim)
		c.Next()
	}
}
