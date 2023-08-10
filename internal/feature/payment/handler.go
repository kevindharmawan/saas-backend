package payment

import "github.com/gin-gonic/gin"

type PaymentHandler struct {
	paymentService PaymentService
}

func NewPaymentHandler(paymentService PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

func (h *PaymentHandler) GetCreditCards(c *gin.Context) {
	return
}

func (h *PaymentHandler) AddCreditCard(c *gin.Context) {
	return
}

func (h *PaymentHandler) RemoveCreditCard(c *gin.Context) {
	return
}

func (h *PaymentHandler) SetDefaultCreditCard(c *gin.Context) {
	return
}

func (h *PaymentHandler) GetSubscriptionInfo(c *gin.Context) {
	return
}

func (h *PaymentHandler) Subscribe(c *gin.Context) {
	return
}

func (h *PaymentHandler) CancelSubscription(c *gin.Context) {
	return
}
