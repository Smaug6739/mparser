package mparser

import (
	"testing"
)

func TestParse(t *testing.T) {
	Parse(`####### Hello world
## Header 2`)
}
