package goply_test

import (
	"fmt"
	"testing"

	"github.com/nayas360/goply"
)

// tests lexer in lenient mode
func TestNewLexer(t *testing.T) {
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

	// lexer is in lenient mode,
	// the following should not raise an error
	//lexer.AddRule("<assign>", ":=")
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

// tests lexer in strict mode
func TestNewLexerStrict(t *testing.T) {
	source := "num := 123"
	lexer := goply.NewLexerStrict(source)

	// the lexer is in strict mode,
	// and cannot match the ':' character to any rule
	// this should raise an error
	//lexer.AddRule("<assign>", ":=")

	lexer.AddRule("<identifier>", "[A-Za-z_][A-Za-z0-9]+")
	lexer.AddRule("<number>", "[0-9]+")

	lexer.Ignore("\\s+")

	_, err := lexer.GetTokens()
	if err == nil {
		t.Errorf("expected an error, got none")
	}
}

// tests custom error handler in strict mode
func TestLexerStrict_SetLexerErrorFunc(t *testing.T) {
	lexer := goply.NewLexerStrict("123")
	lexer.SetLexerErrorFunc(func(ls goply.LexerState) error {
		return fmt.Errorf("there was an error")
	})
	_, err := lexer.GetTokens()
	if fmt.Sprint(err) != "there was an error" {
		t.Error("the custom error function was not set")
	}
}

// tests custom error handler in lenient mode
// the error function is ignored since no error is raised
func TestLexer_SetLexerErrorFunc(t *testing.T) {
	lexer := goply.NewLexer("123")
	// the lexer error functions are ignored in lenient mode
	lexer.SetLexerErrorFunc(func(ls goply.LexerState) error {
		return fmt.Errorf("there was an error")
	})
	_, err := lexer.GetTokens()
	if err != nil {
		t.Error("the lexer returned an error in lenient mode")
	}
}
