package regex

import (
	"fmt"
	"go-custom-compiler/regex/constante"
	"go-custom-compiler/regex/variable"
	"regexp"
)

//CustomRegex ...
type CustomRegex struct {
	RegexConstante *constante.RegexConstante
	RegexVariable  *variable.RegexVariable
}

//NewRegex ...
func NewRegex() (*CustomRegex, error) {
	constanteBuilder, _ := constante.NewRegexConstante()
	variableBuilder, _ := variable.NewRegexVariable()

	return &CustomRegex{
		RegexConstante: constanteBuilder,
		RegexVariable:  variableBuilder,
	}, nil
}

//StartsWith ...
func (r CustomRegex) StartsWith(prefix, strToTest string) (bool, error) {
	compiled, err := regexp.Compile("^" + prefix)
	if err != nil {
		return false, fmt.Errorf("[ERROR] %+v", err)
	}

	return compiled.MatchString(strToTest), nil
}

//EndsWith ...
func (r CustomRegex) EndsWith(suffix, strToTest string) (bool, error) {
	compiled, err := regexp.Compile(suffix + "$")
	if err != nil {
		return false, fmt.Errorf("[ERROR] %+v", err)
	}

	return compiled.MatchString(strToTest), nil
}
