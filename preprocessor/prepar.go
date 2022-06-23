package preprocessor

import (
	"strings"
)

func prepar(str string) *Markdown {
	// The main Markdown instance
	lines := strings.Split(str, "\n")
	var instance Markdown = Markdown{
		Source:   str,
		Lines:    lines,
		MaxIndex: len(lines) - 1,
		Tokens: Tokens{
			Token{
				Token:    "",
				Html:     "",
				Markdown: "",
				Line:     -1, // +1 = 0; so index match
				Block:    false,
			},
		},
	}
	return &instance
}
