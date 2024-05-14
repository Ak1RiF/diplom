package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/todoApp/pkg/dtos"
)

func (h *Handler) GetUserInfo(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	user, err := h.services.UserInfo(userId)
	if err != nil {
		c.JSON(500, gin.H{"message": "user not found"})
		return
	}

	c.JSON(200, gin.H{"user_info": user})
}

func (h *Handler) AddExpToUser(c *gin.Context) {
	var request dtos.UserExperienceInput

	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.UpdateUserExperience(userId, request); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, "User update")
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var request *dtos.UpdateUserFrom

	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.UpdateUser(userId, *request); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User was updated"})
}
