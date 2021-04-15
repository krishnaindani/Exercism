package allergies

var products = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

//Allergies ...
func Allergies(score uint) []string {
	var product []string
	if v, ok := products[score]; ok {
		product = append(product, v)
	} else {
		evenValue := score / 2 
		oddValue := score % 2 
		
	}
	return product
}

//AllergicTo ...
func AllergicTo(score uint, substance string) bool {
	return false
}
