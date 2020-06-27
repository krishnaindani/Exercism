package strain

//Ints is the slice of int
type Ints []int

//Lists ...
type Lists [][]int

//Strings ...
type Strings []string

//Keep will keep the slice of int based on function called
func (i Ints) Keep(f func(int) bool) Ints {
	var ints Ints

	for _, v := range i {
		if f(v) {
			ints = append(ints, v)
		}
	}
	return ints
}

//Discard will retrun list of the discarded ints based on condition
func (i Ints) Discard(f func(int) bool) Ints {
	var ints Ints

	for _, v := range i {
		if !f(v) {
			ints = append(ints, v)
		}
	}
	return ints
}

//Keep will keep if the function condition return true
func (l Lists) Keep(f func([]int) bool) Lists {
	var lists Lists

	for _, v := range l {
		if f(v) {
			lists = append(lists, v)
		}
	}
	return lists
}

//Keep will keep strings and return Strings if condition is true
func (s Strings) Keep(f func(string) bool) Strings {
	var strings Strings

	for _, v := range s {
		if f(v) {
			strings = append(strings, v)
		}
	}
	return strings
}
