package money

// Expression is an interface of expression.
type Expression interface {
	Times(multiplier int) Expression
	Plus(addend Expression) Expression
	Reduce(bank *Bank, to string) *Money
}
