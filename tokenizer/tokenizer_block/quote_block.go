package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeQuoteBlock(state *preprocessor.Markdown, options Options) bool {
	var open_quote bool = false

	data, ok := state.GetData(options.offset)
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
			max_delimiters := countDelimiters(content)
			if !inQuoteblock(&state.Tokens) {
				for i := index + 1; i <= state.MaxIndex && isQuote(state.Lines[i]); i++ {
					delimiters := countDelimiters(state.GetLine(i))
					if delimiters <= max_delimiters {
						state.Lines[i] = quoteOffset(state.GetLine(i), -1)
						if state.GetLine(i)[0] == ' ' {
							state.Lines[i] = removeFirstCharOfString(state.GetLine(i))
						}
					} else {
						state.Lines[i] = quoteOffset(state.Lines[i], max_delimiters-1) // -1 because it should have one more quote
						max_delimiters = delimiters
					}
				}
			}
			new_str := removeFirstCharOfString(content)
			state.Lines[index] = new_str
			if state.GetLine(index)[0] == ' ' {
				state.Lines[index] = removeFirstCharOfString(state.GetLine(index))
			}
			max_index := state.MaxIndex
			if !isQuote(state.GetLine(index + 1)) {
				max_index = index
			}

			openQuote(state, index, &open_quote)
			TokenizeBlock(state, Options{max_index: max_index}, "paragraph")
		} else {
			last_index_before := len(state.Tokens) - 1
			if TokenizeBlock(state, Options{max_index: state.MaxIndex}, "no_end") {
				insert(&state.Tokens, preprocessor.Token{Token: "blockquote_close", Html: "</blockquote>", Line: index, Block: true}, last_index_before+1)
				return true
			} else {
				TokenizeBlock(state, Options{max_index: state.MaxIndex}, "paragraph")
			}
		}
		index = state.GetLastToken().Line + 1
	}

	closeQuote(state, index-1, &open_quote)

	return true
}
func inQuoteblock(tokens *preprocessor.Tokens) bool {
	for i := len(*tokens) - 1; i >= 0; i-- {
		if (*tokens)[i].Token == "quote_block_close" {
			break
		}
		if (*tokens)[i].Token == "quote_block_open" {
			return true
		}
	}
	return false
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
func quoteOffset(str string, max int) string {
	quotes := 0
	for _, ch := range str {
		if ch == '>' && (quotes < max || max == -1) {
			str = removeFirstCharOfString(str)
			quotes++
		} else if ch == ' ' && quotes == 0 {
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
