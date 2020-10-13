package regex

import (
	"fmt"
	"go-custom-compiler/regex/constant/regexfloat"
	"go-custom-compiler/regex/constant/regexint"
	"go-custom-compiler/regex/reserved/regexconstante"
	"go-custom-compiler/regex/reserved/regexvariable"
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
	GL *log.Logger
}

//NewRegex ...
func NewRegex(EL *log.Logger, LL *log.Logger, GL *log.Logger) (*CustomRegex, error) {
	var moduleName string = "[regex][NewRegex()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
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
		GL:             GL,
	}, nil
}

//StartsWith ...
func (r CustomRegex) StartsWith(prefix, strToTest string) (bool, error) {
	var moduleName string = "[regex][StartsWith()]"

	compiled, err := regexp.Compile("^" + prefix)
	if err != nil {
		return false, fmt.Errorf("[ERROR]%+v %+v", moduleName, err)
	}

	return compiled.MatchString(strToTest), nil
}

//EndsWith ...
func (r CustomRegex) EndsWith(suffix, strToTest string) (bool, error) {
	var moduleName string = "[regex][EndsWith()]"

	compiled, err := regexp.Compile(suffix + "$")
	if err != nil {
		return false, fmt.Errorf("[ERROR]%+v %+v", moduleName, err)
	}

	return compiled.MatchString(strToTest), nil
}
