package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rayjosong/splitbill/pkg/protocol/rest/handler/user"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	userHandler := user.NewUserHandler()
	router.POST("/users", userHandler.Create)

	router.Run()
}
