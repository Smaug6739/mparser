package mparser

import (
	"testing"
)

func test(t *testing.T, name, input string, markdown, html, content []string) {
	tokenized := Tokenize(input)
	for index, v := range tokenized.Tokens[1:] {
		/*if v.Markdown != markdown[index] {
			t.Error("[TEST FAIL (MARKDOWN)]: ", name, "\nInput: ", input, "\nExepted result: '", markdown[index]+"'", "\nResult: '"+v.Markdown+"'")
		}*/
		if v.Html != html[index] {
			t.Error("[TEST FAIL (HTML)]: ", name, "\nInput: ", input, "\nExepted result: '", html[index]+"'", "\nResult: '"+v.Html+"'")
		}
		if v.Content != content[index] {
			t.Error("[TEST FAIL (CONTENT)]: ", name, "\nInput: ", input, "\nExepted result: '"+content[index]+"'", "\nResult: '"+v.Content+"'")
		}
	}
}
func TestTokenizeAuto(t *testing.T) {
	// Headers
	test(t, "Headers 1", "# Header 1", []string{"# ", "", ""}, []string{"<h1>", "", "</h1>"}, []string{"", "Header 1", ""})
	/*test(t, "Headers 2", "## Header 2", []string{"## ", "", ""}, []string{"<h2>", "", "</h2>"}, []string{"", "Header 2", ""})
	test(t, "Headers ", "### Header ", []string{"### ", "", ""}, []string{"<h>", "", "</h>"}, []string{"", "Header ", ""})
	test(t, "Headers 4", "#### Header 4", []string{"#### ", "", ""}, []string{"<h4>", "", "</h4>"}, []string{"", "Header 4", ""})
	test(t, "Headers 5", "##### Header 5", []string{"##### ", "", ""}, []string{"<h5>", "", "</h5>"}, []string{"", "Header 5", ""})
	test(t, "Headers 6", "###### Header 6", []string{"###### ", "", ""}, []string{"<h6>", "", "</h6>"}, []string{"", "Header 6", ""})
	test(t, "Headers 7", "#", []string{"# ", "", ""}, []string{"<h1>", "", "</h1>"}, []string{"", "", ""})
	//TODO: #Header = paragraph
	test(t, "Headers 8", " ## Header 8", []string{"# ", "", ""}, []string{"<h2>", "", "</h2>"}, []string{"", "Header 8", ""})
	test(t, "Headers 9", "  ### Header 9", []string{"# ", "", ""}, []string{"<h>", "", "</h>"}, []string{"", "Header 9", ""})
	test(t, "Headers 10", "   # Header 10", []string{"# ", "", ""}, []string{"<h1>", "", "</h1>"}, []string{"", "Header 10", ""})*/
	// Paragraph
}

/*func TestTokenize(t *testing.T) {
	input := `# Headers 1
 Paragraph one

  Paragraph two
  Paragraph two`
	tokenized := Tokenize(input)
	logger.New().Details(tokenized)
}*/
