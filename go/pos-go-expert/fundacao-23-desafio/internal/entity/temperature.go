package entity

import "math"

type Temperature struct {
	Celcius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func TemperatureResponse(celcius float64) *Temperature {
	return &Temperature{
		Celcius:    celcius,
		Kelvin:     convertCelciusToKelvin(celcius),
		Fahrenheit: convertCelciusToFahrenheit(celcius),
	}
}

func convertCelciusToKelvin(celcius float64) float64 {
	kelvinBase := float64(273.15)

	return math.Round((celcius+kelvinBase)*100) / 100
}

func convertCelciusToFahrenheit(celcius float64) float64 {
	fahrenheitBase := 32
	fahrenheitMultiplier := float64(1.8)

	return math.Round((celcius*fahrenheitMultiplier+float64(fahrenheitBase))*100) / 100
}
