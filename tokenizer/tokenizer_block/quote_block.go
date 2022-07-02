package tokenizer_block

import (
	"fmt"
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeQuoteBlock(state *preprocessor.Markdown, offset int) bool {

	var open_quote bool = false

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
		if isEmptyLine(content) {
			break
		} else if is_quote {
			number_of_quotes := countDelimiters(content)
			new_str := removeFirstCharOfString(strings.Trim(content, " "))
			state.Lines[index] = new_str
			fmt.Println(new_str)
			for i := index + 1; i < state.MaxIndex && isQuote(state.Lines[i]); i++ {
				number_of_quotes_2 := countDelimiters(state.Lines[i])
				if number_of_quotes_2 < number_of_quotes {
					state.Lines[i] = quoteOffset(state.Lines[i])
				}
			}
			openQuote(state, index, &open_quote)
			TokenizeBlock(state, 0, "paragraph")
		} else {
			//TODO: Test
			if TokenizeBlock(state, 0, "no_end") {
				insert(&state.Tokens, preprocessor.Token{Token: "quote_block_close", Html: "</blockquote>", Line: index, Block: true}, index)
			} else {
				TokenizeBlock(state, 0, "paragraph")
			}
		}
		index = state.GetLastToken().Line + 1
	}

	closeQuote(state, index-1, &open_quote)

	return true
}

// Return true if the string is a block quote and false if not.
func isQuote(str string) bool {
	sep := false
	for _, ch := range str {
		if !sep && ch == ' ' {
			continue
		} else if !sep && ch == '>' {
			sep = true
			break
		} else {
			break
		}
	}
	return sep
}
func countDelimiters(str string) int {
	size := 0
	for _, ch := range str {
		if ch == ' ' {
			continue
		} else if ch == '>' {
			size++
		} else {
			break
		}
	}
	return size
}
func quoteOffset(str string) string {
	quotes := 0
	for _, ch := range str {
		if ch == '>' {
			str = removeFirstCharOfString(str)
			quotes++
		} else if ch == ' ' {
			str = removeFirstCharOfString(str)
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
