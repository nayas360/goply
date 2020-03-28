package expression

import (
	"github.com/nayas360/goply"
	"github.com/nayas360/goply/examples/algebrica/ast"
	"github.com/nayas360/goply/examples/algebrica/ast/node"
)

type Operator struct {
	E node.Expression
}

func (o *Operator) Parse(tokens *goply.TokenStream) error {
	if tokens.EOS() {
		return ast.EndOfTokenStreamErr
	}
	if o.E != nil {
		o.E = &Binary{Left: o.E}
	} else {
		// no left context unary operator
		o.E = &Binary{}
	}
	return o.E.Parse(tokens)
}

func (o *Operator) String() string {
	return o.E.String()
}

func (o *Operator) Expression() {}
func (o *Operator) Operator()   {}
