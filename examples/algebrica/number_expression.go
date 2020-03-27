package main

import (
	"fmt"
	"strconv"

	"github.com/nayas360/goply"
)

type Number struct {
	Token *goply.Token
	Value uint64
}

func (n *Number) Parse(tokens []*goply.Token, id int) error {
	if id >= len(tokens) {
		return fmt.Errorf("reached end of token stream")
	}
	n.Token = tokens[id]
	val, err := strconv.ParseUint(n.Token.Value, 10, 64)
	if err != nil {
		return err
	}
	n.Value = val
	return nil
}

func (n *Number) String() string {
	return n.Token.Value
}

func (n *Number) Expression() {}
