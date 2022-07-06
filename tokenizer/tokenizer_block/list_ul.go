package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeList(state *preprocessor.Markdown, options Options) bool {

	var open_ul, open_li bool = false, false
	var first_leading_spaces int = -1
	var empty_lines int = 0
	// Get common informations
	data, ok := state.GetData(options.offset)
	if !ok {
		return false
	}

	// If the line is not a list, return false
	if !isUL(strings.TrimLeft(data.LineContent, " ")) {
		return false
	}

	// Handle other lines
	index := data.LineIndex
	for index <= options.max_index {

		content := state.GetLine(index)
		leading_spaces := countLeadingSpaces(content, strings.TrimLeft(content, " "))

		// First iteration
		if first_leading_spaces == -1 {
			// Save the first leading spaces & offset
			first_leading_spaces = leading_spaces
		}

		// If the line is not empty, reset the empty lines counter
		if !isEmptyLine(content) {
			empty_lines = 0
		}

		// If the line is empty, increment the empty lines counter and add "empty" token
		if isEmptyLine(content) {
			empty_lines++
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token:  "empty",
				Line:   index,
				Closer: true,
			})

			// If the line is not an item and the empty lines counter is greater than 0, close the list
		} else if empty_lines > 0 && !isUL(content) {
			break

			// If the line is a new item (offset of previous token is equal to the offset of the current line) AND the list is an item
		} else if leading_spaces == first_leading_spaces && isUL(content) {
			closeLI(state, index-1, &open_li) // -1 because the line the line is from the previous line
			openLI(state, index, &open_ul, &open_li)
			TokenizeBlock(state, Options{offset: leading_spaces + 2, max_index: state.MaxIndex}, "inline")

			// Next item with indentation (minumum previous offset spaces + 2 AND ) => New list
		} else if leading_spaces >= 2+options.offset && isUL(content) {
			tokenizeList(state, Options{offset: leading_spaces, max_index: state.MaxIndex})

			// If the line is indented before the offset, close the list
		} else if leading_spaces < 2+options.offset && isUL(content) {
			/*
			   If the line is a list but with previous indentation ()
			   - Item A (handled by ul1)
			     - Item B (handled by ul2)
			   - Item C <= Handle this case (handled by ul1 NOT ul2)
			*/
			break
		} else {
			last_index_before := len(state.Tokens) - 1

			// If the line is tokenized (not a paragraph/inline end)
			if TokenizeBlock(state, Options{offset: options.offset + leading_spaces}, "no_end") {
				insert(&state.Tokens, preprocessor.Token{Token: "li_close", Html: "</li>", Line: index, Closer: true}, last_index_before+1)
				insert(&state.Tokens, preprocessor.Token{Token: "ul_close", Html: "</ul>", Line: index, Closer: true}, last_index_before+2)
				return true

				// If the line before is empty and the actual line is not tokenized, close the list
			} else if state.GetLastToken().Token == "empty" {
				// Nothing to do (it should be a paragraph block)
				break

				// If the line is not tokenized and the previous line is not empty, add empty token
			} else {
				state.Tokens = append(state.Tokens, preprocessor.Token{Token: "inline", Content: content, Line: index, Closer: true})
			}
		}
		index = state.GetLastToken().Line + 1
	}
	closeLI(state, index-1, &open_li)
	closeUL(state, index-1, &open_ul)
	return true
}

func countListULOffset(str string) int {
	// The offset is :
	// - The number of spaces before the first character
	// - The length of the prefix
	// - The number of spaces after the first character
	offset := 0
	delimiter_len := 0
	for _, c := range str {
		if c == ' ' {
			offset++
		} else if c == '-' && delimiter_len == 0 {
			delimiter_len++
		} else {
			break
		}
	}
	return offset + delimiter_len
}

func isUL(str string) bool {
	left_trimed := strings.TrimLeft(str, " ")
	if len(left_trimed) >= 2 && left_trimed[0] == '-' && left_trimed[1] == ' ' {
		return true
	}
	return false
}
func openUL(state *preprocessor.Markdown, index int, open_ul *bool) {
	if !*open_ul {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token:  "ul_open",
			Html:   "<ul>",
			Line:   index,
			Closer: false,
		})
		*open_ul = true
	}
}
func closeUL(state *preprocessor.Markdown, index int, open_ul *bool) {
	if *open_ul {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token:  "ul_close",
			Html:   "</ul>",
			Line:   index,
			Closer: true,
		})
		*open_ul = false
	}
}
func openLI(state *preprocessor.Markdown, index int, open_ul, open_li *bool) {
	if !*open_ul {
		openUL(state, index, open_ul)
		*open_ul = true
	}
	if !*open_li {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token:  "li_open",
			Html:   "<li>",
			Line:   index,
			Closer: false,
		})
		*open_li = true
	}
}
func closeLI(state *preprocessor.Markdown, index int, open_li *bool) {
	if *open_li {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token:  "li_close",
			Html:   "</li>",
			Line:   index,
			Closer: true,
		})
		*open_li = false
	}
}
