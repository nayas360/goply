package main

import (
	"fmt"

	"github.com/nayas360/goply"
)

var source = `func test() {
	num := 123
	var num2 := 123
}
`

func main() {
	l := goply.NewLexerStrict(source)
	l.AddRule("<lparen>", "\\(")
	l.AddRule("<rparen>", "\\)")
	l.AddRule("<lbrace>", "{")
	l.AddRule("<rbrace>", "}")

	l.AddRule("<assign>", ":=")
	l.AddRule("<eq>", "=")

	l.AddRule("<func_kw>", "func")
	l.AddRule("<var_kw>", "var")

	l.AddRule("<identifier>", "[A-Za-z_][A-Za-z0-9]+")
	l.AddRule("<number>", "[0-9]+")

	l.Ignore("\\s+")

	tokens, err := l.GetTokens()
	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		fmt.Printf("Got %s : \"%s\" on line %d column %d \n", token.Type, token.Value, token.LineNum,
			token.ColNum)
	}
}
