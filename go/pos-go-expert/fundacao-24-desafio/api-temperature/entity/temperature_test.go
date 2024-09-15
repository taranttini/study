package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TemperatureValid(t *testing.T) {

	result := TemperatureResponse("sp", 30)

	expectResult := Temperature{
		City:       "sp",
		Celcius:    30,
		Fahrenheit: 86,
		Kelvin:     303.15,
	}
	assert.Equal(t, expectResult, result)

	assert.Equal(t, expectResult.Celcius, result.Celcius)
	assert.Equal(t, expectResult.Fahrenheit, result.Fahrenheit)
	assert.Equal(t, expectResult.Kelvin, result.Kelvin)
}

func Test_TemperatureRounding(t *testing.T) {

	result := TemperatureResponse("sp", 25.80)

	expectResult := Temperature{
		City:       "sp",
		Celcius:    25.8,
		Fahrenheit: 78.44,
		Kelvin:     298.95,
	}
	assert.Equal(t, expectResult, result)

	assert.Equal(t, expectResult.Celcius, result.Celcius)
	assert.Equal(t, expectResult.Fahrenheit, result.Fahrenheit)
	assert.Equal(t, expectResult.Kelvin, result.Kelvin)
}

func Test_TemperatureInvalid(t *testing.T) {

	result := TemperatureResponse("sp", 30)

	expectResult := Temperature{
		City:       "sp",
		Celcius:    30.1,
		Fahrenheit: 86.1,
		Kelvin:     303.16,
	}
	assert.NotEqual(t, expectResult, result)

	assert.NotEqual(t, expectResult.Celcius, result.Celcius)
	assert.NotEqual(t, expectResult.Fahrenheit, result.Fahrenheit)
	assert.NotEqual(t, expectResult.Kelvin, result.Kelvin)
}
