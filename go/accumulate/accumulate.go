package accumulate

//Accumulate ...
func Accumulate(words []string, operation func(x string) string) []string {

	var output []string

	for _, word := range words {
		output = append(output, operation(word))
	}
	return output
}
