package money

// Franc is a struct to express dollar.
type Franc struct {
	amount int
}

// NewFranc is a constructor of Franc.
func NewFranc(amount int) *Franc {
	dollar := &Franc{amount: amount}
	return dollar
}

// Times multiplies number to Franc.
func (d *Franc) Times(multiplier int) *Franc {
	return NewFranc(d.amount * multiplier)
}

// Equals compares two Francs.
func (d *Franc) Equals(o interface{}) bool {
	if value, ok := o.(*Franc); ok {
		return *d == *value
	}

	return false
}
