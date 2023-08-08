package payment

import "github.com/kevindharmawan/saas-backend/internal/model"

type paymentRepositoryImpl struct{}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepositoryImpl{}
}

func (r *paymentRepositoryImpl) CreateCustomer(user *model.User) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *paymentRepositoryImpl) GetCustomerByUserId(userId int64) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *paymentRepositoryImpl) GetCreditCards(customerId string) ([]*model.Card, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *paymentRepositoryImpl) AddCreditCard(customerId string, cardToken string) (*model.Card, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *paymentRepositoryImpl) RemoveCreditCard(customerId string, cardId string) *model.AppError {
	return model.NewInternalServerError()
}

func (r *paymentRepositoryImpl) SetDefaultCreditCard(customerId string, cardId string) *model.AppError {
	return model.NewInternalServerError()
}

func (r *paymentRepositoryImpl) GetSubscriptionInfo(customerId string) (*model.Subscription, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *paymentRepositoryImpl) Subscribe(customerId string, priceId string) (*model.Subscription, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *paymentRepositoryImpl) CancelSubscription(customerId string) *model.AppError {
	return model.NewInternalServerError()
}
