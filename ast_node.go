package goply

type AstNode interface {
	// This is the entry point for parsing
	Parse(tokens *TokenStream) error
	// Should print the tree representation
	String() string
}
