package main

import (
	"net/http"

	user "go-news/pkg/user"

	"github.com/gin-gonic/gin"

	mysql "go-news/internal/database"
)

func main() {
	db := mysql.DB()
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/signup", userHandler.SignUpHandler)
	r.POST("/signin", userHandler.SignInHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run("127.0.0.1:8000")
}
