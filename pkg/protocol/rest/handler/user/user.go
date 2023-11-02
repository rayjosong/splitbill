package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Create(c *gin.Context) {
	// handler user creation
	fmt.Println("creating user")
}

func (u *UserHandler) Get(c *gin.Context) {
	// Handle getting a user
}

func (u *UserHandler) Update(c *gin.Context) {
	// Handle updating a user
}
