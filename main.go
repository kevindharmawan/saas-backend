package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kevindharmawan/saas-backend/internal/api"
	"github.com/kevindharmawan/saas-backend/internal/feature/auth"
	"github.com/kevindharmawan/saas-backend/internal/feature/payment"
	"github.com/kevindharmawan/saas-backend/internal/feature/user"
	"github.com/kevindharmawan/saas-backend/internal/pkg"
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

	// Create data sources
	pkg.InitializeStripe(c.Stripe)
	db, err := pkg.InitializeSqlite(c.Database)
	if err != nil {
		panic(err)
	}

	// Create repositories
	authRepo := auth.NewAuthRepository(db)
	hashRepo := auth.NewHashRepository()
	tokenRepo := auth.NewTokenRepository()
	userRepo := user.NewUserRepository(db)
	paymentRepo := payment.NewPaymentRepository()

	// Create services
	authService := auth.NewAuthService(authRepo, hashRepo, tokenRepo, userRepo)
	userService := user.NewUserService(userRepo)
	paymentService := payment.NewPaymentService(paymentRepo)

	s := api.InitializeApi("", c.Server.Port, authService, userService, paymentService)

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
