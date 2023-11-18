package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rayjosong/splitbill/internal/models"
)

type GroupService interface {
	CreateGroup(models.Group) error
}

type GroupHandler struct {
	service           GroupService
	userSessionGetter userSessionGetter
}

type userSessionGetter interface {
	GetCurrentUser(c *gin.Context) (models.User, error)
}

func NewGroupHandler(s GroupService, userSessionGetter userSessionGetter) GroupHandler {
	return GroupHandler{
		service:           s,
		userSessionGetter: userSessionGetter,
	}
}

func (h GroupHandler) HandlePost(c *gin.Context) {
	var reqBody struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userSessionGetter.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user must be logged in"})
		return
	}

	group := models.Group{
		GroupID: uuid.New().String(),
		Name:    reqBody.Name,
		Members: []models.User{user},
	}

	if err := h.service.CreateGroup(group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, group)
}
