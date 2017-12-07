package goply_test

import (
	"fmt"

	"github.com/nayas360/goply"
)

func ExampleNewLexer() {
	// Create a new lexer with lisp like syntax
	lexer := goply.NewLexer("(+ 10 20)")
	// match left parenthesis
	lexer.AddRule("<lparen>", "\\(")
	// match right parenthesis
	lexer.AddRule("<rparen>", "\\)")
	// operator +
	lexer.AddRule("<op_plus>", "\\+")
	// a integer number
	lexer.AddRule("<number>", "[0-9]+")
	// ignore all whitespace
	lexer.Ignore("\\s+")
	// get the tokens
	tokens, err := lexer.GetTokens()
	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		fmt.Printf("Got %s : %s\n", token.Type, token.Value)
	}
	// Output:
	// Got <lparen> : (
	// Got <op_plus> : +
	// Got <number> : 10
	// Got <number> : 20
	// Got <rparen> : )
}
