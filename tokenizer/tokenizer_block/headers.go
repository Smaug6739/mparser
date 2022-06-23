package tokenizer_block

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func countLeadingSpaces(str1, trimmedStr string) int {
	return len(str1) - len(trimmedStr)
}
func tokenizeBlockHeader(state *preprocessor.Markdown) {

	lastToken := state.Tokens[len(state.Tokens)-1]
	lineNumber := lastToken.Line + 1
	line := state.Lines[lineNumber]

	// If the string start by more than 3 spaces, returns
	leftTrimmed := strings.TrimLeft(line, " ")
	if countLeadingSpaces(line, leftTrimmed) >= 4 {
		return
	}

	if !strings.HasPrefix(leftTrimmed, "#") {
		fmt.Println("No header prefix")
		return
	}
	for level := 6; level >= 1; level-- {
		prefix := strings.Repeat("#", level) + " "
		if strings.HasPrefix(line, prefix) {
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token:    "header_start",
				Html:     `<h` + strconv.Itoa(level) + `>`,
				Markdown: prefix,
				Line:     lineNumber,
				Block:    true,
			})
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token:   "inline",
				Content: leftTrimmed[len(prefix):],
				Line:    lineNumber,
				Block:   false,
			})
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token: "header_end",
				Html:  `<h` + strconv.Itoa(level) + `/>`,
				Line:  lineNumber,
				Block: true,
			})

		}
	}
}
