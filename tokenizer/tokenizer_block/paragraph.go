package tokenizer_block

import (
	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeBlockParagraph(state *preprocessor.Markdown, skip int) bool {
	data, ok := state.GetData()
	if !ok {
		return false
	}
	if data.LastToken.Token == "paragraph_close" {
		state.Tokens[data.LastTokenSliceIndex].Line = data.LineIndex
		insert(&state.Tokens, preprocessor.Token{
			Token:   "inline",
			Line:    data.LineIndex,
			Content: data.LineContent,
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
			Content: data.LineContent,
			Line:    data.LineIndex,
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
