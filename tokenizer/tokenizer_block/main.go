package tokenizer_block

import (
	"github.com/Smaug6739/mparser/internal/logger"
	"github.com/Smaug6739/mparser/preprocessor"
)

func New(state *preprocessor.Markdown) {
	for state.Tokens[len(state.Tokens)-1].Line < state.MaxIndex {
		logger.New().Warn("Loop tokenizer main")
		r := TokenizeBlock(state, 0, "paragraph")
		if !r {
			break
		}
	}
}

func TokenizeBlock(state *preprocessor.Markdown, skip int, end string) bool {
	if tokenizeEmpty(state, skip) {
		return true
	}
	if tokenizeList(state, skip) {
		return true
	}
	if tokenizeQuoteBlock(state, skip) {
		return true
	}
	if tokenizeIndentedCode(state, skip) {
		return true
	}
	if end == "paragraph" && tokenizeParagraph(state, skip) {
		return true
	} else if end == "inline" && tokenizeInline(state, skip) {
		return true
	}
	return false
}
