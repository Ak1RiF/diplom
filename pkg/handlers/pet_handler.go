package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPets(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	pets, err := h.services.GetUserPets(userId)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, pets)
}

func (h *Handler) GetPetsById(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	petId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	pet, err := h.services.GetUserPet(userId, petId)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, pet)
}
func (h *Handler) PostPets(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	petId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if err = h.services.AddPetToUser(userId, petId); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "User has pet"})
}
