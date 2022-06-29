package mparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"testing"

	"github.com/Smaug6739/mparser/internal/logger"
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
	/*test(t, "Headers 1", "# Header 1", []string{"# ", "", ""}, []string{"<h1>", "", "</h1>"}, []string{"", "Header 1", ""})
	test(t, "Headers 2", "## Header 2", []string{"## ", "", ""}, []string{"<h2>", "", "</h2>"}, []string{"", "Header 2", ""})
	test(t, "Headers 3", "### Header 3", []string{"### ", "", ""}, []string{"<h3>", "", "</h3>"}, []string{"", "Header 3", ""})
	test(t, "Headers 4", "#### Header 4", []string{"#### ", "", ""}, []string{"<h4>", "", "</h4>"}, []string{"", "Header 4", ""})
	test(t, "Headers 5", "##### Header 5", []string{"##### ", "", ""}, []string{"<h5>", "", "</h5>"}, []string{"", "Header 5", ""})
	test(t, "Headers 6", "###### Header 6", []string{"###### ", "", ""}, []string{"<h6>", "", "</h6>"}, []string{"", "Header 6", ""})
	test(t, "Headers 7 (empty)", "#", []string{"# ", "", ""}, []string{"<h1>", "", "</h1>"}, []string{"", "", ""})
	test(t, "Headers 8", " ## Header 8", []string{"# ", "", ""}, []string{"<h2>", "", "</h2>"}, []string{"", "Header 8", ""})
	test(t, "Headers 9", "  ### Header 9", []string{"# ", "", ""}, []string{"<h3>", "", "</h3>"}, []string{"", "Header 9", ""})
	test(t, "Headers 10", "   # Header 10", []string{"# ", "", ""}, []string{"<h1>", "", "</h1>"}, []string{"", "Header 10", ""})
	test(t, "Headers 11 (trim)", "   #         Header 11        ", []string{"# ", "", ""}, []string{"<h1>", "", "</h1>"}, []string{"", "Header 11", ""})
	test(t, "Headers 12 (paragraph)", "   #Header 12        ", []string{"", "", ""}, []string{"<p>", "", "</p>"}, []string{"", "#Header 12", ""})

	// Thematic breaks
	test(t, "Thematic breaks 1", "---", []string{"---"}, []string{"</hr>"}, []string{""})
	test(t, "Thematic breaks 2", "___", []string{"---"}, []string{"</hr>"}, []string{""})
	test(t, "Thematic breaks 3", "***", []string{"---"}, []string{"</hr>"}, []string{""})
	test(t, "Thematic breaks 4", "  ***", []string{"---"}, []string{"</hr>"}, []string{""})
	test(t, "Thematic breaks 5", "   ***", []string{"---"}, []string{"</hr>"}, []string{""})
	test(t, "Thematic breaks 6", "   *  *     *", []string{"---"}, []string{"</hr>"}, []string{""})
	test(t, "Thematic breaks 7", "   **   *", []string{"---"}, []string{"</hr>"}, []string{""})
	//test(t, "Thematic breaks 8", "       *  a**", []string{"", "", ""}, []string{"<p>", "", "</p>"}, []string{"", "*  a**", ""})
	//test(t, "Thematic breaks 9", "    *  **", []string{"", "", ""}, []string{"<p>", "", "</p>"}, []string{"", "*  **", ""}) //TODO: 4 spaces => remove for code block

	// Lheaders
	test(t, "Line headers 1", "Header\n=", []string{"", "", "===", "==="}, []string{"<h1>", "", "</h1>", ""}, []string{"", "Header", "", "==="})
	test(t, "Line headers 2", "Header\n---", []string{"", "", "---", "---"}, []string{"<h2>", "", "</h2>", ""}, []string{"", "Header", "", "---"})

	// Indented code
	test(t, "Indented code 1", "    code", []string{"", "", "    ", "    "}, []string{"<pre>", "<code>", "", "</code>", "</pre>", ""}, []string{"", "", "    code", "", ""})
	test(t, "Indented code 2", "    code\n    code", []string{"", "", "", "", "", ""}, []string{"<pre>", "<code>", "", "", "</code>", "</pre>"}, []string{"", "", "    code", "    code", "", ""})

	// Paragraph
	test(t, "Paragraph 1", "Text", []string{"", "", ""}, []string{"<p>", "", "</p>"}, []string{"", "Text", ""})
	test(t, "Paragraph 2", "Text multiple words", []string{"", "", ""}, []string{"<p>", "", "</p>"}, []string{"", "Text multiple words", ""})*/
}
func TestTokenize(t *testing.T) {
	input := ``
	tokenized := Tokenize(input)
	logger.New().Details(tokenized)
	HTML := "<div>"
	for _, v := range tokenized.Tokens {
		HTML += v.Html
		HTML += v.Content
	}
	HTML += "</div>"
	r, e := formatXML([]byte(HTML))
	if e != nil {
		panic(e)
	} else {
		fmt.Println(string(r))
	}
}

func formatXML(data []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return b.Bytes(), nil
		}
		if err != nil {
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
}
