package mparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"testing"

	"github.com/Smaug6739/mparser/internal/logger"
	"github.com/Smaug6739/mparser/preprocessor"
)

func test(t *testing.T, name, input string, markdown, html, content []string) {
	tokenized := Tokenize(input)
	for index, v := range tokenized.Tokens[1:] {
		if v.Html != html[index] {
			t.Error("[TEST FAIL (HTML)]: ", name, "\nInput: ", input, "\nExepted result: '", html[index]+"'", "\nResult: '"+v.Html+"'")
		}
		if v.Content != content[index] {
			t.Error("[TEST FAIL (CONTENT)]: ", name, "\nInput: ", input, "\nExepted result: '"+content[index]+"'", "\nResult: '"+v.Content+"'")
		}
	}
}
func test2(t *testing.T, name, input string, result []string) {
	tokenized := Tokenize(input)
	for index, token := range tokenized.Tokens[1:] {
		if token.Html != "" {
			if result[index] != token.Html {
				t.Errorf("TEST FAIL (HTML) %s :\nInput:%s\nExepted result: %s\nActual result:%s", name, input, result[index], token.Html)
			}
		} else {
			if result[index] != token.Content {
				t.Errorf("TEST FAIL (CONTENT) %s :\nInput:%s\nExepted result: %s\nActual result:%s", name, input, result[index], token.Content)
			}
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
	*/

	//Indented code classic
	test2(t, "Indented code 1", "    Code 1", []string{"<pre>", "<code>", "Code 1", "</code>", "</pre>"})
	test2(t, "Indented code 2", "    Code 1\n    Code 2", []string{"<pre>", "<code>", "Code 1", "Code 2", "</code>", "</pre>"})
	test2(t, "Indented code 3", "    Code 1\n    Code 2\n    Code 3", []string{"<pre>", "<code>", "Code 1", "Code 2", "Code 3", "</code>", "</pre>"})
	test2(t, "Indented code 4", "    Code 1\n    Code 2\n      Code 3", []string{"<pre>", "<code>", "Code 1", "Code 2", "  Code 3", "</code>", "</pre>"})
	test2(t, "Indented code 5", "    Code 1\n    Code 2\n   Code 3", []string{"<pre>", "<code>", "Code 1", "Code 2", "</code>", "</pre>", "<p>", "Code 3", "</p>"})
	// Indented code with delimiter
	test2(t, "Indented code delimiter 1", "    - Code 1", []string{"<pre>", "<code>", "- Code 1", "</code>", "</pre>"})
	test2(t, "Indented code delimiter 2", "    - Code 1\n    - Code 2", []string{"<pre>", "<code>", "- Code 1", "- Code 2", "</code>", "</pre>"})

	// Paragraph
	test(t, "Paragraph 1", "Text", []string{"", "", ""}, []string{"<p>", "", "</p>"}, []string{"", "Text", ""})
	test(t, "Paragraph 2", "Text multiple words", []string{"", "", ""}, []string{"<p>", "", "</p>"}, []string{"", "Text multiple words", ""})
	test2(t, "Paragraph 3", "Line one\n\nLine two", []string{"<p>", "Line one", "</p>", "", "<p>", "Line two", "</p>"})
	test2(t, "Paragraph 4", "Line one\n\nLine two\nLine three", []string{"<p>", "Line one", "</p>", "", "<p>", "Line two", "Line three", "</p>"})
	test2(t, "Paragraph 5", "Line one\n\nLine two\nLine three\n\n\nLine four", []string{"<p>", "Line one", "</p>", "", "<p>", "Line two", "Line three", "</p>", "", "", "<p>", "Line four", "</p>"})

	// LIST: Classic
	test2(t, "List", "- Item 1", []string{"<ul>", "<li>", "Item 1", "</li>", "</ul>"})
	test2(t, "List (two items)", "- Item 1\n- Item 2", []string{"<ul>", "<li>", "Item 1", "</li>", "<li>", "Item 2", "</li>", "</ul>"})
	// LIST: Indented
	test2(t, "List (indented 1)", "- Item 1\n  - Item 2", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "</li>", "</ul>", "</li>", "</ul>"})
	test2(t, "List (indented 2)", "- Item 1\n  - Item 2\n  - Item 3", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "</li>", "<li>", "Item 3", "</li>", "</ul>", "</li>", "</ul>"})
	test2(t, "List (indented 3)", "- Item 1\n  - Item 2\n    - Item 3", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "<ul>", "<li>", "Item 3", "</li>", "</ul>", "</li>", "</ul>", "</li>", "</ul>"})
	test2(t, "List (indented 4)", "- Item 1\n  - Item 2\n- Item 3", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "</li>", "</ul>", "</li>", "<li>", "Item 3", "</li>", "</ul>"})
	test2(t, "List (indented 5)", "- Item 1\n  - Item 2\n - Item 3", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "</li>", "</ul>", "</li>", "<li>", "Item 3", "</li>", "</ul>"})
	test2(t, "List (indented 6)", "- Item 1\n  - Item 2\n - Item 3\n\n    - Item 4", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "</li>", "</ul>", "</li>", "<li>", "Item 3", "", "<ul>", "<li>", "Item 4", "</li>", "</ul>", "</li>", "</ul>"})
	test2(t, "List (indented 7)", "- Item 1\n  - Item 2\n - Item 3\n\n    - Item 4\n", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "</li>", "</ul>", "</li>", "<li>", "Item 3", "", "<ul>", "<li>", "Item 4", "", "</li>", "</ul>", "</li>", "</ul>"})
	test2(t, "List (indented 8)", "- Item 1\n  - Item 2\n - Item 3\n\n    - Item 4\n- Item 5", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "</li>", "</ul>", "</li>", "<li>", "Item 3", "", "<ul>", "<li>", "Item 4", "</li>", "</ul>", "</li>", "<li>", "Item 5", "</li>", "</ul>"})
	// LIST: Start with spaces
	test2(t, "List  starts with spaces 1", "  - Item 1\n  - Item 2\n- Item 3", []string{"<ul>", "<li>", "Item 1", "</li>", "<li>", "Item 2", "</li>", "<li>", "Item 3", "</li>", "</ul>"})
	test2(t, "List  starts with spaces 2", "  - Item 1\n- Item 2\n- Item 3", []string{"<ul>", "<li>", "Item 1", "</li>", "<li>", "Item 2", "</li>", "<li>", "Item 3", "</li>", "</ul>"})
	test2(t, "List  starts with spaces 3", "   - Item 1\n- Item 2\n- Item 3", []string{"<ul>", "<li>", "Item 1", "</li>", "<li>", "Item 2", "</li>", "<li>", "Item 3", "</li>", "</ul>"})
	test2(t, "List  starts with spaces 4", "   - Item 1\n   - Item 2\n- Item 3", []string{"<ul>", "<li>", "Item 1", "</li>", "<li>", "Item 2", "</li>", "<li>", "Item 3", "</li>", "</ul>"})
	test2(t, "List  starts with spaces 5", "   - Item 1\n   - Item 2\n   - Item 3", []string{"<ul>", "<li>", "Item 1", "</li>", "<li>", "Item 2", "</li>", "<li>", "Item 3", "</li>", "</ul>"})
	test2(t, "List  starts with spaces 6", "   - Item 1\n   - Item 2\n    - Item 3", []string{"<ul>", "<li>", "Item 1", "</li>", "<li>", "Item 2", "</li>", "<li>", "Item 3", "</li>", "</ul>"})
	// LIST: Empty lines
	test2(t, "List (test blank lines 1)", "- Item 1\n\n- Item 2", []string{"<ul>", "<li>", "Item 1", "", "</li>", "<li>", "Item 2", "</li>", "</ul>"})
	test2(t, "List (test blank lines 2)", "- Item 1\n\n\n- Item 2", []string{"<ul>", "<li>", "Item 1", "", "", "</li>", "<li>", "Item 2", "</li>", "</ul>"})
	test2(t, "List (test blank lines 3)", "- Item 1\n\n\n\n- Item 2", []string{"<ul>", "<li>", "Item 1", "", "", "", "</li>", "<li>", "Item 2", "</li>", "</ul>"})
	test2(t, "List (test blank lines 4)", "- Item 1\n\n\n- Item 2\n- Item 3", []string{"<ul>", "<li>", "Item 1", "", "", "</li>", "<li>", "Item 2", "</li>", "<li>", "Item 3", "</li>", "</ul>"})
	test2(t, "List (test blank lines with paragraph 1)", "- Item 1\n\n- Item 2\n\nItem 3", []string{"<ul>", "<li>", "Item 1", "", "</li>", "<li>", "Item 2", "", "</li>", "</ul>", "<p>", "Item 3", "</p>"})
	// LIST: Offset
	test2(t, "List offset 1", "- Item 1\n - Item 2", []string{"<ul>", "<li>", "Item 1", "</li>", "<li>", "Item 2", "</li>", "</ul>"})
	test2(t, "List offset 2", "- Item 1\n  - Item 2", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "</li>", "</ul>", "</li>", "</ul>"})
	test2(t, "List offset 3", "-  Item 1\n   - Item 2", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "</li>", "</ul>", "</li>", "</ul>"})
	test2(t, "List offset 4", "-  Item 1\n    - Item 2", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "</li>", "</ul>", "</li>", "</ul>"})
	test2(t, "List offset 5", "-   Item 1\n     - Item 2", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "</li>", "</ul>", "</li>", "</ul>"})
	// LIST: Lazy continuation line
	test2(t, "List lazy continuation lines 1", "- Item 1\nText (in the item one)", []string{"<ul>", "<li>", "Item 1", "Text (in the item one)", "</li>", "</ul>"})
	test2(t, "List lazy continuation lines 2", "- Item 1\nText (in the item one)\nItem 1", []string{"<ul>", "<li>", "Item 1", "Text (in the item one)", "Item 1", "</li>", "</ul>"})
	test2(t, "List lazy continuation lines 3", "- Item 1\nText (in the item one)\nItem 1\n\nParagraph 1", []string{"<ul>", "<li>", "Item 1", "Text (in the item one)", "Item 1", "", "</li>", "</ul>", "<p>", "Paragraph 1", "</p>"})
	test2(t, "List lazy continuation lines 4", "- Item 1\n  - Item 2\nfoo\nbar", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2", "foo", "bar", "</li>", "</ul>", "</li>", "</ul>"})
	// LIST: Lazy continuation line and new item
	test2(t, "List lazy lines + item 1", "- Item 1\n  - Item 2 (indented)\nLazy continuation item 2\n  - Item 3", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2 (indented)", "Lazy continuation item 2", "</li>", "<li>", "Item 3", "</li>", "</ul>", "</li>", "</ul>"})
	test2(t, "List lazy lines + item 1", "- Item 1\n  - Item 2 (indented)\nLazy continuation item 2\n  - Item 3\nItem 3 lazy", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2 (indented)", "Lazy continuation item 2", "</li>", "<li>", "Item 3", "Item 3 lazy", "</li>", "</ul>", "</li>", "</ul>"})
	test2(t, "List lazy lines + item 1", "- Item 1\n  - Item 2 (indented)\nLazy continuation item 2\n  - Item 3\n- Item 4", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2 (indented)", "Lazy continuation item 2", "</li>", "<li>", "Item 3", "</li>", "</ul>", "</li>", "<li>", "Item 4", "</li>", "</ul>"})
	test2(t, "List lazy lines + item 1", "- Item 1\n  - Item 2 (indented)\nLazy continuation item 2\n  - Item 3\n- Item 4\nItem 4 lazy", []string{"<ul>", "<li>", "Item 1", "<ul>", "<li>", "Item 2 (indented)", "Lazy continuation item 2", "</li>", "<li>", "Item 3", "</li>", "</ul>", "</li>", "<li>", "Item 4", "Item 4 lazy", "</li>", "</ul>"})
	// LIST: With other blocks
	test2(t, "List indented code", "-     Item 1", []string{"<ul>", "<li>", "<pre>", "<code>", "Item 1", "</code>", "</pre>", "</li>", "</ul>"})
	test2(t, "List indented code", "-     Item 1\n-     Item 2", []string{"<ul>", "<li>", "<pre>", "<code>", "Item 1", "Item 2", "</code>", "</pre>", "</li>", "</ul>"})
	test2(t, "List indented code", "-     Item 1\n -     Item 2", []string{"<ul>", "<li>", "<pre>", "<code>", "Item 1", " Item 2", "</code>", "</pre>", "</li>", "</ul>"})

	// Quotes (citations)
	/*	test2(t, "Quote 1 (normal)", "> Citation 1", []string{"<blockquote>", "<p>", "Citation 1", "</p>", "</blockquote>"})
		test2(t, "Quote 2 (next line 1)", "> Citation 1\nCitation 2", []string{"<blockquote>", "<p>", "Citation 1", "Citation 2", "</p>", "</blockquote>"})
		test2(t, "Quote 3 (next line 2)", "> Citation 1\n> Citation 2", []string{"<blockquote>", "<p>", "Citation 1", "Citation 2", "</p>", "</blockquote>"})
		test2(t, "Quote 4 (other block 1)", "> Citation 1\n>> Citation 2", []string{"<blockquote>", "<p>", "Citation 1", "</p>", "<blockquote>", "<p>", "Citation 2", "</p>", "</blockquote>", "</blockquote>"})
		test2(t, "Quote 4 (other block 2)", "> Citation 1\n>>> Citation 2", []string{"<blockquote>", "<p>", "Citation 1", "</p>", "<blockquote>", "<blockquote>", "<p>", "Citation 2", "</p>", "</blockquote>", "</blockquote>", "</blockquote>"})
		test2(t, "Quote 5 (other block 3)", "> Citation 1\n>>> Citation 2\nCitation 3", []string{"<blockquote>", "<p>", "Citation 1", "</p>", "<blockquote>", "<blockquote>", "<p>", "Citation 2", "Citation 3", "</p>", "</blockquote>", "</blockquote>", "</blockquote>"})
		test2(t, "Quote 6 (other block 4)", "> Citation 1\n>>> Citation 2\n> Citation 3", []string{"<blockquote>", "<p>", "Citation 1", "</p>", "<blockquote>", "<blockquote>", "<p>", "Citation 2", "Citation 3", "</p>", "</blockquote>", "</blockquote>", "</blockquote>"})
		test2(t, "Quote 7 (other block 5)", "> Citation 1\n>>> Citation 2\n>Citation 3", []string{"<blockquote>", "<p>", "Citation 1", "</p>", "<blockquote>", "<blockquote>", "<p>", "Citation 2", "Citation 3", "</p>", "</blockquote>", "</blockquote>", "</blockquote>"})
		test2(t, "Quote 8 (other block 6)", "> Citation 1\n>>> Citation 2\n>>Citation 3", []string{"<blockquote>", "<p>", "Citation 1", "</p>", "<blockquote>", "<blockquote>", "<p>", "Citation 2", "Citation 3", "</p>", "</blockquote>", "</blockquote>", "</blockquote>"})
		test2(t, "Quote 9 (other block 7)", "> Citation 1\n>>> Citation 2\n> > Citation 3", []string{"<blockquote>", "<p>", "Citation 1", "</p>", "<blockquote>", "<blockquote>", "<p>", "Citation 2", "Citation 3", "</p>", "</blockquote>", "</blockquote>", "</blockquote>"})
		test2(t, "Quote 10 (other block 8)", "> Citation 1\n>>> Citation 2\n>>> Citation 3", []string{"<blockquote>", "<p>", "Citation 1", "</p>", "<blockquote>", "<blockquote>", "<p>", "Citation 2", "Citation 3", "</p>", "</blockquote>", "</blockquote>", "</blockquote>"})
	*/
}

func TestTokenize(t *testing.T) {
	input := `
- Item 1
  - Item 2 (indented)
Item 2 (lazy continuation)
  - Item three (indented)
`
	/*input := `
	  `*/ // TODO: Paragraph empty
	tokenized := Tokenize(input)
	var last_token preprocessor.Token = tokenized.Tokens[0]
	HTML := "<div>"
	for _, v := range tokenized.Tokens {
		HTML += v.Html
		if last_token.Content != "" && v.Content != "" {
			HTML += "\n" + v.Content
		} else {
			HTML += v.Content
		}
		last_token = v
	}
	HTML += "</div>"
	logger.New().Details(tokenized)
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
