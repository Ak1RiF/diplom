package handlers

import "github.com/gin-gonic/gin"

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

	c.JSON(200, user)
}
