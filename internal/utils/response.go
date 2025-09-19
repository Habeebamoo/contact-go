package utils

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"success": true,
		"status": code,
		"message": msg,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"success": false,
		"status": code,
		"message": msg,
	})
}

func Abort(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, gin.H{
		"success": false,
		"status": code,
		"message": msg,
	})
}