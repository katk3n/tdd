package money

// Pair is a struct to express currency rate
type Pair struct {
	from string
	to   string
}

// Equals compares two Pairs.
func (p *Pair) Equals(o *Pair) bool {
	return p.from == o.from && p.to == o.to
}

// HashCode returns hash code.
func (p *Pair) HashCode() int {
	return 0
}
