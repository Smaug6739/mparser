package tokenizer_block

import (
	"strconv"
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeBlockHeader(state *preprocessor.Markdown) bool {

	// Get common informations
	data, err := getInfos(state)
	if err != nil {
		return false
	}

	// If the string start by more than 3 spaces (it should be a code-block), returns
	leftTrimmed := strings.TrimLeft(data.lineContent, " ")
	leadingSpaces := countLeadingSpaces(data.lineContent, leftTrimmed)
	if leadingSpaces >= 4 {
		return false
	}

	for level := 6; level >= 1; level-- {
		prefix := strings.Repeat(" ", leadingSpaces) + strings.Repeat("#", level) // The markdown prefix

		if strings.HasPrefix(data.lineContent, prefix) { // No space after # is allowed for empty headers
			content := data.lineContent[len(prefix):] // The content based on prefix length
			if strings.HasPrefix(data.lineContent, prefix+" ") {
				prefix += " "                            // Increase the prefix with white space
				content = data.lineContent[len(prefix):] // Update the content based on prefix
			} else if len(content) > 0 { // If there has no space *and* the content is not empty, returns
				return false
			}

			// All verifications pass, update state
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token:    "header_start",
				Html:     `<h` + strconv.Itoa(level) + `>`,
				Markdown: prefix,
				Line:     data.lineIndex,
				Block:    true,
			})
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token:   "inline",
				Content: strings.Trim(content, " "),
				Line:    data.lineIndex,
				Block:   false,
			})
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token: "header_end",
				Html:  `</h` + strconv.Itoa(level) + `>`,
				Line:  data.lineIndex,
				Block: true,
			})
			return true // The state was updated
		}
	}
	return false
}
