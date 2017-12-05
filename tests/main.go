package main

import (
	"fmt"

	"github.com/nayas360/goply"
)

var source = `func test() {
	number := 123
}
`

//var source = "root test = true {\n\tnode {}\n}"

func main() {
	l := goply.NewLexer(source)
	l.AddRule("<lparen>", "\\(")
	l.AddRule("<rparen>", "\\)")
	l.AddRule("<lbrace>", "{")
	l.AddRule("<rbrace>", "}")
	l.AddRule("<assign>", ":=")
	l.AddRule("<identifier>", "[A-Za-z_][A-Za-z0-9]+")
	l.AddRule("<number>", "[0-9]+")
	l.Ignore("\\s+")
	/*l.SetLexerErrorFunc(func(l lexer.Lexer) error {
		return fmt.Errorf("there was an error, %v", l)
	})*/

	tokens, err := l.GetTokens()
	if err != nil {
		panic(err)
	}

	for _, t := range tokens {
		fmt.Printf("Got %s : '%s'(%d,%d) @{%d, %d}\n", t.Type, t.Value, t.StartingPosition,
			t.StartingPosition+t.Length-1, t.LineNum, t.ColNum)
	}
}