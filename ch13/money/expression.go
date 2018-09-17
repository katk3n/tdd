package money

// Expression is an interface of expression.
type Expression interface {
	Reduce(to string) *Money
}
