package tokenizer_block

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeBlockHeader(state *preprocessor.Markdown) {
	lastToken := state.Tokens[len(state.Tokens)-1]
	lineNumber := lastToken.Line + 1
	line := state.Lines[lineNumber]
	fmt.Println("Last token: ", lastToken)
	fmt.Println("Line to parse: ", line)
	if !strings.HasPrefix(line, "#") {
		fmt.Println("No header prefix")
		return
	}
	for level := 6; level >= 1; level-- {
		prefix := strings.Repeat("#", level) + " "
		if strings.HasPrefix(line, prefix) {
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token:    "header_start",
				Html:     `<h` + strconv.Itoa(level) + `>`,
				Markdown: prefix,
				Line:     lineNumber,
				Block:    true,
			})
		}
	}
}
