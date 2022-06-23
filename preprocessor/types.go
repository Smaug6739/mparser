package preprocessor

type Markdown struct {
	Source       string   // The brute string
	Lines        []string // The source splited on character "\n"
	TotalIndexes int      // Total of indexes
	Tokens       Tokens   // Tokens (see: Token)
}

type Tokens = []Token

type Token struct {
	Token    string // token type
	Html     string // html render
	Markdown string // markdown render
	Content  string // The content of literal
	Line     int    // the line in array of lines (start to line -1 (index+1) so first line is index [0])
	Block    bool   // If the token is a block or inline
}
