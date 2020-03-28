package goply

type Parser struct {
	lex *Lexer // the lexer
}

func (p Parser) Parse(source string, root AstNode) error {
	tokens, err := p.lex.GetTokenStream(source)
	if err != nil {
		return err
	}
	return root.Parse(tokens)
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{
		lex: lexer,
	}
}
