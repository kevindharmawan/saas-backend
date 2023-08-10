package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kevindharmawan/saas-backend/internal/api/middleware"
	"github.com/kevindharmawan/saas-backend/internal/feature/auth"
	"github.com/kevindharmawan/saas-backend/internal/feature/payment"
	"github.com/kevindharmawan/saas-backend/internal/feature/user"
)

func InitializeApi(
	host string,
	port int,
	authService auth.AuthService,
	userService user.UserService,
	paymentService payment.PaymentService,
) *http.Server {
	corsMiddleware := middleware.NewCorsMiddleware()
	authMiddleware := middleware.NewAuthMiddleware(authService, userService)

	authHandler := auth.NewAuthHandler(authService)
	userHandler := user.NewUserHandler(userService)
	paymentHandler := payment.NewPaymentHandler(paymentService)

	return &http.Server{
		Addr: fmt.Sprintf("%v:%d", host, port),
		Handler: InitializeRouter(
			corsMiddleware,
			authMiddleware,
			authHandler,
			userHandler,
			paymentHandler,
		),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
