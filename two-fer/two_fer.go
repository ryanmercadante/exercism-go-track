// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith takes a name as a string and returns a statement
// which includes the name, unless an empty string is provided,
// in which case name is replaced with 'you'
func ShareWith(name string) string {
	n := "you"
	if name != "" {
		n = name
	}
	return fmt.Sprintf("One for %s, one for me.", n)
}
