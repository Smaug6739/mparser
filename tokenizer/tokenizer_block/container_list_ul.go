package tokenizer_block

import (
	"fmt"
	"strings"

	"github.com/Smaug6739/mparser/internal/logger"
	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeList(state *preprocessor.Markdown, options Options) bool {
	data, ok := state.GetData(options.offset)
	// VERIFICATIONS
	if !ok {
		return false
	}
	if !isUL(data.LineContent) {
		return false
	}
	if countLeadingSpaces(data.LineContent, strings.TrimLeft(data.LineContent, " ")) >= 4 {
		return false
	}
	// #END VERIFICATIONS
	var index int = data.LineIndex           // The index of the current line
	var empty_lines int = 0                  // The number of empty lines
	var open_ul, open_li bool = false, false // The state of the UL

	// META-DATA:
	var normal_leading_spaces int = 0 // The minimum number of leading spaces
	var last_leading_spaces int = -1
	var last_leading_offset int = -1
	// #END META-DATA
	for index <= options.max_index {
		line_content := state.GetLine(index)
		line_leading_spaces := len(line_content) - len(strings.TrimLeft(line_content, " "))

		if last_leading_spaces == -1 && last_leading_offset == -1 {
			normal_leading_spaces = line_leading_spaces
			last_leading_spaces = line_leading_spaces
			last_leading_offset = countListULOffset(line_content)
		}
		if line_content == "bim" {
			fmt.Println(0)
		}
		if isEmptyLine(line_content) {
			if line_content == "bim" {
				fmt.Println(1)
			}
			empty_lines++
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token:  "empty",
				Line:   index,
				Closer: true,
			})
			// NEW ITEM (LI) IN THE LIST
		} else if isUL(line_content) && (normal_leading_spaces == line_leading_spaces || normal_leading_spaces == line_leading_spaces-1 || (line_leading_spaces < normal_leading_spaces && countOpensBlocks(state, "ul_open", "ul_close") == 1)) {
			closeLI(state, index-1, &open_ul, &open_li) // -1 because the line the line is from the previous line
			openLI(state, index, &open_ul, &open_li)
			TokenizeBlock(state, Options{offset: line_leading_spaces + 2, max_index: state.MaxIndex}, "inline")
			if line_content == "bim" {
				fmt.Println(2)
			}
			// NEW LIST (UL) IN THE LIST
		} else if isUL(line_content) && line_leading_spaces >= last_leading_spaces+2 && line_leading_spaces >= last_leading_offset && line_leading_spaces-last_leading_offset <= 4 {
			tokenizeList(state, Options{offset: line_leading_spaces, max_index: state.MaxIndex})
			if line_content == "bim" {
				fmt.Println(3)
			}
		} else if line_leading_spaces < last_leading_spaces && countOpensBlocks(state, "ul_open", "ul_close") > 1 {
			if line_content == "bim" {
				fmt.Println(4)
			}
			break // END OF INDENTED LIST
		} else {
			if line_content == "bim" {
				fmt.Println(5)
			}
			var slice_index_before int = len(state.Tokens) - 1
			if TokenizeBlock(state, Options{offset: line_leading_spaces, max_index: state.MaxIndex}, "no_end") || state.GetLastToken().Token == "empty" {
				insert(&state.Tokens, preprocessor.Token{Token: "li_close", Html: "</li>", Line: index - 1, Closer: true}, slice_index_before+1)
				insert(&state.Tokens, preprocessor.Token{Token: "ul_close", Html: "</ul>", Line: index - 1, Closer: true}, slice_index_before+2)
				return true
			} else {
				if state.GetLastToken().Token == "ul_close" {
					logger.New().Details(state.Tokens)
					insert(&state.Tokens, preprocessor.Token{Token: "inline", Content: line_content, Line: index}, slice_index_before-1)
					state.Tokens[state.GetLastTokenSliceIndex()].Line = index
				} else {
					TokenizeBlock(state, Options{offset: line_leading_spaces, max_index: state.MaxIndex}, "inline")
				}
			}
		}
		last_leading_spaces = line_leading_spaces
		last_leading_offset = countListULOffset(line_content)
		index = state.GetLastToken().Line + 1
	}
	closeLI(state, index-1, &open_ul, &open_li)
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
func closeLI(state *preprocessor.Markdown, index int, open_ul, open_li *bool) {
	if !*open_ul {
		return
	}
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
