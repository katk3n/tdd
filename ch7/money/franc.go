package money

// Franc is a struct to express dollar.
type Franc struct {
	Money
}

// NewFranc is a constructor of Franc.
func NewFranc(amount int) *Franc {
	dollar := &Franc{Money{amount: amount}}
	return dollar
}

// Times multiplies number to Franc.
func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.amount * multiplier)
}

// Equals compares two Francs.
func (f *Franc) Equals(f2 *Franc) bool {
	return f.Money.Equals(&f2.Money)
}
