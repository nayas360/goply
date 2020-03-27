package main

import (
	"fmt"

	"github.com/nayas360/goply"
)

type GenericExpression struct {
	expr Expression
}

func (g *GenericExpression) Parse(tokens []*goply.Token, id int) error {
	if id >= len(tokens) {
		return fmt.Errorf("reached end of token stream")
	}
	if tokens[id].Type == NumberT && id+1 < len(tokens) && tokens[id+1].Type == AddOpT {
		g.expr = &AddOp{}
		return g.expr.Parse(tokens, id+1)
	} else if tokens[id].Type == NumberT && id+1 >= len(tokens) {
		g.expr = &Number{}
		return g.expr.Parse(tokens, id)
	}
	return fmt.Errorf("invalid literal '%s' at line %d column %d",
		tokens[id].Value, tokens[id].LineNum, tokens[id].ColNum)
}

func (g *GenericExpression) String() string {
	if g.expr != nil {
		return g.expr.String()
	}
	return ""
}

func (g *GenericExpression) Expression() {}
