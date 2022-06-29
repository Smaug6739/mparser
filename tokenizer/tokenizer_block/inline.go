package tokenizer_block

import (
	"github.com/Smaug6739/mparser/preprocessor"
)

// END: Choice between inline and paragraph
func tokenizeInline(state *preprocessor.Markdown, offset int) bool {
	data, ok := state.GetData(offset)
	if !ok {
		return false
	}
	lastRealToken := state.GetLastToken() // Block or inline item
	if lastRealToken.Token == "paragraph_close" {
		state.Tokens[data.LastTokenSliceIndex].Line = data.LineIndex
		insert(&state.Tokens, preprocessor.Token{
			Token:   "inline",
			Line:    data.LineIndex,
			Content: data.LineContent,
		}, data.LastTokenSliceIndex)
	} else {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token:   "inline",
			Content: data.LineContent,
			Line:    data.LineIndex,
			Block:   true,
		})

	}
	return true
}
