package goply_test

import (
	"fmt"
	"testing"

	"github.com/nayas360/goply"
)

func TestLexer(t *testing.T) {
	source := `func test() {
	num := 123
	var num2 = 123
}
`
	lexer := goply.NewLexer(source)
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

	tokens, err := lexer.GetTokens()
	if err != nil {
		t.Errorf("got error instead of tokens, %s", err)
	}

	if len(tokens) != 13 {
		t.Error("expected 13 tokens got,", len(tokens))
	}
}

func TestLexer_SetLexerErrorFunc(t *testing.T) {
	lexer := goply.NewLexer("123")
	lexer.SetLexerErrorFunc(func(ls goply.LexerState) error {
		return fmt.Errorf("there was an error")
	})
	_, err := lexer.GetTokens()
	if fmt.Sprint(err) != "there was an error" {
		t.Error("the custom error function was not set")
	}
}
