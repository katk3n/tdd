package money

// Bank is a struct to express bank.
type Bank struct {
}

// NewBank creates new Bank instance.
func NewBank() *Bank {
	return &Bank{}
}

// Reduce reduces source.
func (b *Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(to)
}
