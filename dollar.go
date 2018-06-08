package money

// Dollar holds money in the USD currency
type Dollar struct {
	amount int
}

// Times multiply the value of Dollar by a given number
func (d *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{d.amount * multiplier}
}

// Equals compare dollars with everything else
func (d *Dollar) Equals(i interface{}) bool {
	dollar, ok := i.(Dollar)
	return ok && d.amount == dollar.amount
}
