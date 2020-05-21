package clock

import "fmt"

//Clock type
type Clock int

//New created a new clock
func New(hour, minute int) Clock {
	c := (hour*60 + minute) % (24 * 60)
	if c < 0 {
		c += (24 * 60)
	}
	return Clock(c)
}

//String return representation of the clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

//Add adds minutes to the clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

//Subtract adds minutes to the clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)

}
