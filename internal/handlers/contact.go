package handlers

import (
	"net/http"

	"github.com/Habeebamoo/contact-go/internal/models"
	"github.com/Habeebamoo/contact-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func Contact(c *gin.Context) {
	contactReq := &models.Contact{}
	if err := c.ShouldBindJSON(&contactReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"success": false,
			"message": "Invalid JSON Format",
		})
		return
	}

	if err := contactReq.Validate(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 400,
			"success": false,
			"message": err.Error(),
		})	
		return
	}

	err := utils.NotifyMe(contactReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"success": false,
			"message": err.Error(),
		})	
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"success": true,
		"message": "Message Sent",
	})
}