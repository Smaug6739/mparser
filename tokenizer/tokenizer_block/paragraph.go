package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

// TODO: Return bool
func tokenizeBlockParagraph(state *preprocessor.Markdown) {
	lastTokenIndex := len(state.Tokens) - 1
	lastToken := state.Tokens[lastTokenIndex]
	lineNumber := lastToken.Line + 1

	line := strings.Trim(state.Lines[lineNumber], " ")

	tokenStart := preprocessor.Token{
		Token: "paragraph_start",
		Html:  "<p>",
		Line:  lineNumber,
		Block: true,
	}
	tokenInline := preprocessor.Token{
		Token:   "inline",
		Content: line,
		Line:    lineNumber,
		Block:   false,
	}
	tokenClose := preprocessor.Token{
		Token: "paragraph_close",
		Html:  "</p>",
		Line:  lineNumber,
		Block: true,
	}
	//OPTIMIZATION: Repetition of len(line)
	if lastToken.Token == "paragraph_close" && len(line) > 0 {
		state.Tokens[lastTokenIndex] = tokenInline
		state.Tokens = append(state.Tokens, tokenClose)
	} else if len(line) > 0 {
		state.Tokens = append(state.Tokens, tokenStart)
		state.Tokens = append(state.Tokens, tokenInline)
		state.Tokens = append(state.Tokens, tokenClose)
	} else {
		state.Tokens = append(state.Tokens, preprocessor.Token{Token: "empty", Line: lineNumber})
	}
}
