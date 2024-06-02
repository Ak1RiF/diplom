package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/todoApp/pkg/dtos"
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

	if len(pets) == 0 {
		c.JSON(200, gin.H{"pets": []*dtos.OutputPet{}})
		return
	}
	c.JSON(200, gin.H{"pets": pets})
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
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}
	var request dtos.CreatePet
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err = h.services.CreatePetToUser(userId, request); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "User has pet"})
}

func (h *Handler) UpdatePet(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	var request dtos.UpdatePet
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.ChangePet(userId, request); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "The pet has been updated"})
}
