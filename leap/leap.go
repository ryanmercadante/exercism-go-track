// Package leap should have a package comment that summarizes what it's about.
package leap

// IsLeapYear that tells you if a year is a leap year
func IsLeapYear(x int) bool {
	if x%400 == 0 {
		return true
	} else if x%100 == 0 {
		return false
	} else if x%4 == 0 {
		return true
	}
	return false
}
