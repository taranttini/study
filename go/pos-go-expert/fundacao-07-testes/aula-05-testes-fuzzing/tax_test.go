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
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)
		}
	}
}

// go test -bench=.
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)

	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500)

	}
}

// go test --fuzz=.
// go test --fuzz=. --run=^#
func FuzzCalculateTaxt(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f, but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Received %f, but expected 20", result)
		}
	})
}
