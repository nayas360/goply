package goply

type AstNode interface {
	// This is the entry point for parsing
	Parse(tokens []*Token, id int) error
	// Should print the tree representation
	String() string
}
