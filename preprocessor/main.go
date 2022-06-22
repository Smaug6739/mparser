package preprocessor

func New(str string) *Markdown {
	str = normalize(str) // Normalize line ending
	return prepar(str)   // Markdown instance
}
