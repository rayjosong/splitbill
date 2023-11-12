package handler

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rayjosong/splitbill/pkg/user"
)

// type UserCredentialsService interface {
// 	SaveCredentials(user user.User, username string, password string) error
// 	GetCredentialsForUser(username string) (*usercredentials.UserCredentials, error)
// 	IsCredentialValid(username string, password string) (bool, error)
// }

type UserSessionService interface {
	CreateUser(user.User, string, string) error
	GetUser(username string) (user.User, error)
}

type CredentialsVerifier interface {
	IsCredentialValid(inputPassword string, savedPassword string) bool
}

type SessionHandler struct {
	userSessionService UserSessionService
	verifier           CredentialsVerifier
}

func NewSessionHandler(userSessionService UserSessionService, verifier CredentialsVerifier) SessionHandler {
	return SessionHandler{
		userSessionService: userSessionService,
		verifier:           verifier,
	}
}

func (s SessionHandler) HandleSignup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	email := c.PostForm("email")
	name := c.PostForm("name")

	user := user.User{Name: name, Email: email}

	if err := s.userSessionService.CreateUser(user, username, password); err != nil {
		c.JSON(500, gin.H{"message": "Could not save credentials"})
	}

	c.JSON(201, gin.H{"message": "Saved user successfuly"})
}

func (s SessionHandler) HandleLogin(c *gin.Context) {
	username := c.PostForm("username")
	inputPassword := c.PostForm("password")

	user, err := s.userSessionService.GetUser(username)
	if err != nil {
		c.JSON(500, gin.H{"message": "failed to get credentials for user"})
		return
	}

	valid := s.verifier.IsCredentialValid(inputPassword, user.Password)
	if err != nil {
		c.JSON(500, gin.H{"message": "cannot find user"})
		return
	}

	if !valid {
		c.JSON(401, gin.H{"message": "Credentials incorrect"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.JSON(200, gin.H{
		"message": "Logged in successfully",
	})
}

func getCurrentUser(c *gin.Context) (*user.User, error) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID == nil {
		return nil, errors.New("user not logged in")
	}

	var user user.User
	result := s.db.First(&user, "id = ?", userID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
