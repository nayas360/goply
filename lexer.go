// package name: goply
package goply

import (
	"fmt"
	"regexp"
	"strings"
)

//var nullToken = &Token{"", "", 0, 0, 0, 0}

var newlineChars = regexp.MustCompile("\n")

// the lexer struct
type Lexer struct {
	sourceLength int                       // the length of the source string
	source       string                    // the source string itself
	rules        map[*regexp.Regexp]string // mapping from regex Rules to Type names to be used with a token
	ignoreRules  []*regexp.Regexp          // regular expressions to be ignored
	curPosition  int                       // current position in the source string
	curLineNum   int                       // current Line number
	curColNum    int                       // current column number
	/*lexerErrorFunc func(l Lexer) error*/ // func to call for error
}

// Create a new lexer for a given source string
func NewLexer(source string) *Lexer {
	rules := make(map[*regexp.Regexp]string)
	return &Lexer{sourceLength: len(source) - 1, source: source, rules: rules /*, lexerErrorFunc: defaultLexerError*/ }
}

// Tokens are returned only for these rules
func (l *Lexer) AddRule(Type, regexVal string) {
	// "^" is added as a prefix to all the regular expressions to match at the front
	l.rules[regexp.MustCompile("^"+regexVal)] = Type
}

// Tokens are not created for these regular expressions
func (l *Lexer) Ignore(regexVal string) {
	// "^" is added as a prefix to all the regular expressions to match at the front
	l.ignoreRules = append(l.ignoreRules, regexp.MustCompile("^"+regexVal))
}

// returns a slice of tokens to
func (l *Lexer) GetTokens() ([]*Token, error) {
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

/*func (l *Lexer) lexerError() error {
	return l.lexerErrorFunc(l)
}

func (l Lexer) SetLexerErrorFunc(f func(l Lexer) error) {
	l.lexerErrorFunc = f
}

func defaultLexerError(l Lexer) error {
	return fmt.Errorf("could not match '%c'@%d with any rule", l.source[l.curPosition], l.curPosition)
}*/

func (l *Lexer) next() (*Token, error) {
	if l.curPosition <= l.sourceLength {

		// go through all the ignored rules
		for _, rule := range l.ignoreRules {
			// check if there is a match
			if rule.MatchString(l.source[l.curPosition:]) {
				// add the length of token to be ignored and skip by recursively calling myself
				l.curPosition += len(rule.FindString(l.source[l.curPosition:]))
				return l.next()
			}
		}

		// go through all the rules to tokenize
		for rule, Type := range l.rules {
			if rule.MatchString(l.source[l.curPosition:]) {
				value := rule.FindString(l.source[l.curPosition:])
				lineNum := strings.Count(l.source[:l.curPosition], "\n")
				newLineIndex := newlineChars.FindAllStringIndex(l.source[:l.curPosition], lineNum)
				colNum := l.curPosition
				if len(newLineIndex) > 0 {
					//fmt.Println(newLineIndex[len(newLineIndex)-1][0] - 1)
					colNum = l.curPosition - newLineIndex[len(newLineIndex)-1][0] - 1
				}
				l.curLineNum = lineNum
				l.curColNum = colNum
				token := newToken(Type, value, l.curPosition, lineNum, colNum)
				// after processing add to the curpos
				l.curPosition += len(value)
				return token, nil
			}
		}

		// If here then Could not match anything
		return nil, fmt.Errorf("could not match '%c'@%d with any rule", l.source[l.curPosition], l.curPosition)
	} else {
		return nil, nil
	}
}
