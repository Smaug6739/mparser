package tokenizer_block

import (
	"github.com/Smaug6739/mparser/internal/logger"
	"github.com/Smaug6739/mparser/preprocessor"
)

func New(state *preprocessor.Markdown) {
	for state.Tokens[len(state.Tokens)-1].Line < state.MaxIndex {
		logger.New().Warn("Loop tokenizer main")
		r := TokenizeBlock(state, 0)
		if !r {
			break
		}
	}
}

func TokenizeBlock(state *preprocessor.Markdown, skip int) bool {
	if tokenizeBlockEmpty(state, skip) {
		return true
	}
	return tokenizeBlockParagraph(state, skip)
}
