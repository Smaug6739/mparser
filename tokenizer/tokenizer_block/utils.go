package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

// Param one: the string
// Param two: the left-trimmed string (strings.Trim(str, " "))
func countLeadingSpaces(str1, trimmedStr string) int {
	return len(str1) - len(trimmedStr)
}
func isSpaceOrTab(character rune) bool {
	return character == 32 /* space */ || character == 9 /* tab */
}
func isEmptyLine(str string) bool {
	return len(strings.Trim(str, " ")) == 0
}
func removeSpaces(str string, nbToRemove int) string {
	return str[nbToRemove:]
}

func insert(slice *preprocessor.Tokens, value preprocessor.Token, index int) {

	// Make space in the array for a new element. You can assign it any value.
	*slice = append(*slice, preprocessor.Token{})

	// Copy over elements sourced from index 2, into elements starting at index 3.
	copy((*slice)[index+1:], (*slice)[index:])

	(*slice)[index] = value
}
