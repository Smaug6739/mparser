package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeBlockParagraph(state *preprocessor.Markdown, skip int) bool {
	data, err := getInfos(state, skip)
	if err != nil {
		return false
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
		Content: strings.Trim(data.lineContent, " "),
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
	if state.Tokens[data.lastTokenIndex].Token == "paragraph_close" && data.lastToken.Line == data.lineIndex-1 && len(data.lineContent) > 0 {
		state.Tokens[data.lastTokenIndex] = tokenInline
		insert(&state.Tokens, tokenClose, data.lastTokenIndex+1)
	} else if len(data.lineContent) > 0 {
		state.Tokens = append(state.Tokens, tokenStart)
		state.Tokens = append(state.Tokens, tokenInline)
		state.Tokens = append(state.Tokens, tokenClose)
	} else {
		state.Tokens = append(state.Tokens, preprocessor.Token{Token: "empty", Line: data.lineIndex, Block: true})
	}
	return true
}
