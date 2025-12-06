package main

import (
	"mockery-demo/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuggestOutfit_Cold(t *testing.T) {
	mockService := mocks.NewWeatherService(t)
	mockService.On("GetTemperature", "London").Return(10, nil)
	result := SuggestOutfit(mockService, "London")
	assert.Equal(t, "Wear a Coat", result)
}

func TestSuggestOutfit_Warm(t *testing.T) {
	mockService := mocks.NewWeatherService(t)
	mockService.On("GetTemperature", "Miami").Return(30, nil)
	result := SuggestOutfit(mockService, "Miami")
	assert.Equal(t, "Wear a T-shirt", result)
}
