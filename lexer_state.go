package goply

// Struct having fields for representing the state of a lexer
// it is passed as an argument to lexerErrorFunc
type LexerState struct {
	SourceLength int    // the length of the source string
	Source       string // the source string itself
	Position     int    // current position in the source string
	LineNum      int    // current Line number
	ColNum       int    // current column number
}
