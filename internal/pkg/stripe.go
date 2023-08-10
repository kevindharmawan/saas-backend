package pkg

import (
	"github.com/kevindharmawan/saas-backend/internal/shared/config"
	"github.com/stripe/stripe-go/v74"
)

func InitializeStripe(stripeConfig config.StripeConfig) {
	stripe.Key = stripeConfig.Key
}
