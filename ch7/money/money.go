package money

// Money is a struct to express money.
type Money struct {
	amount int
}

// Equals compares two Moneys.
func (m1 *Money) Equals(m2 *Money) bool {
	return m1.amount == m2.amount
}
