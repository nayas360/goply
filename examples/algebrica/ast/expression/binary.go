package expression

import (
	"github.com/nayas360/goply"
	"github.com/nayas360/goply/examples/algebrica/ast"
	"github.com/nayas360/goply/examples/algebrica/ast/node"
	"golang.org/x/exp/errors/fmt"
)

type Binary struct {
	Token *goply.Token
	Left  node.Expression
	Right node.Expression
}

func (o *Binary) Parse(tokens *goply.TokenStream) error {
	if tokens.EOS() {
		return ast.EndOfTokenStreamErr
	}
	o.Token = tokens.Get()
	rExpr := &Expression{}
	err := rExpr.Parse(tokens.PosDelta(1))
	if err != nil {
		return err
	}
	o.Right = rExpr.E
	return nil
}

func (o *Binary) String() string {
	return fmt.Sprintf("(%s %s %s)", o.Left, o.Token.Value, o.Right)
}

func (o *Binary) Expression() {}
func (o *Binary) Operator()   {}
