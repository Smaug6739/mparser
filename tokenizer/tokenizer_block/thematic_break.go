package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeBlockThematicBreak(state *preprocessor.Markdown) bool {
	// Get common informations
	data, err := getInfos(state)
	if err != nil {
		return false
	}

	// If the string start by more than 3 spaces (it should be a code-block), returns
	leftTrimmed := strings.TrimLeft(data.lineContent, " ")
	if len(leftTrimmed) < 1 {
		return false
	}
	leadingSpaces := countLeadingSpaces(data.lineContent, leftTrimmed)
	if leadingSpaces >= 4 {
		return false
	}
	delimiter := rune(leftTrimmed[0])
	delimitersCount := 0
	if delimiter != 42 /* * */ && delimiter != 95 /* _ */ && delimiter != 45 /* - */ {
		return false
	}

	for _, character := range leftTrimmed {
		if isSpaceOrTab(character) {
			continue
		}
		if character != delimiter {
			return false
		}
		delimitersCount++
	}

	if delimitersCount < 3 {
		return false
	}
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token:    "thematic_break",
		Html:     "</hr>",
		Markdown: "---",
		Line:     data.lineIndex,
		Block:    true,
	})

	return false

}
