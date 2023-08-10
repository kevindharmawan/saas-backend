package user

import "github.com/gin-gonic/gin"

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	return
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	return
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	return
}
