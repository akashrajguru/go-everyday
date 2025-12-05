package main

import "fmt"

// The below example will mock the SMS service for payment processor app

// The Interface
// This is what we will mock. The PaymentProcessor doesn't care
// if it's a real SMS service or a fake one, as long as it has a Send() method.
type MessageService interface {
	Send(phoneNumber string, content string) bool
}

// The struct that performs the logic
type PaymentProcessor struct {
	Messenger MessageService
}

// Function which we will test
func (p *PaymentProcessor) ProcessPayment(amount int, phone string) string {
	if amount <= 0 {
		return "Invalid Amount"
	}
	if amount > 100 {
		success := p.Messenger.Send(phone, "Payment processed")
		if !success {
			return "Payment succeeded but SMS failed"
		}
	}
	return "Payment Successful"
}

func main() {
	fmt.Println("Run go test to see the mock in action.")
}
