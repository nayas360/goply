package goply_test

import (
	"testing"

	"github.com/nayas360/goply"
)

func TestTokenStream_PosDelta(t *testing.T) {
	l := makeLexer()
	tokens, err := l.GetTokenStream("123 456")
	if err != nil {
		t.Error(err)
	}
	if tokens.Pos() != 0 {
		t.Error("pos should start from 0")
	}
	tokens.PosDelta(1)
	if tokens.Pos() != 1 {
		t.Error("pos should be at 1")
	}
	if tokens.PosDelta(1).Pos() != 2 {
		t.Error("pos should be at 2")
	}
	if !tokens.EOS() {
		t.Error("should be end of stream now")
	}
	tokens.ResetPos()
	if tokens.Pos() != 0 {
		t.Error("should be 0")
	}
	if tokens.PosDelta(2).ResetPos().Pos() != 0 {
		t.Error("should be 0")
	}
	if !tokens.PosDelta(2).EOS() {
		t.Error("should be end of stream now")
	}
	tokens.PosDelta(-1)
	if tokens.Pos() != 1 {
		t.Error("pos should be at 1")
	}
	if tokens.PosDelta(-1).Pos() != 0 {
		t.Error("pos should be at 0")
	}
}

func makeLexer() *goply.Lexer {
	l := goply.NewLexer(true)
	l.AddRule("num", "[0-9]+")
	l.Ignore("\\s+")
	return l
}
