package regex

import (
	"fmt"
	"go-custom-compiler/regex/regexconstante"
	"go-custom-compiler/regex/regexfloat"
	"go-custom-compiler/regex/regexint"
	"go-custom-compiler/regex/regexvariable"
	"log"
	"regexp"
)

//CustomRegex ...
type CustomRegex struct {
	RegexConstante *regexconstante.RegexConstante
	RegexVariable  *regexvariable.RegexVariable
	RegexFloat     *regexfloat.RegexFloat
	RegexInt       *regexint.RegexInt

	EL *log.Logger
	LL *log.Logger
}

//NewRegex ...
func NewRegex(EL *log.Logger, LL *log.Logger) (*CustomRegex, error) {
	if EL == nil || LL == nil {
		return nil, fmt.Errorf("EL or LL loggers came empty")
	}
	constanteBuilder, _ := regexconstante.NewRegexConstante(EL, LL)
	variableBuilder, _ := regexvariable.NewRegexVariable()
	floatBuilder, _ := regexfloat.NewRegexFloat()
	intBuilder, _ := regexint.NewRegexInt()

	return &CustomRegex{
		RegexConstante: constanteBuilder,
		RegexVariable:  variableBuilder,
		RegexFloat:     floatBuilder,
		RegexInt:       intBuilder,
		EL:             EL,
		LL:             LL,
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
