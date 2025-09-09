package handlers

import (
	"net/http"

	"github.com/Habeebamoo/contact-go/internal/models"
	"github.com/Habeebamoo/contact-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func Contact(c *gin.Context) {
	var contactReq models.Contact
	if err := c.ShouldBindJSON(&contactReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"success": false,
			"message": "Invalid JSON Format",
		})
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