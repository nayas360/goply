package main

import (
	"github.com/nayas360/goply"
)

func main() {

	//source := "(+ 1 2)"

	lex := goply.NewLexer(true)
	attachLexerRules(lex)
	pr := goply.NewParserFromLexer(lex)
	pr.AddRule("program : expr")
	pr.AddRule("expr : <lparen> expr <rparen>")
	pr.AddRule("expr : <op_plus> expr expr")
	pr.AddRule("expr : <integer>")
}

// attach the rules to a lexer in another function
// to simplify main
func attachLexerRules(lex *goply.Lexer) {
	lex.AddRule("<lparen>", "\\(")
	lex.AddRule("<rparen>", "\\)")
	lex.AddRule("<integer>", "\\d+")
	lex.AddRule("<op_plus>", "\\+")
}
