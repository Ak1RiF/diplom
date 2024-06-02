package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type eggResponse struct {
	Id     int    `json:"id"`
	Rarity string `json:"rarity"`
	Count  int    `json:"count"`
}

func (h *Handler) GetEggById(c *gin.Context) {
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

	egg, err := h.services.GetUserEggById(userId, eggId)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"egg": egg})
}

func (h *Handler) GetEggs(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	eggs, err := h.services.GetUserEggs(userId)

	response := []eggResponse{
		eggResponse{Id: 1, Rarity: "common", Count: eggs.CountCommon},
		eggResponse{Id: 2, Rarity: "rare", Count: eggs.CountRare},
		eggResponse{Id: 3, Rarity: "epic", Count: eggs.CountEpic},
		eggResponse{Id: 4, Rarity: "legendary", Count: eggs.CountLegendary},
	}
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user_eggs": response})
}

func (h *Handler) AddToCount(c *gin.Context) {
	eggId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.UpdateCountEggs(userId, eggId, "add"); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, "Ok!")
}

func (h *Handler) RemoveFromCount(c *gin.Context) {
	eggId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.UpdateCountEggs(userId, eggId, "remove"); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, "Ok!")
}
