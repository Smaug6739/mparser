package mparser

import (
	"testing"

	"github.com/Smaug6739/mparser/internal/logger"
)

func test(t *testing.T, name, input string, markdown, html, content [3]string) {
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
	/*test(t, "Headers 1", "# Header 1", [3]string{"# ", "", ""}, [3]string{"<h1>", "", "</h1>"}, [3]string{"", "Header 1", ""})
	test(t, "Headers 2", "## Header 2", [3]string{"## ", "", ""}, [3]string{"<h2>", "", "</h2>"}, [3]string{"", "Header 2", ""})
	test(t, "Headers 3", "### Header 3", [3]string{"### ", "", ""}, [3]string{"<h3>", "", "</h3>"}, [3]string{"", "Header 3", ""})
	test(t, "Headers 4", "#### Header 4", [3]string{"#### ", "", ""}, [3]string{"<h4>", "", "</h4>"}, [3]string{"", "Header 4", ""})
	test(t, "Headers 5", "##### Header 5", [3]string{"##### ", "", ""}, [3]string{"<h5>", "", "</h5>"}, [3]string{"", "Header 5", ""})
	test(t, "Headers 6", "###### Header 6", [3]string{"###### ", "", ""}, [3]string{"<h6>", "", "</h6>"}, [3]string{"", "Header 6", ""})
	test(t, "Headers 7", "#", [3]string{"# ", "", ""}, [3]string{"<h1>", "", "</h1>"}, [3]string{"", "", ""})
	//TODO: #Header = paragraph
	test(t, "Headers 8", " ## Header 8", [3]string{"# ", "", ""}, [3]string{"<h2>", "", "</h2>"}, [3]string{"", "Header 8", ""})
	test(t, "Headers 9", "  ### Header 9", [3]string{"# ", "", ""}, [3]string{"<h3>", "", "</h3>"}, [3]string{"", "Header 9", ""})
	test(t, "Headers 10", "   # Header 10", [3]string{"# ", "", ""}, [3]string{"<h1>", "", "</h1>"}, [3]string{"", "Header 10", ""})*/
}
func TestTokenize(t *testing.T) {
	input := `# Headers 1
 Paragraph one

  Paragraph two
  Paragraph two`
	tokenized := Tokenize(input)
	logger.New().Details(tokenized)
}
