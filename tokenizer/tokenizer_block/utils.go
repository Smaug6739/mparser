package tokenizer_block

import (
	"errors"
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
func isNumber(character rune) bool {
	return character >= 48 && character <= 57
}
func isEmptyLine(str string) bool {
	return strings.Trim(str, " ") == ""
}
func removeSpaces(str string, nbToRemove int) string {
	return str[nbToRemove:]
}

type Data struct {
	lastTokenIndex int
	lastToken      *preprocessor.Token
	lineIndex      int
	lineContent    string
}

func getInfos(state *preprocessor.Markdown, offset int) (*Data, error) {
	lastTokenIndex := len(state.Tokens) - 1
	var lastToken preprocessor.Token
	for i := lastTokenIndex; i >= 0; i-- {
		if state.Tokens[i].Block {
			lastToken = state.Tokens[i]
			break
		}
	}
	lineIndex := lastToken.Line + 1
	if lineIndex > state.MaxIndex {
		return nil, errors.New("line number exceeds total of indexs")
	}
	lineContent := state.Lines[lineIndex]
	if len(lineContent) < offset {
		return nil, errors.New("offset exceeds line length")
	}
	return &Data{
		lastTokenIndex: lastTokenIndex,
		lastToken:      &lastToken,
		lineIndex:      lineIndex,
		lineContent:    lineContent[offset:],
	}, nil
}

// ------------------LIST--------------------
func ulItem(str string) (bool, int /* offset */, int /* blank_spaces */) {
	var delimiter string = ""
	var start_spaces int = 0
	for _, ch := range str {

		if ch == 45 /* - */ {
			delimiter = "-"
			break
		} else if delimiter == "" && ch == 32 /* space */ {
			start_spaces++
		} else if ch == 32 /* space */ {
			delimiter = "- "
		} else {
			break
		}
	}
	if delimiter != "" {
		return true, len(delimiter) + start_spaces, start_spaces
	}
	return false, 0, 0
}
func isUlItem(str string) bool {
	var delimiter string
	var spaces int
	for _, ch := range str {
		if ch == 45 /* - */ {
			delimiter = "-"
			break
		} else if delimiter == "" && ch == 32 /* space */ {
			spaces++
		} else if ch == 32 /* space */ {
			delimiter = "- "
		} else {
			break
		}
	}
	return delimiter != ""

}

func isOlItem(str string) (bool, int) {
	var delimiter, nb string
	var spaces int
	for _, ch := range str {
		if isNumber(ch) {
			nb += string(ch)
		}
		if ch == 41 /* ) */ {
			delimiter = ")"
			break
		}
		if ch == 46 /* . */ {
			delimiter = "."
			break
		}
		if ch == 32 /* space */ {
			spaces++
		}
	}
	if delimiter != "" && spaces > 0 {
		return true, len(nb) + len(delimiter) + spaces
	}
	return false, 0
}

func insert(slice *preprocessor.Tokens, value preprocessor.Token, index int) {

	// Make space in the array for a new element. You can assign it any value.
	*slice = append(*slice, preprocessor.Token{})

	// Copy over elements sourced from index 2, into elements starting at index 3.
	copy((*slice)[index+1:], (*slice)[index:])

	(*slice)[index] = value
}
