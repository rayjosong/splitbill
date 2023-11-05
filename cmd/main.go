package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/rayjosong/splitbill/pkg/protocol/rest/handler"
	"github.com/rayjosong/splitbill/pkg/protocol/rest/handler/user"
	userModel "github.com/rayjosong/splitbill/pkg/user"
	"github.com/rayjosong/splitbill/pkg/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func main() {
	dbCfg := DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.DBName, dbCfg.Password)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&userModel.User{}, &userModel.UserCredentials{})

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.POST("/signup", handler.HandleSignup)

	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		var credentials userModel.UserCredentials
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
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	userRepository := repository.NewUserRepo(db)
	userHandler := user.NewUserHandler(userRepository)
	router.POST("/users", userHandler.Create)
	router.POST("/users/:id/friends", userHandler.HandlePostFriends)

	router.Run()
}
