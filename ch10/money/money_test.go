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
	testCases := []struct {
		desc string
		in1  *Money
		in2  *Money
		want bool
	}{
		{"$5 == $5", NewMoney(5, "USD"), NewMoney(5, "USD"), true},
		{"$5 != $6", NewMoney(5, "USD"), NewMoney(6, "USD"), false},
		{"5 CHF == 5 CHF", NewMoney(5, "CHF"), NewMoney(5, "CHF"), true},
		{"5 CHF != 6 CHF", NewMoney(5, "CHF"), NewMoney(6, "CHF"), false},
		{"5 CHF != $5", NewMoney(5, "CHF"), NewMoney(5, "USD"), false},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := test.in1.Equals(test.in2); got != test.want {
				t.Errorf("%v: got %v want %v", test.desc, got, test.want)
			}
		})
	}
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

func TestDifferentClassEquality(t *testing.T) {
	assert.True(t, NewMoney(10, "CHF").Equals(NewMoney(10, "CHF")))
}
