package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	five.Times(2)
	if five.Amount != 10 {
		t.Errorf("Expected: 10, got: %d", five.Amount)
	}
}
