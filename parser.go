package goply

import "fmt"

// production rules lexer specific constants
const (
	plx_TERMINAL_B   = "<terminal_"
	plx_TERMINAL_E   = ">"
	plx_NON_TERMINAL = "<non_terminal>"
	plx_SEPARATOR    = "<colon>"
	plx_SEPARATOR_V  = ":"
	plx_WHITESPACE   = "[ \t]+"
)

type Parser struct {
	slx       *Lexer              // this performs lexing on the source
	plx       *Lexer              // this performs lexing on the production rules
	sProdTok  string              // the first non-terminal token, everything should reduce to this
	prodRules map[string][]*Token // mapping from nonterminal RHS token to slice of tokens in LHS
}

// Construct and make a new parser from an existing lexer
// Overrides defaultLexerError
func NewParserFromLexer(lex *Lexer) *Parser {
	// create a new lexer for production rules
	pLex := NewLexer(true)
	// set custom lexer error
	pLex.SetLexerErrorFunc(plxDefaultErrorFunc)
	// add slx types as keywords that are used in production
	// rules as terminals
	for _, Type := range lex.lexRulesKeyOrder {
		// add the types of all ruels to the production lexer to make them unique
		pLex.AddRule(plx_TERMINAL_B+Type+plx_TERMINAL_E, Type)
	}
	// add <non_terminals> to identify non terminal symbols
	// general variable naming rules apply
	pLex.AddRule(plx_NON_TERMINAL, "[_A-Za-z][_A-Za-z0-9]+")
	pLex.AddRule(plx_SEPARATOR, plx_SEPARATOR_V)
	pLex.Ignore(plx_WHITESPACE)

	return &Parser{slx: lex, plx: pLex, prodRules: make(map[string][]*Token)}
}

// adds a rule to the parser
func (p *Parser) AddRule(expr string) {
	tokens, err := p.plx.GetTokens(expr)
	if err != nil {
		panic(err)
	}
	// position one has to be a plx_SEPARATOR
	if tokens[1].Type != plx_SEPARATOR {
		panic(fmt.Errorf("expected ':' after %s at position %d in '%s'", tokens[0].Value,
			tokens[0].StartingPosition+tokens[0].Length+1, expr))
	}

	for id, token := range tokens {
		// this should be the RHS, before plx_SEPARATOR
		if id == 0 {
			// the token on the right hand side cannot be a terminal token
			if isTerminalType(token.Type) {
				panic(fmt.Errorf("%s is a terminal token, it must not be in right hand side"))
			}
			// this should only be set if it was previously unset
			if p.sProdTok == "" {
				p.sProdTok = token.Value
			}
			fmt.Printf("RHS(%s) :", token.Value)
			continue
		}
		if id == 1 { // skip plx_SEPARATOR
			continue
		}
		// this should be LHS, after plx_SEPARATOR
		fmt.Printf(" LHS(%s, %s)", token.Type, token.Value)
	}
	fmt.Println()
}
