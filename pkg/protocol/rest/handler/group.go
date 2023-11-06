package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rayjosong/splitbill/pkg/group"
)

type GroupService interface {
	CreateGroup(group.Group) error
}

type GroupHandler struct {
	service GroupService
}

func NewGroupHandler(s GroupService) GroupHandler {
	return GroupHandler{service: s}
}

func (h GroupHandler) HandlePost(c *gin.Context) {
	var reqBody struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the currently logged in user
	user, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user must be logged in"})
		return
	}

	// Create new group
	group := group.Group{
		GroupID: uuid.New().String(),
		Name:    reqBody.Name,
		Members: []user.User{*user},
	}

	if err := h.service.CreateGroup(group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, group)
}
