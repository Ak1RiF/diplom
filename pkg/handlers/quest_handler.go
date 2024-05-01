package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/todoApp/pkg/dtos"
)

func (h *Handler) AllQuests(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	quests, err := h.services.GetUserQuests(userId)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, quests)
}
func (h *Handler) ByIdQuest(c *gin.Context) {
	questId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	quest, err := h.services.GetUserQuestById(questId, userId)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, quest)
}
func (h *Handler) PostQuest(c *gin.Context) {
	var request dtos.InputQuestDto

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if err = h.services.AddUserQuest(request, userId); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "quest added"})
}
func (h *Handler) PutQuest(c *gin.Context) {
	questId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	var request dtos.InputQuestDto

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if err = h.services.UpdateUserQuest(questId, userId, request); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "quest update"})
}
func (h *Handler) DeleteQuest(c *gin.Context) {
	questId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	userId, err := h.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	if err = h.services.RemoveUserQuest(questId, userId); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "quest deleted"})
}
