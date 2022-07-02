package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

// END: Choice between inline and paragraph
func tokenizeParagraph(state *preprocessor.Markdown, offset int) bool {
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
			Content: strings.Trim(data.LineContent, " "),
		}, data.LastTokenSliceIndex)
	} else {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token: "paragraph_open",
			Html:  "<p>",
			Line:  data.LineIndex,
			Block: true,
		})
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token:   "inline",
			Content: strings.Trim(data.LineContent, " "),
			Line:    data.LineIndex,
			Block:   true,
		})
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token: "paragraph_close",
			Html:  "</p>",
			Line:  data.LineIndex,
			Block: true,
		})
	}
	return true
}
