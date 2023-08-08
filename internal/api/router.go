package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kevindharmawan/saas-backend/internal/api/middleware"
	"github.com/kevindharmawan/saas-backend/internal/feature/auth"
	"github.com/kevindharmawan/saas-backend/internal/feature/payment"
	"github.com/kevindharmawan/saas-backend/internal/feature/user"
)

func InitRouter(
	corsMiddleware *middleware.CorsMiddleware,
	authMiddleware *middleware.AuthMiddleware,
	authHandler *auth.AuthHandler,
	userHandler *user.UserHandler,
	paymentHandler *payment.PaymentHandler,
) *gin.Engine {
	app := gin.Default()

	apiRoute := app.Group("/api")
	apiRoute.Use(corsMiddleware.CorsMiddleware)
	apiRoute.Use(authMiddleware.AuthMiddleware)
	{
		apiRoute.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// router.AuthRouter(apiRoute, as)
	// router.UserRouter(apiRoute, us)
	// router.CardRouter(apiRoute, ps, us)

	return app
}
