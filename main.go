package main

import (
	"fmt"
	"regexp"
)

var rules = [...]Rule{
	// Headers
	{
		name:           "Header 6",
		markdownToHtml: Key{compiled: regexp.MustCompile(`#{6}\s?([^\n]+)`), string: `<h6>$1</h6>`},
		htmlToMarkdown: Key{compiled: regexp.MustCompile(`<h6>\s?([^\n]+)</h6>`), string: `###### $1`},
	},
	{
		name:           "Header 5",
		markdownToHtml: Key{compiled: regexp.MustCompile(`#{5}\s?([^\n]+)`), string: `<h5>$1</h5>`},
		htmlToMarkdown: Key{compiled: regexp.MustCompile(`<h5>\s?([^\n]+)</h5>`), string: `##### $1`},
	},
	{
		name:           "Header 4",
		markdownToHtml: Key{compiled: regexp.MustCompile(`#{4}\s?([^\n]+)`), string: `<h4>$1</h4>`},
		htmlToMarkdown: Key{compiled: regexp.MustCompile(`<h4>\s?([^\n]+)</h4>`), string: `#### $1`},
	},
	{
		name:           "Header 3",
		markdownToHtml: Key{compiled: regexp.MustCompile(`#{3}\s?([^\n]+)`), string: `<h3>$1</h3>`},
		htmlToMarkdown: Key{compiled: regexp.MustCompile(`<h3>\s?([^\n]+)</h3>`), string: `### $1`},
	},
	{
		name:           "Header 2",
		markdownToHtml: Key{compiled: regexp.MustCompile(`#{2}\s?([^\n]+)`), string: `<h2>$1</h2>`},
		htmlToMarkdown: Key{compiled: regexp.MustCompile(`<h2>\s?([^\n]+)</h6>`), string: `## $1`},
	},
	{
		name:           "Header 1",
		markdownToHtml: Key{compiled: regexp.MustCompile(`#{1}\s?([^\n]+)`), string: `<h1>$1</h1>`},
		htmlToMarkdown: Key{compiled: regexp.MustCompile(`<h1>\s?([^\n]+)</h1>`), string: `# $1`},
	},
	// Syntax style
	{
		name:           "bold",
		markdownToHtml: Key{compiled: regexp.MustCompile(`(?s)\*\*(.*?)\*\*`), string: "<b>$1</b>"},
		htmlToMarkdown: Key{compiled: regexp.MustCompile(`<b>\s?([^\n]+)</b>`), string: `**$1**`},
	},
	// Markdown table

}

func main() {
	input := `   ___
	du texte
___
des séparateurs
  ***
___
`

	result1 := parseMarkdownRegex(input, leafBlocsRules)
	result2 := parseHTMLRegex(result1, leafBlocsRules)

	fmt.Println("------------------MARKDOWN-TO-HTML------------------")
	fmt.Println(result1)
	fmt.Println("------------------HTML-TO-MARKDOWN------------------")
	fmt.Println(result2)
}
