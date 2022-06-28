package tokenizer_block

import (
	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeBlockEmpty(state *preprocessor.Markdown, skip int) bool {
	data, ok := state.GetData()
	if !ok {
		return false
	}
	if isEmptyLine(data.LineContent) {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token: "empty",
			Line:  data.LineIndex,
			Block: true,
		})
		return true
	}
	return false
}
