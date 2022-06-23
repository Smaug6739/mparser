package tokenizer_block

import (
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
	leadingSpaces := countLeadingSpaces(line, leftTrimmed)
	if leadingSpaces >= 4 {
		return
	}

	for level := 6; level >= 1; level-- {
		prefix := strings.Repeat(" ", leadingSpaces) + strings.Repeat("#", level) // The markdown prefix

		if strings.HasPrefix(line, prefix) {
			content := line[len(prefix):] // The content based on prefix length
			if strings.HasPrefix(line, prefix+" ") {
				prefix += " "                // Increase the prefix with white space
				content = line[len(prefix):] // Update the content based on prefix
			} else if content != "" {
				return // If there has no space *and* the content is not empty, returns
			}

			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token:    "header_start",
				Html:     `<h` + strconv.Itoa(level) + `>`,
				Markdown: prefix,
				Line:     lineNumber,
				Block:    true,
			})
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token:   "inline",
				Content: strings.Trim(line[len(prefix):], " "),
				Line:    lineNumber,
				Block:   false,
			})
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token: "header_end",
				Html:  `</h` + strconv.Itoa(level) + `>`,
				Line:  lineNumber,
				Block: true,
			})

		}
	}
}
