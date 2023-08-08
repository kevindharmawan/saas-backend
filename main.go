package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kevindharmawan/saas-backend/internal/api"
	"github.com/kevindharmawan/saas-backend/internal/feature/auth"
	"github.com/kevindharmawan/saas-backend/internal/feature/payment"
	"github.com/kevindharmawan/saas-backend/internal/feature/user"
	"github.com/kevindharmawan/saas-backend/internal/shared/config"
)

// TODO: Make logger
func main() {
	c, err := config.InitializeConfig()
	if err != nil {
		panic(err)
	}

	if c.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if err != nil {
		panic(err)
	}

	// TODO: Create data sources

	// Create repositories
	authRepo := auth.NewAuthRepository()
	hashRepo := auth.NewHashRepository()
	tokenRepo := auth.NewTokenRepository()
	userRepo := user.NewUserRepository()
	paymentRepo := payment.NewPaymentRepository()

	// Create services
	authService := auth.NewAuthService(authRepo, hashRepo, tokenRepo)
	userService := user.NewUserService(userRepo)
	paymentService := payment.NewPaymentService(paymentRepo)

	s := api.InitApi("", c.Server.Port, authService, userService, paymentService)

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
