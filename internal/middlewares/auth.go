package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RequireAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientAPiKey := c.GetHeader("X-API-KEY")

		if clientAPiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"success": false,
				"message": "API Key Missing",
			})
			return 
		}

		apiKey := os.Getenv("API_KEY")

		if clientAPiKey != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"success": false,
				"message": "Invalid API Key",
			})
			return 
		}

		c.Next()
	}
}