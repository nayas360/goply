package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nayas360/goply"
	"github.com/nayas360/goply/examples/algebrica/ast/program"
	"github.com/nayas360/goply/examples/algebrica/ast/types"
)

const lexRulesYamlConfig = `
version : "0.0.1"
lexer:
  rules :
    - type  : "%s"
      regex : "\\("
    - type  : "%s"
      regex : "\\)"
    - type  : "%s"
      regex : "\\+"
    - type  : "%s"
      regex : "-"
    - type  : "%s"
      regex : "\\*"
    - type  : "%s"
      regex : "/"
    - type  : "%s"
      regex : "[0-9]+"
  ignore :
    - "\\s+"
`

func main() {
	l, err := goply.NewLexerFromYamlConfig([]byte(fmt.Sprintf(lexRulesYamlConfig,
		types.LParenT, types.RParenT, types.AddOpT, types.SubOpT, types.MulOpT, types.DivOpT,
		types.NumberT)))
	if err != nil {
		log.Fatal(err)
	}
	p := goply.NewParser(l)
	a := &program.Program{}
	err = p.Parse(strings.Join(os.Args[1:], ""), a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}
