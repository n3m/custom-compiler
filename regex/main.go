package regex

import (
	"fmt"
	"go-custom-compiler/regex/regexconstante"
	"go-custom-compiler/regex/regexfloat"
	"go-custom-compiler/regex/regexint"
	"go-custom-compiler/regex/regexvariable"
	"regexp"
)

//CustomRegex ...
type CustomRegex struct {
	RegexConstante *regexconstante.RegexConstante
	RegexVariable  *regexvariable.RegexVariable
	RegexFloat     *regexfloat.RegexFloat
	RegexInt       *regexint.RegexInt
}

//NewRegex ...
func NewRegex() (*CustomRegex, error) {
	constanteBuilder, _ := regexconstante.NewRegexConstante()
	variableBuilder, _ := regexvariable.NewRegexVariable()
	floatBuilder, _ := regexfloat.NewRegexFloat()
	intBuilder, _ := regexint.NewRegexInt()

	return &CustomRegex{
		RegexConstante: constanteBuilder,
		RegexVariable:  variableBuilder,
		RegexFloat:     floatBuilder,
		RegexInt:       intBuilder,
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
