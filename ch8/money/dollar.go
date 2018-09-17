package money

// Dollar is a struct to express dollar.
type Dollar struct {
	Money
}

// Times multiplies number to Dollar.
func (d *Dollar) Times(multiplier int) *Money {
	return NewMoney(d.amount * multiplier)
}

// Equals compares two Dollars.
func (d *Dollar) Equals(d2 *Dollar) bool {
	return d.Money.Equals(&d2.Money)
}
