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

	r.name = generateName()
	for names[r.name] {
		r.name = generateName()
	}

	names[r.name] = true
	return r.name, nil

}

//Reset will reset the name
func (r *Robot) Reset() {
	r.name = ""
}

func checkName(name string) bool {
	if _, ok := names[name]; ok {
		return false
	}
	return true
}

func generateName() string {
	name := randLetters() + randNumbers()
	return name
}

//randLetters generated random letters
func randLetters() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVXWYZ")
	n := make([]rune, 2)
	for i := range n {
		n[i] = letters[rand.Intn(len(letters))]
	}
	return string(n)
}

//randNumbers generated ramdom number
func randNumbers() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}
