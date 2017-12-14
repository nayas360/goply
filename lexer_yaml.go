package goply

import (
	"fmt"
	"regexp"

	"gopkg.in/yaml.v2"
)

var goplyConfVersion = "0.0.1"

// Struct used to read the yaml config into
type goplyYamlConfig struct {
	ConfVersion string `yaml:"version"`
	Lexer       struct {
		StrictMode bool `yaml:"strict_mode,omitempty"`
		Rules      []struct {
			Type  string `yaml:"type"`
			Regex string `yaml:"regex"`
		} `yaml:"rules"`
		Ignore []string `yaml:"ignore"`
	} `yaml:"lexer"`
}

// Create a lexer from a yaml config
// the config should be the config source and not a file path
// this allows loading the config from file as well as memory
// like the source file
// returns an error if could not read the yaml properly
func NewLexerFromYamlConfig(yamlConfig []byte, source string) (*Lexer, error) {
	var gyc goplyYamlConfig
	// strict mode set to true by default
	gyc.Lexer.StrictMode = true
	err := yaml.UnmarshalStrict([]byte(yamlConfig), &gyc)
	if err != nil {
		return nil, err
	}

	if gyc.ConfVersion != goplyConfVersion {
		return nil, fmt.Errorf("expected yaml conf version %s, got %s", goplyConfVersion, gyc.ConfVersion)
	}

	lex := &Lexer{ls: LexerState{SourceLength: len(source) - 1, Source: source},
		lexRules: make(map[string]*regexp.Regexp), lexerErrorFunc: defaultLexerError, strictMode: gyc.Lexer.StrictMode}

	for _, rule := range gyc.Lexer.Rules {
		if rule.Type != "" && rule.Regex != "" {
			lex.AddRule(rule.Type, rule.Regex)
		} else {
			return nil, fmt.Errorf("malformed config file, \"type\" or \"regex\" fields missing from a rule")
		}
	}
	for _, rule := range gyc.Lexer.Ignore {
		if rule != "" {
			lex.Ignore(rule)
		} else {
			return nil, fmt.Errorf("malformed config file, empty rule in \"ignore\"")
		}
	}
	return lex, nil
}
