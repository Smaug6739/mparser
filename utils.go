package main

import (
	"regexp"
)

type Key struct {
	compiled *regexp.Regexp
	string   string
}
type Rule struct {
	name           string
	markdownToHtml Key
	htmlToMarkdown Key
}

func parseMarkdownRegex(text string, rules []Rule) string {
	for _, value := range rules {
		/*test := value.markdownToHtml.compiled.FindAllString(text, -1)
		for i, tc := range test {
			fmt.Println(i, tc)
		}*/
		text = value.markdownToHtml.compiled.ReplaceAllString(text, value.markdownToHtml.string)
	}
	return text
}

func parseHTMLRegex(text string, rules []Rule) string {
	for _, value := range rules {
		text = value.htmlToMarkdown.compiled.ReplaceAllString(text, value.htmlToMarkdown.string)
	}
	return text
}
