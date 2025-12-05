package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define mock object
type MockMessenger struct {
	mock.Mock
}

// Implement an interface using the mock
func (m *MockMessenger) Send(phoneNumber string, content string) bool {
	args := m.Called(phoneNumber, content)
	return args.Bool(0)
}

func TestProcessPayment_HighValue(t *testing.T) {
	mockObj := new(MockMessenger)

	mockObj.On("Send", "+177788800", "Payment processed").Return(true)

	processor := PaymentProcessor{Messenger: mockObj}

	result := processor.ProcessPayment(200, "+177788800")

	assert.Equal(t, "Payment Successful", result)

	mockObj.AssertExpectations(t)
}

func TestProcessPayment_LowValue(t *testing.T) {
	mockObj := new(MockMessenger)

	processor := PaymentProcessor{Messenger: mockObj}

	result := processor.ProcessPayment(50, "+15550000")

	assert.Equal(t, "Payment Successful", result)

	mockObj.AssertExpectations(t)
}
