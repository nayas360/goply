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
	token, err := l.nextToken()
	if err != nil {
		return nil, err
	}
	for token != nil {
		tokens = append(tokens, token)
		token, err = l.nextToken()
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
	return fmt.Errorf("line %d, column %d: could not match '%c' with any rule", ls.LineNum, ls.ColNum,
		ls.Source[ls.Position])
}

// returns the nextToken token from the source
func (l *Lexer) nextToken() (*Token, error) {
	if l.ls.Position <= l.ls.SourceLength {

		// go through all the ignored lexRules
		for _, lexRule := range l.ignoreRules {
			// check if there is a match
			if lexRule.MatchString(l.ls.Source[l.ls.Position:]) {
				// update the lexer state
				l.updateLexerState()
				// add the length of token to be ignored and skip by recursively calling myself
				l.ls.Position += len(lexRule.FindString(l.ls.Source[l.ls.Position:]))
				return l.nextToken()
			}
		}

		// go through all the lexRules to tokenize
		for _, tokenType := range l.lexRulesKeyOrder {
			lexRule := l.lexRules[tokenType]
			if lexRule.MatchString(l.ls.Source[l.ls.Position:]) {
				value := lexRule.FindString(l.ls.Source[l.ls.Position:])
				// update the lexer state
				l.updateLexerState()
				// create the token to return later
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

// calculates and updates the lexer state based on the current position in the source
func (l *Lexer) updateLexerState() {
	l.ls.LineNum = strings.Count(l.ls.Source[:l.ls.Position], "\n")
	newLineIndex := newlineChars.FindAllStringIndex(l.ls.Source[:l.ls.Position], l.ls.LineNum)
	if len(newLineIndex) > 0 {
		l.ls.ColNum = l.ls.Position - newLineIndex[len(newLineIndex)-1][0] - 1
	} else {
		l.ls.ColNum = l.ls.Position
	}
}
