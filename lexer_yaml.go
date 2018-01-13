package goply

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var goplyConfVersion = "0.0.1"

// Struct used to read the yaml config into
type goplyYamlConfig struct {
	ConfVersion string `yaml:"version,omitempty"`
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
// returns an error if could not read the yaml properly
func NewLexerFromYamlConfig(yamlConfig []byte) (*Lexer, error) {
	var gyc goplyYamlConfig
	// strict mode set to true by default
	gyc.Lexer.StrictMode = true
	//gyc.ConfVersion = "invalid"
	err := yaml.UnmarshalStrict([]byte(yamlConfig), &gyc)
	if err != nil {
		return nil, err
	}

	if gyc.ConfVersion != goplyConfVersion {
		if gyc.ConfVersion == "" {
			gyc.ConfVersion = "none"
		}

		//fmt.Printf("%v", gyc.ConfVersion)

		return nil, fmt.Errorf("expected yaml conf version %s, got %s", goplyConfVersion, gyc.ConfVersion)
	}

	lex := NewLexer(gyc.Lexer.StrictMode)

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
