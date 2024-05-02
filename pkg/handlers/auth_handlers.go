package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/todoApp/pkg/dtos"
)

func (h *Handler) SignIn(c *gin.Context) {
	var request dtos.InputUserForm

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid body request",
		})
		return
	}

	accessToken, err := h.services.GenerateJwt(request)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token": accessToken,
	})
}

func (h *Handler) SignUp(c *gin.Context) {
	var request dtos.InputUserForm

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid body request",
		})
		return
	}

	err = h.services.AddUser(request)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(200)
}

func (h *Handler) LogOut(c *gin.Context) {
	c.Header("Authorization", "")
	c.JSON(200, gin.H{
		"message": "goodbye!",
	})
}
