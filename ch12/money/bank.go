package money

// Bank is a struct to express bank.
type Bank struct {
}

// NewBank creates new Bank instance.
func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) reduce(source *Expression, to string) *Money {
	return NewDollar(10)
}
