package money

import (
	"testing"
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

func TestCurrency(t *testing.T) {
	testCases := []struct {
		desc string
		in   *Money
		want string
	}{
		{"$1 is USD", NewDollar(1), "USD"},
		{"1 CHF is CHF", NewFranc(1), "CHF"},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := test.in.Currency(); got != test.want {
				t.Errorf("%v: got %v want %v", test.desc, got, test.want)
			}
		})
	}
}

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	var sum Expression
	sum = five.Plus(five)
	bank := NewBank()
	reduced := bank.reduce(&sum, "USD")
	if *reduced != *NewDollar(10) {
		t.Errorf("%v: got %v want %v", "reduced", NewDollar(10), reduced)
	}
}
