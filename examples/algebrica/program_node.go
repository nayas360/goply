package main

import (
	"fmt"

	"github.com/nayas360/goply"
)

type Program struct {
	Expr Expression
}

func (p *Program) Parse(tokens []*goply.Token, id int) error {
	p.Expr = &GenericExpression{}
	return p.Expr.Parse(tokens, id)
}

func (p *Program) String() string {
	return fmt.Sprintf("Program { %s }", p.Expr)
}
