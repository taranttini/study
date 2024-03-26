package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	// testando usando o modelo de if else sem importacoes
	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}

	// gera erro pois o valor nao esta ok
	if result != 6.0 {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
