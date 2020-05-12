package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Robot is name
type Robot struct {
	name string
}

var names = make(map[string]bool)

//init function
func init() {
	rand.Seed(time.Now().UnixNano())
}

//Name returns new name
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(names) > 26*26*10*10*10 {
		return "", errors.New("error: names exhausted")
	}

	r.name = newName()
	for names[r.name] {
		r.name = newName()
	}

	names[r.name] = true
	return r.name, nil

}

//Reset will reset the name
func (r *Robot) Reset() {
	r.name = ""
}

func newName() string {
	r1 := rand.Intn(26) + 'A'
	r2 := rand.Intn(26) + 'A'
	num := rand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}
