package main

// Real external service
type PaymentGateway interface {
	ProcessPayment(amount float64, cardToken string) (string, error)
}
