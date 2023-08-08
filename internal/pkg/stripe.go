package pkg

import (
	"github.com/stripe/stripe-go/v74"
)

func InitializeStripe(key string) {
	stripe.Key = key
}
