package expression

import (
	"github.com/nayas360/goply"
	"github.com/nayas360/goply/examples/algebrica/ast"
	"github.com/nayas360/goply/examples/algebrica/ast/node"
)

type Operator struct {
	E node.Expression
}

func (o Operator) Parse(tokens *goply.TokenStream) error {
	if tokens.EOS() {
		return ast.EndOfTokenStreamErr
	}
	var op node.Operator
	if o.E != nil {
		op = &Binary{Left: o.E}
	} else {
		op = &Binary{}
	}
	err := op.Parse(tokens)
	if err != nil {
		return err
	}
	o.E = op
	return nil
}

func (o Operator) String() string {
	return o.E.String()
}

func (o Operator) Expression() {}
func (o Operator) Operator()   {}
