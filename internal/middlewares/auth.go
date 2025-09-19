package middlewares

import (
	"os"

	"github.com/Habeebamoo/contact-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func RequireAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientAPiKey := c.GetHeader("X-API-KEY")

		if clientAPiKey == "" {
			utils.Abort(c, 401, "API Key Missing")
			return 
		}

		apiKey := os.Getenv("API_KEY")

		if clientAPiKey != apiKey {
			utils.Abort(c, 401, "Invalid API Key")
			return 
		}

		c.Next()
	}
}