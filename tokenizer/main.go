package tokenizer

import (
	"github.com/Smaug6739/mparser/preprocessor"
	"github.com/Smaug6739/mparser/tokenizer/tokenizer_block"
)

func New(state *preprocessor.Markdown) {
	tokenizer_block.New(state)
}
