package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kevindharmawan/saas-backend/internal/api/middleware"
	"github.com/kevindharmawan/saas-backend/internal/feature/auth"
	"github.com/kevindharmawan/saas-backend/internal/feature/payment"
	"github.com/kevindharmawan/saas-backend/internal/feature/user"
)

func InitializeRouter(
	corsMiddleware *middleware.CorsMiddleware,
	authMiddleware *middleware.AuthMiddleware,
	authHandler *auth.AuthHandler,
	userHandler *user.UserHandler,
	paymentHandler *payment.PaymentHandler,
) *gin.Engine {
	app := gin.Default()

	apiRoute := app.Group("/api")
	apiRoute.Use(corsMiddleware.CorsMiddleware)
	{
		apiRoute.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	authRoute := apiRoute.Group("/auth")
	{
		authRoute.POST("/signin", authHandler.SignIn)
		authRoute.POST("/signup", authHandler.SignUp)
	}

	userRoute := apiRoute.Group("/user")
	userRoute.Use(authMiddleware.AuthMiddleware)
	{
		userRoute.GET("", userHandler.GetCurrentUser)
		userRoute.PUT("", userHandler.UpdateUser)
		userRoute.DELETE("", userHandler.DeleteUser)
	}

	paymentRoute := apiRoute.Group("/payment")
	paymentRoute.Use(authMiddleware.AuthMiddleware)
	{
		cardRoute := paymentRoute.Group("/card")
		{
			cardRoute.GET("/card", paymentHandler.GetCreditCards)
			cardRoute.POST("/card", paymentHandler.AddCreditCard)
			cardRoute.DELETE("/card/:cardId", paymentHandler.RemoveCreditCard)
			cardRoute.PUT("/card/:cardId", paymentHandler.SetDefaultCreditCard)
		}

		subscriptionRoute := paymentRoute.Group("/subscription")
		{
			subscriptionRoute.GET("", paymentHandler.GetSubscriptionInfo)
			subscriptionRoute.POST("", paymentHandler.Subscribe)
			subscriptionRoute.DELETE("", paymentHandler.CancelSubscription)
		}
	}

	return app
}
