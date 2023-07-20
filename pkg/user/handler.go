package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (uh *UserHandler) SignUpHandler(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = uh.userService.SignUp(user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully"})
}

func (uh *UserHandler) SignInHandler(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := uh.userService.SignIn(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// You can generate a JWT token here and send it back to the client as a response.
	// For simplicity, I'm omitting the token generation in this example.

	// For demonstration purposes, I'm returning the user data in the response.
	c.JSON(http.StatusOK, user)
}
