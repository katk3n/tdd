package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewMoney(5, "USD")
	testCases := []struct {
		desc string
		in   int
		want *Money
	}{
		{"$5 * 2 = $10", 2, NewMoney(10, "USD")},
		{"$5 * 3 = $15", 3, NewMoney(15, "USD")},
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
	assert.True(t, NewMoney(5, "USD").Equals(NewMoney(5, "USD")))
	assert.False(t, NewMoney(5, "USD").Equals(NewMoney(6, "USD")))
	assert.True(t, NewMoney(5, "CHF").Equals(NewMoney(5, "CHF")))
	assert.False(t, NewMoney(5, "CHF").Equals(NewMoney(6, "CHF")))
	// assert.False(t, NewFranc(5).Equals(NewDollar(5)))  // Go does not implement extend
}

func TestFrancMultiplication(t *testing.T) {
	five := NewMoney(5, "CHF")
	testCases := []struct {
		desc string
		in   int
		want *Money
	}{
		{"5 CHF * 2 = 10 CHF", 2, NewMoney(10, "CHF")},
		{"5 CHF * 3 = 15 CHF", 3, NewMoney(15, "CHF")},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := five.Times(test.in); *got != *test.want {
				t.Errorf("%v: got %v want %v", test.desc, got, test.want)
			}
		})
	}
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1, "USD").Currency())
	assert.Equal(t, "CHF", NewFranc(1, "CHF").Currency())
}
