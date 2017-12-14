package goply_test

import (
	"testing"

	"github.com/nayas360/goply"
)

func TestNewLexerFromYamlConfig(t *testing.T) {
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
	lex, err := goply.NewLexerFromYamlConfig([]byte(yamlSource))
	if err != nil {
		t.Errorf("could not create a new lexer from yaml config, %s", err)
	}
	tokens, err := lex.GetTokens(source)
	if err != nil {
		t.Errorf("got error instead of tokens, %s", err)
	}
	if len(tokens) != 3 {
		t.Errorf("expected 3 tokens got, %s", len(tokens))
	}
}
