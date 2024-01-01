package utils

import "regexp"

func isValidFormatEmp(input string) bool {
	// Define the regular expression pattern
	pattern := `^IB\d+$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Use the regular expression to match the input string
	return regex.MatchString(input)
}
