package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsMiddleware struct{}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (cm *CorsMiddleware) CorsMiddleware(c *gin.Context) {
	cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Authorization", "Origin"},
		ExposeHeaders:    []string{"Content-Type", "Content-Length"},
		AllowCredentials: true,
		AllowWebSockets:  true,
		MaxAge:           12 * time.Hour,
	})
	c.Next()
}
