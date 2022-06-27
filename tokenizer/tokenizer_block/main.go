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
	r1 := tokenizeBlockHeader(state, skip)
	if r1 {
		return true
	}
	r2 := tokenizeBlockThematicBreak(state, skip)
	if r2 {
		return true
	}
	r3 := tokenizeBlockList(state, skip)
	if r3 {
		return true
	}
	r4 := tokenizeBlockIndentedCode(state, skip)
	if r4 {
		return true
	}
	r5 := tokenizeBlockParagraph(state, skip)

	r6 := tokenizeBlockLHeader(state, skip)
	if r6 {
		return true
	}

	// Paragraph at the end
	if r5 {
		return true
	}
	return false
}

func ListTokenizeBlock(state *preprocessor.Markdown, skip int) bool {
	r1 := tokenizeBlockHeader(state, skip)
	if r1 {
		return true
	}
	r3 := tokenizeBlockList(state, skip)
	if r3 {
		return true
	}
	r4 := tokenizeBlockIndentedCode(state, skip)
	if r4 {
		return true
	}
	r5 := tokenizeBlockParagraph(state, skip)

	r6 := tokenizeBlockLHeader(state, skip)
	if r6 {
		return true
	}

	// Paragraph at the end
	if r5 {
		return true
	}
	return false
}

func ListTokenizeBlock2(state *preprocessor.Markdown, skip int) bool {
	r1 := tokenizeBlockHeader(state, skip)
	if r1 {
		return true
	}
	r2 := tokenizeBlockList(state, skip)
	if r2 {
		return true
	}
	r3 := tokenizeBlockIndentedCode(state, skip)
	if r3 {
		return true
	}

	return false
}
