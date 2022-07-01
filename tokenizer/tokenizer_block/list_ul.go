package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

func tokenizeList(state *preprocessor.Markdown, offset int) bool {

	var open_ul, open_li bool = false, false
	var first_start_spaces int = -1
	var empty_lines int = 0
	data, ok := state.GetData(offset)
	if !ok {
		return false
	}
	// If the line is not a list, return false
	if !isUL(strings.TrimLeft(data.LineContent, " ")) {
		return false
	}
	index := data.LineIndex
	for index <= state.MaxIndex {
		content := state.Lines[index]
		leading_spaces := countLeadingSpaces(content, strings.TrimLeft(content, " "))
		if first_start_spaces == -1 { // First iteration
			first_start_spaces = leading_spaces // Save the first leading spaces
		}
		if !isEmptyLine(content) {
			empty_lines = 0
		}
		if isEmptyLine(content) {
			empty_lines++
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token: "empty",
				Line:  index,
				Block: true,
			})
			//break // End of list
		} else if empty_lines > 0 && !isUL(content) {
			break
		} else if leading_spaces == first_start_spaces && isUL(content) { // Next item without indentation
			closeLI(state, index-1, &open_li) // -1 because the line the line is from the previous line
			openLI(state, index, &open_ul, &open_li)
			TokenizeBlock(state, leading_spaces+2, "inline")
		} else if leading_spaces >= 2+offset && isUL(content) { // Next item with indentation (minumum two spaces) (new list)
			tokenizeList(state, leading_spaces)
		} else if leading_spaces < 2+offset && isUL(content) {
			/*
				If the line is a list but with previous indentation ()
				- Item A (handled by ul1)
				  - Item B (handled by ul2)
				- Item C <= Handle this case (handled by ul1 NOT ul2)
			*/
			break
		} else {
			//TODO: Test this
			r := TokenizeBlock(state, offset+leading_spaces, "no_end")
			if r {
				insert(&state.Tokens, preprocessor.Token{Token: "li_close", Html: "</li>", Line: index, Block: true}, index)
				insert(&state.Tokens, preprocessor.Token{Token: "ul_close", Html: "</ul>", Line: index, Block: true}, index)
				return false // Return true TODO
			} else if state.GetLastToken().Token == "empty" {
				// Nothing to do (it should be a paragraph block)
				break
			} else {
				state.Tokens = append(state.Tokens, preprocessor.Token{Token: "inline", Content: content, Line: index, Block: true})
			}
		}
		index = state.GetLastToken().Line + 1
	}
	closeLI(state, index-1, &open_li)
	closeUL(state, index-1, &open_ul)
	return true
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
			Token: "ul_open",
			Html:  "<ul>",
			Line:  index,
			Block: false,
		})
		*open_ul = true
	}
}
func closeUL(state *preprocessor.Markdown, index int, open_ul *bool) {
	if *open_ul {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token: "ul_close",
			Html:  "</ul>",
			Line:  index,
			Block: true,
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
			Token: "li_open",
			Html:  "<li>",
			Line:  index,
			Block: false,
		})
		*open_li = true
	}
}
func closeLI(state *preprocessor.Markdown, index int, open_li *bool) {
	if *open_li {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token: "li_close",
			Html:  "</li>",
			Line:  index,
			Block: true,
		})
		*open_li = false
	}
}
