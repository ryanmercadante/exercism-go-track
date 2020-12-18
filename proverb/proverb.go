// Package proverb creates a proverb from an input slice
package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var proverb []string
	for i := 0; i < len(rhyme); i++ {
		if i+1 < len(rhyme) {
			line := "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
			proverb = append(proverb, line)
		} else {
			line := "And all for the want of a " + rhyme[0] + "."
			proverb = append(proverb, line)
		}
	}
	return proverb
}
