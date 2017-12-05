// package name: goply
package goply

// The Token Type
type Token struct {
	Type             string // The Type of token
	Value            string // The Value of the token
	Length           int    // The length of the token
	StartingPosition int    // The starting StartingPosition of the Token in the source
	LineNum          int    // The line number of the token
	ColNum           int    // the column number of the token
}

// Helper function to generate a new token
func newToken(Type, Value string, StartingPosition, LineNum, ColNum int) *Token {
	return &Token{Type: Type, Value: Value, Length: len(Value), StartingPosition: StartingPosition,
		LineNum: LineNum, ColNum: ColNum}
}
