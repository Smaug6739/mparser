package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeBlockIndentedCode(state *preprocessor.Markdown) bool {
	// Get common informations
	data, err := getInfos(state)
	if err != nil {
		return false
	}

	leftTrimmed := strings.TrimLeft(data.lineContent, " ")
	leadingSpaces := countLeadingSpaces(data.lineContent, leftTrimmed)
	if leadingSpaces < 4 {
		return false
	}
	if len(leftTrimmed) <= 0 {
		return false
	}
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token: "pre_start",
		Html:  "<pre>",
		Line:  data.lineIndex,
	})
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token: "code_start",
		Html:  "<code>",
		Line:  data.lineIndex,
	})
	var lastIndex int
	for index := data.lineIndex; index <= state.MaxIndex; index++ {
		content := state.Lines[index]
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
		Token: "code_end",
		Html:  "</code>",
		Line:  lastIndex,
	})
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token: "pre_end",
		Html:  "</pre>",
		Line:  lastIndex,
	})

	return false
}
