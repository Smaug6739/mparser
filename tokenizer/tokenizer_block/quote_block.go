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
			fmt.Println("Chargé en tant que quote: ", content)
			number_of_quotes := countDelimiters(content)
			if !inQuoteblock(&state.Tokens) {
				for i := index + 1; i <= state.MaxIndex && isQuote(state.Lines[i]); i++ {
					fmt.Println("Quote offset from:", content)
					number_of_quotes_2 := countDelimiters(state.Lines[i])
					if number_of_quotes_2 <= number_of_quotes {
						state.Lines[i] = quoteOffset(state.Lines[i], -1)
					} else {
						fmt.Println("Ici on doit trouver 4: ", state.Lines[i])
						state.Lines[i] = quoteOffset(state.Lines[i], number_of_quotes-1) // -1 because it should have one more quote
						number_of_quotes = number_of_quotes_2
						fmt.Println("A la sortie: ", state.Lines[i])
					}
				}
			}
			new_str := removeFirstCharOfString(strings.Trim(content, " "))
			state.Lines[index] = new_str
			openQuote(state, index, &open_quote)
			TokenizeBlock(state, 0, "paragraph")
		} else {
			//TODO: Test
			if TokenizeBlock(state, 0, "no_end") {
				insert(&state.Tokens, preprocessor.Token{Token: "quote_block_close", Html: "</blockquote>", Line: index, Block: true}, index)
				return true
			} else {
				TokenizeBlock(state, 0, "paragraph")
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
