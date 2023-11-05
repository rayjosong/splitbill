package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rayjosong/splitbill/pkg/user"
)

type UserHandler struct {
	service UserService
}

type UserService interface {
	CreateUser(u user.User) error
	GetUserDetailsByID(id string) (user.User, error)
	GetUserDetailsByEmail(email string) (user.User, error)
}

func NewUserHandler(service UserService) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (u UserHandler) Create(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	if err := u.service.CreateUser(user); err != nil {
		c.JSON(500, gin.H{
			"message": "Error inserting user",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func (u *UserHandler) Get(c *gin.Context) {
	// Handle getting a user
}

func (u *UserHandler) Update(c *gin.Context) {
	// Handle updating a user
}
