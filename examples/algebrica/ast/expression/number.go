package expression

import (
	"strconv"

	"github.com/nayas360/goply"
	"github.com/nayas360/goply/examples/algebrica/ast"
	"github.com/nayas360/goply/examples/algebrica/ast/types"
	"github.com/pkg/errors"
)

type Number struct {
	Token *goply.Token
	Value uint64
}

func (n *Number) Parse(tokens *goply.TokenStream) error {
	if tokens.EOS() {
		return ast.EndOfTokenStreamErr
	}
	if !types.NumberT.Equals(tokens.Get().Type) {
		return errors.WithMessagef(ast.UnexpectedTokenErr,
			"[%d:%d]: expected literal of type '%s', got '%s' instead",
			n.Token.LineNum, n.Token.ColNum, types.NumberT, n.Token.Type)
	}
	n.Token = tokens.Get()
	val, err := strconv.ParseUint(n.Token.Value, 10, 64)
	if err != nil {
		return errors.WithMessagef(err,
			"[%d:%d]: exception occurred while parsing integer '%s'",
			n.Token.LineNum, n.Token.ColNum, n.Token.Value)
	}
	n.Value = val
	return nil
}

func (n *Number) String() string {
	return n.Token.Value
}

func (n *Number) Expression() {}
