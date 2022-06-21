package main

import "regexp"

// This section describes the different kinds of leaf block that make up a Markdown document.

var leafBlocsRules = []Rule{
	// Thematic breaks
	{
		name:           "Thematic breaks",
		markdownToHtml: Key{compiled: regexp.MustCompile(`(?m)^[^\w*]{0,3}(\_\_\_)|(\*\*\*)|(-\-\-)]\n`), string: `<hr />`},
		htmlToMarkdown: Key{compiled: regexp.MustCompile(`<hr />`), string: "___\n"},
	},
}
