package tokenizer_block

import (
	"github.com/Smaug6739/mparser/internal/logger"
	"github.com/Smaug6739/mparser/preprocessor"
)

func New(state *preprocessor.Markdown) {
	for state.Tokens[len(state.Tokens)-1].Line < state.TotalIndexes {
		logger.New().Warn("Loop tokenizer main")
		tokenizeBlockHeader(state)
		tokenizeBlockParagraph(state)
	}
}
