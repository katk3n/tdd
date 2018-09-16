package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	if *five.Times(2) != *NewDollar(10) {
		t.Errorf("got %v, want %v", *five.Times(2), *NewDollar(10))
	}
	if *five.Times(3) != *NewDollar(15) {
		t.Errorf("got %v, want %v", *five.Times(3), *NewDollar(15))
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	if *five.Times(2) != *NewFranc(10) {
		t.Errorf("got %v, want %v", *five.Times(2), *NewFranc(10))
	}
	if *five.Times(3) != *NewFranc(15) {
		t.Errorf("got %v, want %v", *five.Times(3), *NewFranc(15))
	}
}
