package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	if product.Amount != 10 {
		t.Errorf("got %v, want %v", five.Amount, 10)
	}
	product = five.Times(3)
	if product.Amount != 15 {
		t.Errorf("got %v, want %v", five.Amount, 15)
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5).Equals(NewDollar(5)) {
		t.Errorf("Two Dollars with same amounts must be equal.")
	}
	if NewDollar(5).Equals(NewDollar(6)) {
		t.Errorf("Two Dollars with different amounts must not be equal.")
	}
}
