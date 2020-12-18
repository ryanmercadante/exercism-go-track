// Package raindrops converts and integer into rain
package raindrops

import "strconv"

// Convert takes an integer and checks if it has 3, 5, or 7 as factors,
// and adds rain drop noises if so. Otherwise it returns the string of
// the original integer
func Convert(x int) string {
	result := ""
	if x%3 == 0 {
		result += "Pling"
	}
	if x%5 == 0 {
		result += "Plang"
	}
	if x%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		return strconv.Itoa(x)
	}
	return result
}
