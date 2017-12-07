package goply

import (
	"fmt"
	"regexp"
	"strings"
)

var newlineChars = regexp.MustCompile("\n")

// the lexer struct
type Lexer struct {
	ls               LexerState                // internal state of the lexer
	lexRules         map[string]*regexp.Regexp // mapping from Type names to regex Rules to be used with a token
	lexRulesKeyOrder []string                  // slice of keys for predictable iteration over lexRules
	ignoreRules      []*regexp.Regexp          // regular expressions to be ignored
	lexerErrorFunc   func(ls LexerState) error // func to call for error
}

// Create a new lexer for a given source string
func NewLexer(source string) *Lexer {
	return &Lexer{ls: LexerState{SourceLength: len(source) - 1, Source: source},
		lexRules: make(map[string]*regexp.Regexp), lexerErrorFunc: defaultLexerError}
}

// Tokens are returned only for these lexRules
func (l *Lexer) AddRule(tokenType, regexv string) {
	// "^" is added as a prefix to all the regular expressions to match at the front
	l.lexRules[tokenType] = regexp.MustCompile("^" + regexv)
	l.lexRulesKeyOrder = append(l.lexRulesKeyOrder, tokenType)
}

// Tokens are not created for these regular expressions
func (l *Lexer) Ignore(regexv string) {
	// "^" is added as a prefix to all the regular expressions to match at the front
	l.ignoreRules = append(l.ignoreRules, regexp.MustCompile("^"+regexv))
}

// returns a slice of tokens to
func (l *Lexer) GetTokens() ([]*Token, error) {
	// build the slice of tokens
	var tokens []*Token
	token, err := l.next()
	if err != nil {
		return nil, err
	}
	for ; token != nil; {
		tokens = append(tokens, token)
		token, err = l.next()
		if err != nil {
			return nil, err
		}
	}
	return tokens, nil
}

// This function is used to replace the default error handler
func (l *Lexer) SetLexerErrorFunc(f func(ls LexerState) error) {
	l.lexerErrorFunc = f
}

// The default error handler
func defaultLexerError(ls LexerState) error {
	return fmt.Errorf("could not match '%c'@%d with any rule", ls.Source[ls.Position], ls.Position)
}

// returns the next token from the source
func (l *Lexer) next() (*Token, error) {
	if l.ls.Position <= l.ls.SourceLength {

		// go through all the ignored lexRules
		for _, lexRule := range l.ignoreRules {
			// check if there is a match
			if lexRule.MatchString(l.ls.Source[l.ls.Position:]) {
				// add the length of token to be ignored and skip by recursively calling myself
				l.ls.Position += len(lexRule.FindString(l.ls.Source[l.ls.Position:]))
				return l.next()
			}
		}

		// go through all the lexRules to tokenize
		for _, tokenType := range l.lexRulesKeyOrder {
			lexRule := l.lexRules[tokenType]
			if lexRule.MatchString(l.ls.Source[l.ls.Position:]) {
				value := lexRule.FindString(l.ls.Source[l.ls.Position:])
				l.ls.LineNum = strings.Count(l.ls.Source[:l.ls.Position], "\n")
				newLineIndex := newlineChars.FindAllStringIndex(l.ls.Source[:l.ls.Position], l.ls.LineNum)
				//colNum := l.ls.Position
				if len(newLineIndex) > 0 {
					//fmt.Println(newLineIndex[len(newLineIndex)-1][0] - 1)
					l.ls.ColNum = l.ls.Position - newLineIndex[len(newLineIndex)-1][0] - 1
				} else {
					l.ls.ColNum = l.ls.Position
				}
				//l.ls.ColNum = colNum
				token := newToken(tokenType, value, l.ls.Position, l.ls.LineNum, l.ls.ColNum)
				// after processing add to the curpos
				l.ls.Position += len(value)
				return token, nil
			}
		}

		// If here then Could not match anything
		return nil, l.lexerErrorFunc(l.ls)
	} else {
		return nil, nil
	}
}
