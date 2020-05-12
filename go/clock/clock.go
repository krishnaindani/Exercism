package clock

//Clock type
type Clock int

//New created a new clock
func New(hout, minute int) Clock {
	var c Clock
	return c
}

//String return representation of the clock
func (c *Clock) String() string {
	return ""
}

//Add adds minutes to the clock
func (c *Clock) Add(minutes int) Clock {
	var cl Clock
	return cl
}

//Subtract subtracts the minutes from clock
func (c *Clock) Subtract(minutes int) Clock {
	var cl Clock
	return cl
}
