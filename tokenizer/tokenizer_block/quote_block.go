package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/internal/logger"
	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeQuoteBlock(state *preprocessor.Markdown, offset int) bool {

	var open_quote bool = false
	data, ok := state.GetData(offset)
	if !ok {
		return false
	}
	// If the line is not a block quote, return false
	if !isQuote(strings.TrimLeft(data.LineContent, " ")) {
		return false
	}
	index := data.LineIndex
	for index <= state.MaxIndex {
		content := state.Lines[index]

		if isEmptyLine(content) {
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token: "empty",
				Line:  index,
				Block: true,
			})
			break
		} else if isQuote(content) {
			openQuote(state, index, &open_quote)
			TokenizeBlock(state, 2, "paragraph")
		} else {
			logger.New().Error("Else")
		}
		index = state.GetLastToken().Line + 1
	}
	closeQuote(state, index-1, &open_quote)
	return true
}

func isQuote(str string) bool {
	left_trimed := strings.TrimLeft(str, " ")
	if len(left_trimed) >= 2 && left_trimed[0] == '>' && left_trimed[1] == ' ' {
		return true
	}
	return false
}

func openQuote(state *preprocessor.Markdown, index int, open_quote *bool) {
	if !*open_quote {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token: "quote_block_open",
			Html:  "<blockquote>",
			Line:  index,
			Block: false,
		})
		*open_quote = true
	}
}
func closeQuote(state *preprocessor.Markdown, index int, open_quote *bool) {
	if *open_quote {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token: "quote_block_close",
			Html:  "</blockquote>",
			Line:  index,
			Block: true,
		})
		*open_quote = false
	}
}
