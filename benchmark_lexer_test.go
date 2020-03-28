package goply_test

import (
	"fmt"
	"testing"

	"github.com/nayas360/goply"
)

func BenchmarkNewLexer(b *testing.B) {
	// the source would repeat after 100 iterations
	// caching should take place
	b.ReportAllocs()
	repeatAfter := 100
	lexer := goply.NewLexer(true)
	lexer.AddRule("<number>", "[0-9]+")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tokens, err := lexer.GetTokenStream(fmt.Sprintf("%d", b.N%repeatAfter))
		if err != nil {
			b.Errorf("got error instead of tokens, %s", err)
		}
		if tokens.Len() != 1 {
			b.Error("expected 1 tokens got,", tokens.Len())
		}
	}
}
