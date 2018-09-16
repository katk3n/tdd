package money

// Franc is a struct to express dollar.
type Franc struct {
	Money
}

// NewFranc is a constructor of Franc.
func NewFranc(amount int) *Franc {
	franc := &Franc{Money{amount: amount}}
	return franc
}

// Times multiplies number to Franc.
func (d *Franc) Times(multiplier int) *Franc {
	return NewFranc(d.amount * multiplier)
}
