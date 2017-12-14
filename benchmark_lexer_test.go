package goply_test

import (
	"testing"

	"github.com/nayas360/goply"
)

func BenchmarkNewLexer(b *testing.B) {
	source := `func test() {
	num := 123
	var num2 = 123
}
`
	for i := 0; i < b.N; i++ {
		lexer := goply.NewLexer(false)
		lexer.AddRule("<lparen>", "\\(")
		lexer.AddRule("<rparen>", "\\)")
		lexer.AddRule("<lbrace>", "{")
		lexer.AddRule("<rbrace>", "}")

		lexer.AddRule("<assign>", ":=")
		lexer.AddRule("<eq>", "=")

		lexer.AddRule("<func_kw>", "func")
		lexer.AddRule("<var_kw>", "var")

		lexer.AddRule("<identifier>", "[A-Za-z_][A-Za-z0-9]+")
		lexer.AddRule("<number>", "[0-9]+")

		lexer.Ignore("\\s+")
		tokens, err := lexer.GetTokens(source)
		if err != nil {
			b.Errorf("got error instead of tokens, %s", err)
		}
		if len(tokens) != 13 {
			b.Error("expected 13 tokens got,", len(tokens))
		}
	}
}
