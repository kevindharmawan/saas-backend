package payment

import "github.com/kevindharmawan/saas-backend/internal/model"

type PaymentRepository interface {
	CreateCustomer(user *model.User) (*model.User, *model.AppError)
	GetCustomerByUserId(userId int64) (*model.User, *model.AppError)
	GetCreditCards(customerId string) ([]*model.Card, *model.AppError)
	AddCreditCard(customerId string, cardToken string) (*model.Card, *model.AppError)
	RemoveCreditCard(customerId string, cardId string) *model.AppError
	SetDefaultCreditCard(customerId string, cardId string) *model.AppError
	GetSubscriptionInfo(customerId string) (*model.Subscription, *model.AppError)
	Subscribe(customerId string, priceId string) (*model.Subscription, *model.AppError)
	CancelSubscription(customerId string) *model.AppError
}

type PaymentService interface {
	GetCreditCards(user *model.User) ([]*model.Card, *model.AppError)
	AddCreditCard(user *model.User, cardToken string) (*model.Card, *model.AppError)
	RemoveCreditCard(user *model.User, cardId string) *model.AppError
	SetDefaultCreditCard(user *model.User, cardId string) *model.AppError
	GetSubscriptionInfo(user *model.User) (*model.Subscription, *model.AppError)
	Subscribe(user *model.User, priceId string) (*model.Subscription, *model.AppError)
	CancelSubscription(user *model.User) *model.AppError
}
