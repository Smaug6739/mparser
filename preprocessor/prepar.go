package preprocessor

import (
	"strings"
)

func prepar(str string) *Markdown {
	// The main Markdown instance
	var instance Markdown = Markdown{
		Source: str,
		Lines:  strings.Split(str, "\n"),
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
