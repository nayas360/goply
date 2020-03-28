package expression

import (
	"github.com/nayas360/goply"
	"github.com/nayas360/goply/examples/algebrica/ast"
	"github.com/nayas360/goply/examples/algebrica/ast/node"
	"github.com/nayas360/goply/examples/algebrica/ast/types"
	"github.com/pkg/errors"
)

type Expression struct {
	E node.Expression
}

func (e *Expression) Parse(tokens *goply.TokenStream) error {
	if tokens.EOS() {
		return ast.EndOfTokenStreamErr
	}
	switch types.NodeType(tokens.Get().Type) {
	case types.AddOpT, types.SubOpT, types.MulOpT, types.DivOpT, types.LParenT, types.RParenT:
		o := &Operator{E: e.E}
		err := o.Parse(tokens)
		if err != nil {
			return err
		}
		e.E = o.E
		return nil
	case types.NumberT:
		e.E = &Number{}
		err := e.E.Parse(tokens)
		if err != nil {
			return err
		}
		// advance if not the last or only token
		if !tokens.PosDelta(1).EOS() {
			return e.Parse(tokens)
		}
		return nil
	default:
		return errors.WithMessagef(ast.UnexpectedTokenTypeErr,
			"[error (%d:%d)]: invalid literal '%s' of unknown type '%s'",
			tokens.Get().LineNum, tokens.Get().ColNum, tokens.Get().Value, tokens.Get().Type)
	}
}

func (e *Expression) String() string {
	return e.E.String()
}

func (e *Expression) Expression() {}
