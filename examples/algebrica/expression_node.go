package main

import (
	"github.com/nayas360/goply"
)

type Expression interface {
	goply.AstNode
	Expression()
}
