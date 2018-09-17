package money

// Expression is an interface of expression.
type Expression interface {
	Reduce(bank *Bank, to string) *Money
}
