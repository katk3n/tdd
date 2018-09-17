package money

// Expression is an interface of expression.
type Expression interface {
	Plus(addend Expression) Expression
	Reduce(bank *Bank, to string) *Money
}
