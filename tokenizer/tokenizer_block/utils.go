package tokenizer_block

import (
	"errors"

	"github.com/Smaug6739/mparser/preprocessor"
)

func countLeadingSpaces(str1, trimmedStr string) int {
	return len(str1) - len(trimmedStr)
}

type Data struct {
	lastTokenIndex int
	lastToken      preprocessor.Token
	lineIndex      int
	lineContent    string
}

func getInfos(state preprocessor.Markdown) (*Data, error) {
	lastTokenIndex := len(state.Tokens) - 1
	lastToken := state.Tokens[lastTokenIndex]
	lineIndex := lastToken.Line + 1
	if lineIndex > state.MaxIndex {
		return nil, errors.New("Line number exceeds total of indexs")
	}
	lineContent := state.Lines[lineIndex]
	return &Data{
		lastTokenIndex: lastTokenIndex,
		lastToken:      lastToken,
		lineIndex:      lineIndex,
		lineContent:    lineContent,
	}, nil
}
