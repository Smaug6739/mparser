package tokenizer_block

func countLeadingSpaces(str1, trimmedStr string) int {
	return len(str1) - len(trimmedStr)
}
