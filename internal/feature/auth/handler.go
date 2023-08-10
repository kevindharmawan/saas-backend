package auth

import "github.com/gin-gonic/gin"

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	return
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	return
}
