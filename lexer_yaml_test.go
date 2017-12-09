package goply_test

import (
	"fmt"
	"testing"

	"github.com/nayas360/goply"
)

func TestNewLexerFromYamlConfig(t *testing.T) {
	yamlSource := `
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
	lex, err := goply.NewLexerFromYamlConfig([]byte(yamlSource), source)
	if err != nil {
		t.Errorf("could not create a new lexer from yaml config, %s", err)
	}
	tokens, err := lex.GetTokens()
	if err != nil {
		t.Errorf("got error instead of tokens, %s", err)
	}
	fmt.Println(tokens)
	if len(tokens) != 3 {
		t.Errorf("expected 3 tokens got, %s", len(tokens))
	}
}
