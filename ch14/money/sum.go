package money

// Sum is a struct to express sum.
type Sum struct {
	augend Money
	addend Money
}

// NewSum creates new Sum instance.
func NewSum(augend *Money, addend *Money) *Sum {
	return &Sum{augend: *augend, addend: *addend}
}

// Reduce reduces sum.
func (s *Sum) Reduce(bank *Bank, to string) *Money {
	amount := s.augend.amount + s.addend.amount
	return NewMoney(amount, to)
}
