package ast

import (
	"errors"
)

var (
	EndOfTokenStreamErr    = errors.New("unexpectedly reached end of token stream")
	UnexpectedTokenErr     = errors.New("unexpected token encountered")
	UnexpectedTokenTypeErr = errors.New("unexpected token type encountered")
)
