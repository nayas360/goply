package goply

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	source := `func test() {
	num := 123
	var num2 = 123
}
`
	lexer := NewLexer(source)
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
	lexer := NewLexer("123")
	lexer.SetLexerErrorFunc(func(ls LexerState) error {
		return fmt.Errorf("there was an error")
	})
	_, err := lexer.GetTokens()
	if fmt.Sprint(err) != "there was an error" {
		t.Error("the custom error function was not set")
	}
}

func TestLexer_NextToken(t *testing.T) {
	lexer := NewLexer("\t\ntest 123\n")
	lexer.AddRule("<test_kw>", "test")
	lexer.AddRule("<number>", "[0-9]+")
	lexer.Ignore("\\s+")

	token, err := lexer.nextToken()
	if err != nil {
		t.Errorf("could not get next token,%s", err)

	}
	if token.Type != "<test_kw>" && token.Value != "test" {
		t.Errorf("expected token.Type = '<test_kw>' and token.Value = 'test',"+
			"got token.Type = %s and token.Value = %s", token.Type, token.Value)
	}

	token, err = lexer.nextToken()
	if err != nil {
		t.Errorf("could not get next token,%s", err)

	}
	if token.Type != "<number>" && token.Value != "123" {
		t.Errorf("expected token.Type = '<number>' and token.Value = '123',"+
			"got token.Type = %s and token.Value = %s", token.Type, token.Value)
	}
}

func BenchmarkNewLexer(b *testing.B) {
	source := `func test() {
	num := 123
	var num2 = 123
}
`
	for i := 0; i < b.N; i++ {
		lexer := NewLexer(source)
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
			b.Errorf("got error instead of tokens, %s", err)
		}
		if len(tokens) != 13 {
			b.Error("expected 13 tokens got,", len(tokens))
		}
	}
}

func ExampleLexer() {
	// sample lisp like code
	source := "(+ 10 20)"
	lexer := NewLexer(source)
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
