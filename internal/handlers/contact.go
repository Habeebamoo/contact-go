package handlers

import (
	"github.com/Habeebamoo/contact-go/internal/models"
	"github.com/Habeebamoo/contact-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func Contact(c *gin.Context) {
	contactReq := &models.Contact{}
	if err := c.ShouldBindJSON(&contactReq); err != nil {
		utils.Error(c, 400, "Invalid JSON Format")
		return
	}

	if err := contactReq.Validate(); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	err := utils.NotifyMe(contactReq)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}

	utils.Success(c, 200, "Message Sent")
}