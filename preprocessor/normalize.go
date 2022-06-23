package preprocessor

import (
	"regexp"
)

var NEWLINES = "\r\n?|\n"
var NULL = `\0`

var (
	NEWLINES_RE = regexp.MustCompile(NEWLINES)
	NULL_RE     = regexp.MustCompile(NULL)
)

func normalize(str string) string {
	str = NEWLINES_RE.ReplaceAllString(str, "\n")
	str = NULL_RE.ReplaceAllString(str, "\uFFFD")
	return str
}
