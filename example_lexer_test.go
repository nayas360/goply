package goply_test

import (
	"fmt"

	"github.com/nayas360/goply"
)

func ExampleNewLexer() {
	// Create a new lexer with lisp like syntax
	// The stray = is has no matching rule but should not cause an error
	lexer := goply.NewLexer(false)
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
	tokens, err := lexer.GetTokenStream("= (+ 10 20)")
	if err != nil {
		panic(err)
	}
	// print out the tokens
	for token := range tokens.Iter() {
		fmt.Printf("Got %s : %s\n", token.Type, token.Value)
	}
	// Output:
	// Got <lparen> : (
	// Got <op_plus> : +
	// Got <number> : 10
	// Got <number> : 20
	// Got <rparen> : )
}

func ExampleNewLexer_strict() {
	// Create a new lexer with lisp like syntax
	lexer := goply.NewLexer(true)
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
	tokens, err := lexer.GetTokenStream("(+ 10 20)")
	if err != nil {
		panic(err)
	}
	// print out the tokens
	for token := range tokens.Iter() {
		fmt.Printf("Got %s : %s\n", token.Type, token.Value)
	}
	// Output:
	// Got <lparen> : (
	// Got <op_plus> : +
	// Got <number> : 10
	// Got <number> : 20
	// Got <rparen> : )
}

func ExampleNewLexerFromYamlConfig() {
	// The yaml config source
	// strict_mode is true by default
	yamlSource := `
version : "0.0.1"
lexer:
  rules :
    - type  : "<var_kw>"
      regex : "var"
    - type  : "<eq>"
      regex : "="
    - type  : "<integer>"
      regex : "[0-9]+"
  ignore :
    - "\\s+"
`
	source := "var = 123"
	// try to generate a lexer from the given source and yaml config
	lex, err := goply.NewLexerFromYamlConfig([]byte(yamlSource))
	if err != nil {
		panic(err)
	}
	// get the tokens
	tokens, err := lex.GetTokenStream(source)
	if err != nil {
		panic(err)
	}
	// print out the tokens
	for token := range tokens.Iter() {
		fmt.Printf("Got %s : %s\n", token.Type, token.Value)
	}
	// Output:
	// Got <var_kw> : var
	// Got <eq> : =
	// Got <integer> : 123
}

func ExampleNewLexerFromYamlConfig_lenient() {
	// The yaml config source
	// The strict_mode field sets the strictness of the lexer
	// it is true by default
	yamlSource := `
version : "0.0.1"
lexer:
  strict_mode : false
  rules :
    - type  : "<var_kw>"
      regex : "var"
    - type  : "<eq>"
      regex : "="
    - type  : "<integer>"
      regex : "[0-9]+"
  ignore :
    - "\\s+"
`
	source := "var = 123"
	// try to generate a lexer from the given source and yaml config
	lex, err := goply.NewLexerFromYamlConfig([]byte(yamlSource))
	if err != nil {
		panic(err)
	}
	// get the tokens
	tokens, err := lex.GetTokenStream(source)
	if err != nil {
		panic(err)
	}
	// print out the tokens
	for token := range tokens.Iter() {
		fmt.Printf("Got %s : %s\n", token.Type, token.Value)
	}
	// Output:
	// Got <var_kw> : var
	// Got <eq> : =
	// Got <integer> : 123
}
