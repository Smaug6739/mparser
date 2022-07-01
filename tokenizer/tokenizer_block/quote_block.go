package tokenizer_block

import (
	"fmt"
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeQuoteBlock(state *preprocessor.Markdown, offset int) bool {

	var open_quote bool = false
	var max_quotes int = 0
	data, ok := state.GetData(offset)
	if !ok {
		return false
	}
	// If the line is not a block quote, return false
	is_quote := isQuote(data.LineContent)
	if !is_quote {
		return false
	}
	index := data.LineIndex
	for index <= state.MaxIndex {
		content := state.Lines[index]
		is_quote := isQuote(strings.TrimLeft(content, " "))
		new_str := quoteOffset(content)
		delimiter_size := countLeadingSpaces(content, new_str)
		fmt.Println("Line:", content, "is quote ?", is_quote)
		fmt.Println("Quote size: ", delimiter_size)
		if isEmptyLine(content) {
			break
		} else if is_quote {
			state.Lines[index] = new_str
			if max_quotes < delimiter_size {
				max_quotes = delimiter_size
				openQuote(state, index, &open_quote)
			}
			TokenizeBlock(state, 0, "paragraph")
		} else {
			TokenizeBlock(state, 0, "paragraph")
		}
		index = state.GetLastToken().Line + 1
	}
	closeQuote(state, index-1, &open_quote)
	return true
}

// Return true if the string is a block quote and false if not.
func isQuote(str string) bool {
	if !strings.HasPrefix(str, ">") {
		return false
	}
	return true
}

// Return also number of quote (>)
func quoteOffset(str string) string {
	separation_char := false
	for _, ch := range str {
		if !separation_char && ch == '>' {
			if !separation_char {
				str = removeFirstCharOfString(str)
				separation_char = true
			}
		} else {
			break
		}
	}
	return str
}

func openQuote(state *preprocessor.Markdown, index int, open_quote *bool) {
	if !*open_quote {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token: "quote_block_open",
			Html:  "<blockquote>",
			Line:  index,
			Block: false,
		})
		*open_quote = true
	}
}
func closeQuote(state *preprocessor.Markdown, index int, open_quote *bool) {
	if *open_quote {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token: "quote_block_close",
			Html:  "</blockquote>",
			Line:  index,
			Block: true,
		})
		*open_quote = false
	}
}
