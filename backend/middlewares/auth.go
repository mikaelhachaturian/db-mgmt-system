package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

func CheckIfHasToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		WATSON_AUTH_TOKEN := os.Getenv("WATSON_AUTH_TOKEN")

		if c.GetHeader("Token") != WATSON_AUTH_TOKEN {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
