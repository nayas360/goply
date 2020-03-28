package program

import (
	"github.com/nayas360/goply"
	"github.com/nayas360/goply/examples/algebrica/ast/expression"
	"github.com/nayas360/goply/examples/algebrica/ast/node"
)

type Program struct {
	E node.Expression
}

func (p *Program) Parse(tokens *goply.TokenStream) error {
	expr := &expression.Expression{}
	err := expr.Parse(tokens)
	if err != nil {
		return err
	}
	p.E = expr.E
	return nil
}

func (p *Program) String() string {
	return p.E.String()
}
