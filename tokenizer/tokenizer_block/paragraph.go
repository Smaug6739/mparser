package tokenizer_block

import (
	"github.com/Smaug6739/mparser/preprocessor"
	"strings"
)

// TODO: Return bool
func tokenizeBlockParagraph(state *preprocessor.Markdown) {
	data, err := getInfos(state)
	if err != nil {
		return
	}
	data.lineContent = strings.Trim(data.lineContent, " ")

	tokenStart := preprocessor.Token{
		Token: "paragraph_open",
		Html:  "<p>",
		Line:  data.lineIndex,
		Block: true,
	}
	tokenInline := preprocessor.Token{
		Token:   "inline",
		Content: data.lineContent,
		Line:    data.lineIndex,
		Block:   false,
	}
	tokenClose := preprocessor.Token{
		Token: "paragraph_close",
		Html:  "</p>",
		Line:  data.lineIndex,
		Block: true,
	}
	//OPTIMIZATION: Repetition of len(line)
	if data.lastToken.Token == "paragraph_close" && data.lastToken.Line == data.lineIndex-1 && len(data.lineContent) > 0 {
		state.Tokens[data.lastTokenIndex] = tokenInline
		state.Tokens = append(state.Tokens, tokenClose)
	} else if len(data.lineContent) > 0 {
		state.Tokens = append(state.Tokens, tokenStart)
		state.Tokens = append(state.Tokens, tokenInline)
		state.Tokens = append(state.Tokens, tokenClose)
	} else {
		state.Tokens = append(state.Tokens, preprocessor.Token{Token: "empty", Line: data.lineIndex})
	}
}
