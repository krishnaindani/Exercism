package bob

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var output string
	switch remark {
	case remark.IsUpper() == true:
		output = "Sure"
	case "WATCH OUT!":
		output = "Whoa, chill out!"
	default:
		output = "Whatever."
	}
	return output
}
