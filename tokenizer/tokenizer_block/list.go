package tokenizer_block

import (
	"fmt"
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

var ul_call int = 0

func tokenizeList(state *preprocessor.Markdown, offset int) bool {

	var open_ul, open_li bool = false, false
	var first_start_spaces int = -1
	data, ok := state.GetData(offset)
	if !ok {
		return false
	}
	// If the line is not a list, return false
	if !isUL(strings.TrimLeft(data.LineContent, " ")) {
		return false
	}
	ul_call++
	index := data.LineIndex
	for index <= state.MaxIndex {
		content := state.Lines[index]
		fmt.Println("content:", content, "ul_call:", ul_call)
		leading_spaces := countLeadingSpaces(content, strings.Trim(content, " "))
		if first_start_spaces == -1 {
			first_start_spaces = leading_spaces
		}
		if isEmptyLine(content) {
			break
		} else if leading_spaces == first_start_spaces && isUL(content) {
			closeLI(state, index-1, &open_li) // -1 because the line the line is from the previous line
			openLI(state, index, &open_ul, &open_li)
			TokenizeBlock(state, leading_spaces+2)
		} else if leading_spaces >= 2+offset && TokenizeBlock(state, 2+offset) {
		} else if leading_spaces < 2+offset {
			break
		} else {
			TokenizeBlock(state, offset+leading_spaces)
		}
		index = state.GetLastToken().Line + 1
	}
	closeLI(state, index-1, &open_li)
	closeUL(state, index-1, &open_ul)
	return true
}

func isUL(str string) bool {
	left_trimed := strings.Trim(str, " ")
	leading_spaces := countLeadingSpaces(str, left_trimed)
	if leading_spaces >= 4 {
		return true
	}
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
