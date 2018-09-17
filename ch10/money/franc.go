package money

// Franc is a struct to express dollar.
type Franc struct {
	Money
}

// Equals compares two Francs.
func (f *Franc) Equals(f2 *Franc) bool {
	return f.Money.Equals(&f2.Money)
}

// Currency returns currency.
func (f *Franc) Currency() string {
	return f.Money.Currency()
}
