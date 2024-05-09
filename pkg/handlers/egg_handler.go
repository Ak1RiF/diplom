package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetEggs(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	eggs, err := h.services.GetUserEggs(userId)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, eggs)
}
func (h *Handler) PostEggs(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	eggId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if err = h.services.AddEggToUser(userId, eggId); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "User has egg"})
}

func (h *Handler) AddToCountEggs(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	eggId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.AddToCountEgg(userId, eggId); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Egg was added"})
}

func (h *Handler) TakeFromCountEggs(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	eggId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.TakeFromCountEgg(userId, eggId); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Egg was removed"})
}
