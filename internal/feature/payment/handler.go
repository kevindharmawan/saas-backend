package payment

type PaymentHandler struct {
	paymentService PaymentService
}

func NewPaymentHandler(paymentService PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}
