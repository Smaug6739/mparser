package tokenizer_block

import (
	"fmt"
	"strconv"
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

func removeIndex(a *[]any, i int) {
	// Remove the element at index i from a.
	(*a)[i] = (*a)[len(*a)-1] // Copy last element to index i.
	(*a)[len(*a)-1] = ""      // Erase last element (write zero value).
	*a = (*a)[:len(*a)-1]     // Truncate slice.
}
func removeCharOfString(str string, index int) string {
	a := str[:index] // First part
	b := str[index+1:]

	fmt.Println("REMOVE_CHAR_OF_STRING: ", strconv.Quote(str), "INDEX:", index, "A:", a, "B:", b)
	return a + b
}
func removeFirstCharOfString(str string) string {
	return str[1:]
}
func removePrefix(str string, nb int) string {
	return str[nb:]
}
func countOpensBlocks(state *preprocessor.Markdown, open, close string) int {
	var open_blocks int = 0
	var close_blocks int = 0
	for _, token := range state.Tokens {
		if token.Token == open {
			open_blocks++
		} else if token.Token == close {
			close_blocks++
		}
	}
	return open_blocks - close_blocks
}
