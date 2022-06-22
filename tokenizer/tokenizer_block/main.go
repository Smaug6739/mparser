package tokenizer_block

import "github.com/Smaug6739/mparser/preprocessor"

func New(state *preprocessor.Markdown) {
	tokenizeBlockHeader(state)
}
