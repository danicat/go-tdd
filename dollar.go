package money

// Dollar holds money in the USD currency
type Dollar struct {
	Amount int
}

// Times multiply the value of Dollar by a given number
func (d *Dollar) Times(multiplier int) {
	d.Amount = d.Amount * multiplier
}
