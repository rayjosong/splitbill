package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rayjosong/splitbill/pkg/user"
)

type UserHandler struct {
	repo UserRepository
}

type UserRepository interface {
	InsertUser(u user.User) error
}

func NewUserHandler(repo UserRepository) UserHandler {
	return UserHandler{
		repo: repo,
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

	if err := u.repo.InsertUser(user); err != nil {
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
