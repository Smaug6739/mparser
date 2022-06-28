package preprocessor

import (
	"strings"
)

func prepar(str string) *Markdown {
	// The main Markdown instance
	lines := strings.Split(str, "\n")
	var instance Markdown = Markdown{
		Source:   str,
		Lines:    lines,
		MaxIndex: len(lines) - 1,
		Tokens: Tokens{
			Token{
				Token:    "",
				Html:     "",
				Markdown: "",
				Line:     -1, // +1 = 0; so index match
				Block:    true,
			},
		},
	}
	return &instance
}
func (md *Markdown) GetLine(index int) string {
	return md.Lines[index]
}
func (md *Markdown) GetLastTokenSliceIndex() int {
	return len(md.Tokens) - 1
}
func (md *Markdown) GetLastToken() Token {
	return md.Tokens[len(md.Tokens)-1]
}

func (md *Markdown) GetLastBlockToken() (Token, int) {
	for i := len(md.Tokens) - 1; i >= 0; i-- {
		if md.Tokens[i].Block {
			return md.Tokens[i], i
		}
	}
	return md.Tokens[0], 0
}
func (md *Markdown) GetToken(index int) Token {
	return md.Tokens[index]
}

type Data struct {
	LastTokenSliceIndex int
	LastToken           *Token
	LineIndex           int
	LineContent         string
}

func (md *Markdown) GetData() (Data, bool) {
	LastTokenSliceIndex := md.GetLastTokenSliceIndex()
	LastToken := md.GetToken(LastTokenSliceIndex)
	LineIndex := LastToken.Line + 1
	LineContent := md.GetLine(LineIndex)
	if LineIndex > md.MaxIndex {
		return Data{}, false
	}
	return Data{
		LastTokenSliceIndex: LastTokenSliceIndex,
		LastToken:           &LastToken,
		LineIndex:           LineIndex,
		LineContent:         LineContent,
	}, true
}
