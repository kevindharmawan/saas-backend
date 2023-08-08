package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kevindharmawan/saas-backend/internal/feature/auth"
	"github.com/kevindharmawan/saas-backend/internal/shared/constant"
)

type AuthMiddleware struct {
	validateService auth.ValidationService
}

func NewAuthMiddleware(vs auth.ValidationService) *AuthMiddleware {
	return &AuthMiddleware{
		validateService: vs,
	}
}

// TODO: Add require verification implementation
func (am *AuthMiddleware) AuthMiddleware(c *gin.Context) {
	authorizationToken := c.GetHeader("Authorization")
	token := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if token == "" {
		c.Set(constant.IsAuthenticatedKey, false)
		c.Next()
		return
	}

	authId, err := am.validateService.ValidateToken(token)
	if err != nil {
		c.Set(constant.IsAuthenticatedKey, false)
		c.Next()
		return
	}

	c.Set(constant.IsAuthenticatedKey, true)
	c.Set(constant.UserAuthIdKey, authId)
	c.Next()
}
