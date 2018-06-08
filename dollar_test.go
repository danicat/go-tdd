package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	product := five.Times(2)
	if product.Amount != 10 {
		t.Errorf("Expected: 10, got: %d", product.Amount)
	}
	product = five.Times(3)
	if product.Amount != 15 {
		t.Errorf("Expected: 15, got: %d", product.Amount)
	}
}
