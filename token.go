package goply

// struct having related fields for representing a token
type Token struct {
	Type             string // the Type of token
	Value            string // the Value of the token
	Length           int    // the length of the token
	StartingPosition int    // the starting position of the Token in the source
	LineNum          int    // the line number of the token
	ColNum           int    // the column number of the token
}

// Helper function to generate a new token
func newToken(Type, Value string, StartingPosition, LineNum, ColNum int) *Token {
	return &Token{Type: Type, Value: Value, Length: len(Value), StartingPosition: StartingPosition,
		LineNum: LineNum, ColNum: ColNum}
}
