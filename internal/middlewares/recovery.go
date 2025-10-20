package middlewares

import (
	"fmt"

	"github.com/Habeebamoo/contact-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(gin.DefaultErrorWriter, func(c *gin.Context, err any) {
		utils.Error(c, 500, fmt.Sprint(err))
	})
}