package goply

// used by lexer internally to maintain its state
// passed to error handler and maybe used to report a useful error
type LexerState struct {
	SourceLength int    // the length of the source string
	Source       string // the source string itself
	Position     int    // current position in the source string
	LineNum      int    // current Line number
	ColNum       int    // current column number
}
