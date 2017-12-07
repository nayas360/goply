package goply

// this is passed to an error func
type LexerState struct {
	SourceLength int    // the length of the source string
	Source       string // the source string itself
	Position     int    // current position in the source string
	LineNum      int    // current Line number
	ColNum       int    // current column number
}
