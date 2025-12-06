package main

type WeatherService interface {
	GetTemperature(city string) (int, error)
}

func SuggestOutfit(ws WeatherService, city string) string {
	temp, err := ws.GetTemperature(city)
	if err != nil {
		return "Stay inside, cannot check weather."
	}
	if temp > 20 {
		return "Wear a T-shirt"
	}
	return "Wear a Coat"
}
