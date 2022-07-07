package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

// END: Choice between inline and paragraph
func tokenizeInline(state *preprocessor.Markdown, options Options) bool {
	data, ok := state.GetData(options.offset)
	if !ok {
		return false
	}
	lastRealToken := state.GetLastToken() // Block or inline item
	if lastRealToken.Token == "paragraph_close" {
		state.Tokens[data.LastTokenSliceIndex].Line = data.LineIndex
		insert(&state.Tokens, preprocessor.Token{
			Token:   "inline",
			Line:    data.LineIndex,
			Content: strings.Trim(data.LineContent, " "),
		}, data.LastTokenSliceIndex)
	} else {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token:   "inline",
			Content: strings.Trim(data.LineContent, " "),
			Line:    data.LineIndex,
			Closer:  true,
		})

	}
	return true
}
