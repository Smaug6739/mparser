package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeIndentedCode(state *preprocessor.Markdown, options Options) bool {
	data, ok := state.GetData(options.offset)
	if !ok {
		return false
	}

	leftTrimmed := strings.TrimLeft(data.LineContent, " ")
	leadingSpaces := countLeadingSpaces(data.LineContent, leftTrimmed)
	if leadingSpaces < 4 {
		return false
	}
	if len(leftTrimmed) <= 0 {
		return false
	}
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token: "pre_start",
		Html:  "<pre>",
		Line:  data.LineIndex,
	})
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token: "code_start",
		Html:  "<code>",
		Line:  data.LineIndex,
	})
	var lastIndex int
	for index := data.LineIndex; index <= state.MaxIndex; index++ {
		content := removeSpaces(state.Lines[index], options.offset)
		if isEmptyLine(content) {
			break
		}
		leftTrimmed := strings.TrimLeft(content, " ")
		leadingSpaces := countLeadingSpaces(content, leftTrimmed)
		if leadingSpaces < 4 {
			break
		}
		lastIndex = index
		//TODO: Remove four spaces
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token:   "inline_code",
			Content: content,
			Line:    index,
		})
	}
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token:  "code_end",
		Html:   "</code>",
		Line:   lastIndex,
		Closer: true,
	})
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token:  "pre_end",
		Html:   "</pre>",
		Line:   lastIndex,
		Closer: true,
	})

	return false
}
