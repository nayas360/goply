package goply

import (
	"crypto/sha1"
	"fmt"
)

// computes and returns a sha1 hash
func computeSha1(text string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(text)))
}

// checks if the given token is a terminal token
func isTerminalType(Type string) bool {
	return len(Type) > len(plx_TERMINAL_B) && plx_TERMINAL_B == Type[:len(plx_TERMINAL_B)]
}

// The default error handler
func defaultLexerError(ls LexerState) error {
	return fmt.Errorf("line %d, column %d: could not match '%c' with any rule", ls.LineNum, ls.ColNum,
		ls.Source[ls.Position])
}

// The parsers default error handler for lexer
func plxDefaultErrorFunc(ls LexerState) error {
	symbol := string(ls.Source[ls.Position])
	if symbol == "\n" {
		symbol = "\\n"
	}
	return fmt.Errorf("unexpected symbol '%s' at position '%d' in '%s'", symbol,
		ls.Position, ls.Source)
}
