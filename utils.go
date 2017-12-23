package goply

import (
	"crypto/sha1"
	"fmt"
)

// computes and returns a sha1 hash
func computeSha1(text string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(text)))
}

// The default error handler
func defaultLexerError(ls LexerState) error {
	return fmt.Errorf("line %d, column %d: could not match '%c' with any rule", ls.LineNum, ls.ColNum,
		ls.Source[ls.Position])
}
