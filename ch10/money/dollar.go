package money

// Dollar is a struct to express dollar.
type Dollar struct {
	Money
}

// Equals compares two Dollars.
func (d *Dollar) Equals(d2 *Dollar) bool {
	return d.Money.Equals(&d2.Money)
}

// Currency returns currency.
func (d *Dollar) Currency() string {
	return d.Money.Currency()
}
