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
			if got := five.Times(test.in); *(got.(*Money)) != *test.want {
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
	reduced := bank.Reduce(sum, "USD")
	if *reduced != *NewDollar(10) {
		t.Errorf("%v: got %v want %v", "reduced", NewDollar(10), reduced)
	}
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	var result Expression
	result = five.Plus(five)
	var sum *Sum
	sum = result.(*Sum)
	if five != sum.augend {
		t.Errorf("%v: got %v want %v", "augend", five, sum.augend)
	}
	if five != sum.addend {
		t.Errorf("%v: got %v want %v", "addend", five, sum.addend)
	}
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	if *result != *NewDollar(7) {
		t.Errorf("%v: got %v want %v", "$3 + $4 = $7", result, NewDollar(7))
	}
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.Reduce(NewDollar(1), "USD")
	if *result != *NewDollar(1) {
		t.Errorf("%v: got %v want %v", "Reduce $1", result, NewDollar(1))
	}

}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(NewFranc(2), "USD")
	if *result != *NewDollar(1) {
		t.Errorf("%v: got %v want %v", "Reduce 2 CHF to USD", result, NewDollar(1))
	}
}

func TestIdentityRate(t *testing.T) {
	if NewBank().Rate("USD", "USD") != 1 {
		t.Errorf("%v: got %v want %v", "Rate from USD to USD", NewBank().Rate("USD", "USD"), 1)
	}
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
	if *result != *NewDollar(10) {
		t.Errorf("%v: got %v want %v", "$5 + 10CHF = $10", result, NewDollar(10))
	}
}
