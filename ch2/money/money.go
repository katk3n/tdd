package money

// Dollar is a struct to express dollar.
type Dollar struct {
	Amount int
}

// NewDollar is a constructor of Dollar.
func NewDollar(amount int) *Dollar {
	dollar := &Dollar{Amount: amount}
	return dollar
}

func (d *Dollar) times(multiplier int) *Dollar {
	return NewDollar(d.Amount * multiplier)
}
