package tokenizer_block

import (
	"github.com/Smaug6739/mparser/preprocessor"
)

func New(state *preprocessor.Markdown) {
	for state.Tokens[len(state.Tokens)-1].Line < state.MaxIndex {
		r := TokenizeBlock(state, Options{max_index: state.MaxIndex}, "paragraph")
		if !r {
			break
		}
	}
}

type Options struct {
	offset      int
	max_index   int
	must_prefix string
}

func TokenizeBlock(state *preprocessor.Markdown, options Options, end string) bool {
	if tokenizeEmpty(state, options) {
		return true
	}
	/*if tokenizeQuoteBlock(state, options) {
		return true
	}*/
	if tokenizeList(state, options) {
		return true
	}
	if tokenizeIndentedCode(state, options) {
		return true
	}
	if end == "paragraph" && tokenizeParagraph(state, options) {
		return true
	} else if end == "inline" && tokenizeInline(state, options) {
		return true
	}
	return false
}
