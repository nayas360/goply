package goply

import (
	"regexp"
	"strings"
)

// Struct having related fields for representing a lexer
type Lexer struct {
	ls               LexerState                // internal state of the lexer
	lexRules         map[string]*regexp.Regexp // mapping from Type names to regex Rules to be used with a token
	lexRulesKeyOrder []string                  // slice of keys for predictable iteration over lexRules
	ignoreRules      []*regexp.Regexp          // regular expressions to be ignored
	lexerErrorFunc   func(ls LexerState) error // func to call for error
	strictMode       bool                      // if true, returns error if no rules can be matched
	tokenCache       map[string][]*Token       // a cache of token slices given a cacheID
}

// Create a new lexer
// if strictMode is false
// the lexer skips over characters that cannot be matched by any rule
// else generates an error for the unmatched symbol
func NewLexer(strictMode bool) *Lexer {
	return &Lexer{ls: LexerState{},
		lexRules: make(map[string]*regexp.Regexp), tokenCache: make(map[string][]*Token),
		lexerErrorFunc: defaultLexerError, strictMode: strictMode}
}

// When processing the source, all patterns matched by the regex
// generates a token with the Token.Type as the tokenType and
// the Token.Value as the pattern that was matched.
func (l *Lexer) AddRule(tokenType, regexv string) {
	// "^" is added as a prefix to all the regular expressions to match at the front
	l.lexRules[tokenType] = regexp.MustCompile("^" + regexv)
	l.lexRulesKeyOrder = append(l.lexRulesKeyOrder, tokenType)
}

// When processing the source,
// all patterns matched by the regex are skipped over.
func (l *Lexer) Ignore(regexv string) {
	// "^" is added as a prefix to all the regular expressions to match at the front
	l.ignoreRules = append(l.ignoreRules, regexp.MustCompile("^"+regexv))
}

// Processes the source text and returns the tokens
func (l *Lexer) GetTokens(sourceText string) ([]*Token, error) {
	// compute sha1 of source text for caching
	sourceSha1 := computeSha1(sourceText)

	// check if tokens for sourceText exists or not
	if l.tokenCache[sourceSha1] != nil {
		// if exists then return
		return l.tokenCache[sourceSha1], nil
	} else {
		l.ls.Source = sourceText
		l.ls.SourceLength = len(sourceText) - 1
		l.ls.Position = 0
		l.ls.LineNum = 0
		l.ls.ColNum = 0
	}
	// build the slice of tokens
	var tokens []*Token
	for token, err := l.nextToken(); ; token, err = l.nextToken() {
		if err != nil {
			return nil, err
		} else if token == nil {
			break
		}
		tokens = append(tokens, token)
	}
	// store the tokens in the cache
	l.tokenCache[sourceSha1] = tokens
	return tokens, nil
}

// Set a custom error handler for the lexer
func (l *Lexer) SetLexerErrorFunc(f func(ls LexerState) error) {
	l.lexerErrorFunc = f
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
		if l.strictMode {
			// strict mode enabled and could not match anything
			return nil, l.lexerErrorFunc(l.ls)
		} else {
			// strict mode disabled skip over unmatched chars one by one
			l.ls.Position += 1
			l.updateLexerState()
			return l.nextToken()
		}
	} else {
		return nil, nil
	}
}

// regex for newline characters
// used only by updateLexerState
var newlineChars = regexp.MustCompile("\n")

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
