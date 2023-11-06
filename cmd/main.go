package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rayjosong/splitbill/pkg/protocol/rest/handler"
	userModel "github.com/rayjosong/splitbill/pkg/user"
	"github.com/rayjosong/splitbill/pkg/usercredentials"
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

	db.AutoMigrate(&userModel.User{}, &usercredentials.UserCredentials{})

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// **************** API *****************
	SessionHandler := handler.NewSessionHandler()
	router.POST("/signup", SessionHandler.HandleSignup)

	router.POST("/login", SessionHandler.HandleLogin)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	GroupHandler := handler.NewGroupHandler()
	router.POST("/api/groups", GroupHandler.HandlePost)
	// ************ END OF API ************

	router.Run()
}
