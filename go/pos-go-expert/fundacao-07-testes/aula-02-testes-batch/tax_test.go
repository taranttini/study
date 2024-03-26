package tax

import (
	"testing"
)

// run go test .
func TestCalculateTax(t *testing.T) {
	// arrange
	// act
	// assert
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	// testando usando o modelo de if else sem importacoes
	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

// run go test -v

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)
		}
	}
}
