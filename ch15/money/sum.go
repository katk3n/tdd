package money

// Sum is a struct to express sum.
type Sum struct {
	augend Expression
	addend Expression
}

// NewSum creates new Sum instance.
func NewSum(augend Expression, addend Expression) *Sum {
	return &Sum{augend: augend, addend: addend}
}

// Reduce reduces sum.
func (s *Sum) Reduce(bank *Bank, to string) *Money {
	amount := s.augend.Reduce(bank, to).amount + s.addend.Reduce(bank, to).amount
	return NewMoney(amount, to)
}

// Plus need to be implemented.
func (s *Sum) Plus(addend Expression) Expression {
	return nil
}
