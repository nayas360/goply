package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/nayas360/goply"
)

var source = `(+ 10 (- 15 5))`

func main() {
	yamlConfig, err := getYamlConfig()
	if err != nil {
		panic(err)
	}
	// get the lexer from the config and source
	lex, err := goply.NewLexerFromYamlConfig(yamlConfig)
	if err != nil {
		panic(err)
	}
	// get the tokens
	tokens, err := lex.GetTokenStream(source)
	if err != nil {
		panic(err)
	}
	// print the tokens
	for token := range tokens.Iter() {
		fmt.Printf("Got %s : %s\n", token.Type, token.Value)
	}
}

// assumes GOPATH="GOPLY_PATH:OTHER_PATH"
// $GOPATH/src/github.com/nayas360/goply/examples/minilisp/lex.yml should exist
func getYamlConfig() ([]byte, error) {
	var GOP string

	for _, p := range os.Environ() {
		ps := strings.Split(p, "=")
		if ps[0] == "GOPATH" {
			GOP = strings.Split(ps[1], ":")[0]
		}
	}

	// load lexer definition from file
	ycp, err := filepath.Abs(GOP + "/src/github.com/nayas360/goply/examples/minilisp/lex.yml")
	if err != nil {
		panic(err)
	}
	// read the config
	yamlConfig, err := ioutil.ReadFile(ycp)
	return yamlConfig, err
}
