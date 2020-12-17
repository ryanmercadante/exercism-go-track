// Package hamming has a function, Distance, which calculates
// the Hamming Distance between two DNA strands.
package hamming

import (
	"errors"
)

// Distance is a function which calcuates the Hamming Distance
// between two DNA strands. It takes two strings, and counts the
// number of differences between the letters. The Hamming Distance
// can only be calculated if the strings are the same length,
// so we must include proper error handling
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the lengths of the string must be equal")
	}

	// don't need to check length of b since we already checked for equal length above
	if len(a) == 0 {
		return 0, nil
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
