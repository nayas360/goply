package node

import (
	"github.com/nayas360/goply"
)

type Expression interface {
	goply.AstNode
	Expression()
}
