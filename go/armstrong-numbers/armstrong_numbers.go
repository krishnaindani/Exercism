package armstrong

import (
	"math"
	"strconv"
	"strings"
)

//IsNumber ...
func IsNumber(number int) bool {
	s := strings.Split(strconv.Itoa(number), "")

	factor := len(s)
	var total float64
	for _, v := range s {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return false
		}
		total += math.Pow(n, float64(factor))
	}
	return number == int(total)
}
