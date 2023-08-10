package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kevindharmawan/saas-backend/internal/feature/auth"
	"github.com/kevindharmawan/saas-backend/internal/feature/user"
	"github.com/kevindharmawan/saas-backend/internal/model"
	"github.com/kevindharmawan/saas-backend/internal/shared/constant"
	"github.com/kevindharmawan/saas-backend/internal/shared/dto"
)

type AuthMiddleware struct {
	validationService auth.ValidationService
	userService       user.UserService
}

func NewAuthMiddleware(
	validationService auth.ValidationService,
	userService user.UserService,
) *AuthMiddleware {
	return &AuthMiddleware{
		validationService: validationService,
		userService:       userService,
	}
}

// TODO: Add require verification implementation
func (am *AuthMiddleware) AuthMiddleware(c *gin.Context) {
	authorizationToken := c.GetHeader("Authorization")
	token := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if token == "" {
		err := model.NewUnauthorizedError(false)
		dto.SendErrorResponse(c, err)
		return
	}

	authId, err := am.validationService.ValidateToken(token)
	if err != nil {
		dto.SendErrorResponse(c, err)
		return
	}

	user, err := am.userService.GetUserByAuthId(authId)
	if err != nil {
		dto.SendErrorResponse(c, err)
		return
	}

	c.Set(constant.UserIdKey, user.ID)
	c.Set(constant.UserAuthIdKey, user.AuthID)
	c.Set(constant.UserCustomerIdKey, user.CustomerID)
	c.Next()
}
