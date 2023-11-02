package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rayjosong/splitbill/pkg/protocol/rest/handler/user"
	userModel "github.com/rayjosong/splitbill/pkg/user"
	"github.com/rayjosong/splitbill/pkg/user/repository"
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

	db.AutoMigrate(&userModel.User{})

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	userRepository := repository.NewUserRepo()
	userHandler := user.NewUserHandler(userRepository)
	router.POST("/users", userHandler.Create)

	router.Run()
}
