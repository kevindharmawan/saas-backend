package payment

import "github.com/kevindharmawan/saas-backend/internal/model"

type paymentServiceImpl struct {
	paymentRepository PaymentRepository
}

func NewPaymentService(paymentRepository PaymentRepository) PaymentService {
	return &paymentServiceImpl{
		paymentRepository: paymentRepository,
	}
}

func (s *paymentServiceImpl) GetCreditCards(user *model.User) ([]*model.Card, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (s *paymentServiceImpl) AddCreditCard(user *model.User, cardToken string) (*model.Card, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (s *paymentServiceImpl) RemoveCreditCard(user *model.User, cardId string) *model.AppError {
	return model.NewInternalServerError()
}

func (s *paymentServiceImpl) SetDefaultCreditCard(user *model.User, cardId string) *model.AppError {
	return model.NewInternalServerError()
}

func (s *paymentServiceImpl) GetSubscriptionInfo(user *model.User) (*model.Subscription, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (s *paymentServiceImpl) Subscribe(user *model.User, priceId string) (*model.Subscription, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (s *paymentServiceImpl) CancelSubscription(user *model.User) *model.AppError {
	return model.NewInternalServerError()
}
