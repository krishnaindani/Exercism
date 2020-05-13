package etl

import "strings"

//Transform transforms data from old format to new format
func Transform(input map[int][]string) map[string]int {
	newScore := make(map[string]int)

	for k, v := range input {
		for i := 0; i < len(v); i++ {
			s := strings.ToLower(v[i])
			newScore[s] = k
		}
	}

	return newScore
}
