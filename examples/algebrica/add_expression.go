package main

import (
	"github.com/nayas360/goply"
	"golang.org/x/exp/errors/fmt"
)

type AddOp struct {
	Token *goply.Token
	Left  Expression
	Right Expression
}

func (o *AddOp) Parse(tokens []*goply.Token, id int) error {
	if id >= len(tokens) {
		return fmt.Errorf("reached end of token stream")
	}
	if tokens[id].Type != AddOpT {
		return fmt.Errorf("invalid literal '%s' at line %d column %d", tokens[id].Value)
	}
	if id-1 < 0 {
		return fmt.Errorf("cannot start with binary operator %d", tokens[id].Value)
	}
	if tokens[id-1].Type != NumberT {
		return fmt.Errorf("invalid operand '%s' at line %d column %d", tokens[id-1].Value)
	}
	o.Token = tokens[id]
	o.Left = &Number{}
	err := o.Left.Parse(tokens, id-1)
	if err != nil {
		return err
	}
	o.Right = &GenericExpression{}
	return o.Right.Parse(tokens, id+1)
}

func (o *AddOp) String() string {
	return fmt.Sprintf("%s + %s", o.Left, o.Right)
}

func (o *AddOp) Expression() {}
