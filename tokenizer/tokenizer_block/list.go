package tokenizer_block

import (
	"strings"

	"github.com/Smaug6739/mparser/preprocessor"
)

//INFO: Ceci n'est que pour les listes de type "ul"
func tokenizeBlockList(state *preprocessor.Markdown, skip int) bool {
	// Get common informations
	data, err := getInfos(state, skip)
	if err != nil {
		return false
	}
	leftTrimmed := strings.TrimLeft(data.lineContent, " ")

	leadingSpaces := countLeadingSpaces(data.lineContent, leftTrimmed)
	/*if leadingSpaces >= 4 {
		return false
	}*/
	if !isUlItem(leftTrimmed) {
		return false
	}
	var ul_opened bool = false
	var li_oppened bool = false                  // true if an "li" token is oppened
	var start_spaces, off int = leadingSpaces, 0 // Déterminé par le premier item de la liste
	var is_item bool = false                     // Vérification
	var index int = 0                            // Index de la ligne analysée

	// On analyse chaque ligne de la liste
	for index = data.lineIndex; index <= state.MaxIndex; index++ {
		content := state.Lines[index]
		start_blank_spaces := 0
		is_item, off, start_blank_spaces = ulItem(content)
		if is_item && start_blank_spaces == start_spaces {
			handleOpenUl(state, index, &ul_opened)
			handleOpenLi(state, index, &li_oppened)
			if ListTokenizeBlock(state, off+skip) {
			} else {
				state.Tokens = append(state.Tokens, preprocessor.Token{
					Token:   "inline",
					Content: content,
					Line:    index,
					Block:   true,
				})
			}
			continue
		}
		// CONDITIONS DE FIN DE LISTE \\
		// 1. La ligne est vide
		// 2. La ligne est un bloc "thematic_break" valide
		// 3. La ligne est un Token avec moins de 2 espaces au début

		if isEmptyLine(content) {
			break
		}
		is_block := ListTokenizeBlock2(state, 0)
		if is_block && start_blank_spaces < start_spaces+skip {
			panic("TODO")
		} else if is_block {
		} else {
			state.Tokens = append(state.Tokens, preprocessor.Token{
				Token:   "inline",
				Content: content,
				Line:    index,
				Block:   false,
			})
		}
	}
	handleCloseLi(state, index, &li_oppened)
	handleUlClose(state, index-1, &ul_opened)
	return true
}

// UTILS: Functions for code lisibility \\
func openLi(state *preprocessor.Markdown, index int) {
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token: "li_open",
		Html:  "<li>",
		Line:  index,
		Block: false,
	})
}

func closeLi(state *preprocessor.Markdown, index int) {
	state.Tokens = append(state.Tokens, preprocessor.Token{
		Token: "li_close",
		Html:  "</li>",
		Line:  index,
		Block: true,
	})
}

func handleOpenLi(state *preprocessor.Markdown, index int, is_oppened *bool) {
	if *is_oppened {
		closeLi(state, index-1)
		openLi(state, index)
	} else {
		openLi(state, index)
		*is_oppened = true
	}
}
func handleCloseLi(state *preprocessor.Markdown, index int, is_oppened *bool) {
	if *is_oppened {
		closeLi(state, index)
		*is_oppened = false
	}
}
func handleOpenUl(state *preprocessor.Markdown, index int, is_oppened *bool) {
	if *is_oppened {
		return
	} else {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token: "ul_open",
			Html:  "<ul>",
			Line:  index,
			Block: false, //TODO: Improve this
		})
		*is_oppened = true
	}
}
func handleUlClose(state *preprocessor.Markdown, index int, is_oppened *bool) {
	if *is_oppened {
		state.Tokens = append(state.Tokens, preprocessor.Token{
			Token: "ul_close",
			Html:  "</ul>",
			Line:  index,
			Block: true,
		})
		*is_oppened = false
	}
}
