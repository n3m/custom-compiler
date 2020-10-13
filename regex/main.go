package regex

import (
	"fmt"
	"go-custom-compiler/regex/constant/regexconstfloat"
	"go-custom-compiler/regex/constant/regexconstint"
	"go-custom-compiler/regex/reserved/regexconstante"
	"go-custom-compiler/regex/reserved/regexvariable"
	"go-custom-compiler/regex/variable/regexvaralfabetico"
	"go-custom-compiler/regex/variable/regexvardefault"
	"go-custom-compiler/regex/variable/regexvarentero"
	"go-custom-compiler/regex/variable/regexvarflotante"
	"go-custom-compiler/regex/variable/regexvarlogico"
	"go-custom-compiler/regex/variable/regexvarreal"
	"log"
	"regexp"
)

//CustomRegex ...
type CustomRegex struct {
	RegexConstante          *regexconstante.RegexConstante
	RegexVariable           *regexvariable.RegexVariable
	RegexConstanteFloat     *regexconstfloat.RegexFloat
	RegexConstanteInt       *regexconstint.RegexInt
	RegexVariableAlfabetico *regexvaralfabetico.RegexVarAlfabetico
	RegexVariableEntero     *regexvarentero.RegexVarEntero
	RegexVariableFlotante   *regexvarflotante.RegexVarFlotante
	RegexVariableLogico     *regexvarlogico.RegexVarLogico
	RegexVariableReal       *regexvarreal.RegexVarReal
	RegexVariableDefault    *regexvardefault.RegexVarDefault

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
	constfloatBuilder, _ := regexconstfloat.NewRegexFloat()
	constintBuilder, _ := regexconstint.NewRegexInt()
	varalfabeticoBuilder, _ := regexvaralfabetico.NewRegexVariableAlfabetico()
	varenteroBuilder, _ := regexvarentero.NewRegexVariableEntero()
	varflotanteBuilder, _ := regexvarflotante.NewRegexVariableFlotante()
	varlogicoBuilder, _ := regexvarlogico.NewRegexVariableLogico()
	varrealBuilder, _ := regexvarreal.NewRegexVariableReal()
	vardefaultBuilder, _ := regexvardefault.NewRegexVariableDefault()

	return &CustomRegex{
		RegexConstante:          constanteBuilder,
		RegexVariable:           variableBuilder,
		RegexConstanteFloat:     constfloatBuilder,
		RegexConstanteInt:       constintBuilder,
		RegexVariableAlfabetico: varalfabeticoBuilder,
		RegexVariableEntero:     varenteroBuilder,
		RegexVariableFlotante:   varflotanteBuilder,
		RegexVariableLogico:     varlogicoBuilder,
		RegexVariableReal:       varrealBuilder,
		RegexVariableDefault:    vardefaultBuilder,
		EL:                      EL,
		LL:                      LL,
		GL:                      GL,
	}, nil
}

//StartsWith ...
func (r CustomRegex) StartsWith(prefix, strToTest string) (bool, error) {
	var moduleName string = "[regex][StartsWith()]"

	compiled, err := regexp.Compile("^" + prefix)
	if err != nil {
		return false, fmt.Errorf("[ERROR]%+v %+v", moduleName, err.Error())
	}

	return compiled.MatchString(strToTest), nil
}

//EndsWith ...
func (r CustomRegex) EndsWith(suffix, strToTest string) (bool, error) {
	var moduleName string = "[regex][EndsWith()]"

	compiled, err := regexp.Compile(suffix + "$")
	if err != nil {
		return false, fmt.Errorf("[ERROR]%+v %+v", moduleName, err.Error())
	}

	return compiled.MatchString(strToTest), nil
}
