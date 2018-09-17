package money

// Money is a struct to express money.
type Money struct {
	amount int
}

// NewMoney is a constructor of Money.
func NewMoney(amount int) *Money {
	return &Money{amount: amount}
}

// Times multiplies number to Money.
func (m *Money) Times(multiplier int) *Money {
	return NewMoney(m.amount * multiplier)
}

// Equals compares two Moneys.
func (m *Money) Equals(m2 *Money) bool {
	return m.amount == m2.amount
}

// NewDollar is a constructor of Dollar.
func NewDollar(amount int) *Dollar {
	dollar := &Dollar{Money{amount: amount}}
	return dollar
}

// NewFranc is a constructor of Franc.
func NewFranc(amount int) *Franc {
	dollar := &Franc{Money{amount: amount}}
	return dollar
}
