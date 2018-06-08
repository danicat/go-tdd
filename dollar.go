package money

// Dollar holds money in the USD currency
type Dollar struct {
	Amount int
}

// Times multiply the value of Dollar by a given number
func (d *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{d.Amount * multiplier}
}

// Equals compare dollars with everything else
func (d *Dollar) Equals(i interface{}) bool {
	dollar, ok := i.(Dollar)
	return ok && d.Amount == dollar.Amount
}
