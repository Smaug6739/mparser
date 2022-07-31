package preprocessor

type Markdown struct {
	Source   string   // The brute string
	Lines    []string // The source splited on character "\n"
	MaxIndex int      // Maximal index
	Tokens   Tokens   // Tokens (see: Token)
	Meta     MetaMarkdown
}

type Tokens = []Token

type Token struct {
	Token    string // token type
	Html     string // html render
	Markdown string // markdown render
	Content  string // The content of literal
	Line     int    // the line in array of lines (start to line -1 (index+1) so first line is index [0])
	Closer   bool   // If the token is a closer or not
	Meta     Meta   // The meta data
}

type Meta struct {
	leading_spaces int
	offset         int
	level          int
}
type MetaMarkdown struct {
	Blank_lines_allowed bool
}
