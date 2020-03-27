package main

import (
	"fmt"
	"log"

	"github.com/nayas360/goply"
)

func main() {
	l := goply.NewLexer(true)
	// the add operator
	l.AddRule(AddOpT, "\\+")
	// some numbers
	l.AddRule(NumberT, "[0-9]+")
	l.Ignore("\\s+")

	p := goply.NewParser(l)
	ast := &Program{}
	err := p.Parse("1 + 2 + 3", ast)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ast)
}
