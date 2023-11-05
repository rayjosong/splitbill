package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rayjosong/splitbill/pkg/user"
	"golang.org/x/crypto/bcrypt"
)

type UserCredentialsService interface {
	SaveCredentials(user user.User, username string, password string) error
	CheckCredentials(username string, password string) (bool, error)
}

type SessionHandler struct {
	service UserCredentialsService
}

func NewSessionHandler() SessionHandler {
	return SessionHandler{}
}

func (s SessionHandler) HandleSignup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	email := c.PostForm("email")
	name := c.PostForm("name")

	user := user.User{Name: name, Email: email}

	if err := s.service.SaveCredentials(user, username, password); err != nil {
		c.JSON(500, gin.H{"message": "Could not save credentials"})
	}

	c.JSON(201, gin.H{"message": "Saved user successfuly"})
}

func (s SessionHandler) HandleLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid, err := s.service.CheckCredentials(username, password)
	if err != nil {
		c.JSON(500, gin.H{"message": "cannot find user"})
	}

	if !valid {
		c.JSON(401, gin.H{"message": "Credentials incorrect"})
	}
	session := sessions.Default(c)
	session.Set("user_id", credentials.UserID)
	session.Save()

	c.JSON(200, gin.H{
		"message": "Logged in successfully",
	})
}
