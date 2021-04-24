package prime

func Factors(num int64) []int64 {

	if num == 1 {
		return []int64{}
	}

	var result []int64

	start := int64(2)
	for num > 1 {

		if num%start == 0 {
			num = num / start
			result = append(result, start)
		} else {
			start++
		}
	}

	return result
}
