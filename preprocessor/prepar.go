package preprocessor

import (
	"strings"
)

func prepar(str string) Markdown {
	// The main Markdown instance
	var instance Markdown = Markdown{
		Source: str,
		Lines:  strings.Split(str, "\n"),
		Tokens: Tokens{},
	}
	return instance
}
