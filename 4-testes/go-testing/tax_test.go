package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	testCases := []struct {
		amount   float64
		expected float64
	}{
		{0, 0},
		{500, 5},
		{1000, 10},
		{1500, 10},
	}

	for _, tc := range testCases {
		result := CalculateTax(tc.amount)
		if result != tc.expected {
			t.Errorf("Expected %f but got %f", tc.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

func BenchmarkCalculateTaxWithDelay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTaxWithDelay(500)
	}
}

func FuzzCalculateTax(f *testing.F) {
	f.Fuzz(func(t *testing.T, amount float64) {
		res := CalculateTax(amount)
		if amount <= 0 && res != 0 {
			t.Errorf("got %f but expected 0", res)
		}
		if amount > 20000 && res != 20 {
			t.Errorf("got %f but expected 20", res)
		}
	})
}
