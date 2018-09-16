package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	if *five.Times(2) != *NewDollar(10) {
		t.Errorf("got %v, want %v", five.Times(2), 10)
	}
	if *five.Times(3) != *NewDollar(15) {
		t.Errorf("got %v, want %v", five.Times(3), 15)
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
