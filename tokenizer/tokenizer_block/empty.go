package tokenizer_block

import (
	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeEmpty(state *preprocessor.Markdown, options Options) bool {
	data, ok := state.GetData(options.offset)
	if !ok {
		return false
	}
	if isEmptyLine(data.LineContent) {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token:  "empty",
			Line:   data.LineIndex,
			Closer: true,
		})
		return true
	}
	return false
}
