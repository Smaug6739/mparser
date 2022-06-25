package tokenizer_block

import (
	"strconv"
	"strings"

	"github.com/Smaug6739/mparser/internal/logger"
	"github.com/Smaug6739/mparser/preprocessor"
)

// TODO: Before paragraph
func tokenizeBlockLHeader(state *preprocessor.Markdown) bool {
	// Get common informations
	data, err := getInfos(state)
	if err != nil {
		return false
	}

	// If the string start by more than 3 spaces (it should be a code-block), returns
	leftTrimmed := strings.TrimLeft(data.lineContent, " ")
	leadingSpaces := countLeadingSpaces(data.lineContent, leftTrimmed)
	if leadingSpaces >= 4 {
		return false
	}

	if len(leftTrimmed) <= 0 {
		return false
	}

	prefix := leftTrimmed[0]
	var token, delimiter string
	if prefix == 61 /* = */ {
		token = "h1"
		delimiter = "==="
	} else if prefix == 45 /* - */ {
		token = "h2"
		delimiter = "---"
	} else {
		return false
	}

	for _, ch := range leftTrimmed {
		if ch != rune(prefix) {
			return false
		}
	}
	if data.lastToken.Token != "paragraph_close" {
		return false
	}
	var openTokenIndex int
	var closeTokenIndex int = data.lastTokenIndex
	for i := data.lastTokenIndex; i >= 0; i-- {
		logger.New().Warn("I: " + strconv.Itoa(i) + " Value: " + state.Tokens[i].Token)
		if state.Tokens[i].Token == "paragraph_open" {
			openTokenIndex = i
			break
		}
	}
	state.Tokens[openTokenIndex].Token = "lheader_open"
	state.Tokens[openTokenIndex].Html = "<" + token + ">"

	state.Tokens[closeTokenIndex].Token = "lheader_close"
	state.Tokens[closeTokenIndex].Html = "</" + token + ">"
	state.Tokens[closeTokenIndex].Markdown = delimiter

	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token:   "internal_lheading_delimiter",
		Content: delimiter,
		Line:    data.lineIndex,
	})

	// Get previous lines: All lines before and dont in a block except blank lines
	/*slice := []preprocessor.Token{}
	for previousLineIndex := data.lineIndex - 1; previousLineIndex >= data.lastToken.Line; previousLineIndex-- {
		lineContent := state.Lines[previousLineIndex]
		if isEmptyLine(lineContent) {
			break
		}
		slice = append(slice, preprocessor.Token{
			Content: lineContent,
			Line:    previousLineIndex,
		})
	}
	if len(slice) <= 0 {
		return false
	}
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token:    "lheader_start",
		Html:     `<` + token + `>`,
		Markdown: "---",
		Line:     slice[0].Line,
		Block:    true,
	})

	for i := len(slice); i > 0; i-- {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token:   "inline",
			Line:    slice[i-1].Line,
			Content: slice[i-1].Content,
			Block:   false,
		})
	}
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token: "lheader_end",
		Html:  `</` + token + `>`,
		Line:  slice[len(slice)-1].Line,
		Block: true,
	})*/
	return false
}
