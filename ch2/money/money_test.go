package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	if product.Amount != 10 {
		t.Errorf("got %v, want %v", five.Amount, 10)
	}
	product = five.times(3)
	if product.Amount != 15 {
		t.Errorf("got %v, want %v", five.Amount, 15)
	}
}
