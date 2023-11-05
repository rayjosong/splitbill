package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rayjosong/splitbill/pkg/user"
)

func HandleSignup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	email := c.PostForm("email")
	name := c.PostForm("name")

	user := user.User{Name: name, Email: email}

	// save user details into DB

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"message": "Could not hash password"})
		return
	}
	credentials := user.UserCredentials{UserID: user.ID, Username: username, Password: string(hashedPassword)}
	// save user credential details into D
}

func HandleLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var credentials user.UserCredentials
	// search DB for credentials based on username and password

	if err := bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password)); err != nil {
		c.JSON(401, gin.H{"message": "Credentials incorrect"})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", credentials.UserID)
	session.Save()

	c.JSON(200, gin.H{
		"message": "Logged in successfully",
	})
}
