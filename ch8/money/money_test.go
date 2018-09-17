package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewMoney(5)
	testCases := []struct {
		desc string
		in   int
		want *Money
	}{
		{"$5 * 2 = $10", 2, NewMoney(10)},
		{"$5 * 3 = $15", 3, NewMoney(15)},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := five.Times(test.in); *got != *test.want {
				t.Errorf("%v: got %v want %v", test.desc, got, test.want)
			}
		})
	}
}

func TestEquality(t *testing.T) {
	assert.True(t, NewMoney(5).Equals(NewMoney(5)))
	assert.False(t, NewMoney(5).Equals(NewMoney(6)))
	assert.True(t, NewMoney(5).Equals(NewMoney(5)))
	assert.False(t, NewMoney(5).Equals(NewMoney(6)))
	// assert.False(t, NewFranc(5).Equals(NewDollar(5)))  // Go does not implement extend
}

func TestFrancMultiplication(t *testing.T) {
	five := NewMoney(5)
	testCases := []struct {
		desc string
		in   int
		want *Money
	}{
		{"5 CHF * 2 = 10 CHF", 2, NewMoney(10)},
		{"5 CHF * 3 = 15 CHF", 3, NewMoney(15)},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := five.Times(test.in); *got != *test.want {
				t.Errorf("%v: got %v want %v", test.desc, got, test.want)
			}
		})
	}
}
