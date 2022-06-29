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
				Token: "internal_init",
				Line:  -1, // +1 = 0; so index match
				Block: true,
			},
		},
	}
	return &instance
}
func (md *Markdown) GetLine(index int) string {
	if index > md.MaxIndex {
		return ""
	}
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

func (md *Markdown) GetData(offset int) (Data, bool) {
	LastToken, LastTokenSliceIndex := md.GetLastBlockToken()
	LineIndex := LastToken.Line + 1
	LineContent := md.GetLine(LineIndex)
	if len(LineContent) >= offset {
		LineContent = LineContent[offset:]
	} else {
		LineContent = ""
	}
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
