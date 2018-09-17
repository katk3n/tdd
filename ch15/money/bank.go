package money

// Bank is a struct to express bank.
type Bank struct {
	rates map[Pair]int
}

// NewBank creates new Bank instance.
func NewBank() *Bank {
	return &Bank{rates: make(map[Pair]int)}
}

// Reduce reduces source.
func (b *Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(b, to)
}

// AddRate adds rate.
func (b *Bank) AddRate(from string, to string, rate int) {
	b.rates[Pair{from, to}] = rate
}

// Rate returns rate of "from" to "to"
func (b *Bank) Rate(from string, to string) int {
	if from == to {
		return 1
	}
	return b.rates[Pair{from, to}]
}
