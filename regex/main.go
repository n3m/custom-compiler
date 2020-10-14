package regex

import (
	"fmt"
	"go-custom-compiler/regex/constant/regexconstfloat"
	"go-custom-compiler/regex/constant/regexconstint"
	"go-custom-compiler/regex/functionproto/regexfunctionprotoalfabetico"
	"go-custom-compiler/regex/functionproto/regexfunctionprotodefault"
	"go-custom-compiler/regex/functionproto/regexfunctionprotoentero"
	"go-custom-compiler/regex/functionproto/regexfunctionprotoflotante"
	"go-custom-compiler/regex/functionproto/regexfunctionprotologico"
	"go-custom-compiler/regex/functionproto/regexfunctionprotoreal"
	"go-custom-compiler/regex/reserved/regexconstante"
	"go-custom-compiler/regex/reserved/regexfuncionproto"
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
	//Constante
	RegexConstante      *regexconstante.RegexConstante
	RegexConstanteFloat *regexconstfloat.RegexFloat
	RegexConstanteInt   *regexconstint.RegexInt
	//Variable
	RegexVariable           *regexvariable.RegexVariable
	RegexVariableAlfabetico *regexvaralfabetico.RegexVarAlfabetico
	RegexVariableEntero     *regexvarentero.RegexVarEntero
	RegexVariableFlotante   *regexvarflotante.RegexVarFlotante
	RegexVariableLogico     *regexvarlogico.RegexVarLogico
	RegexVariableReal       *regexvarreal.RegexVarReal
	RegexVariableDefault    *regexvardefault.RegexVarDefault
	//Function Proto
	RegexFuncionProto            *regexfuncionproto.RegexFuncionProto
	RegexFunctionProtoDefault    *regexfunctionprotodefault.RegexFuncProtoDefault
	RegexFunctionProtoAlfabetico *regexfunctionprotoalfabetico.RegexFuncProtoAlfabetico
	RegexFunctionProtoEntero     *regexfunctionprotoentero.RegexFuncProtoEntero
	RegexFunctionProtoReal       *regexfunctionprotoreal.RegexFuncProtoReal
	RegexFunctionProtoLogico     *regexfunctionprotologico.RegexFuncProtoLogico
	RegexFunctionProtoFlotante   *regexfunctionprotoflotante.RegexFuncProtoFlotante

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

	//Constante
	constanteBuilder, _ := regexconstante.NewRegexConstante(EL, LL, GL)
	constfloatBuilder, _ := regexconstfloat.NewRegexFloat()
	constintBuilder, _ := regexconstint.NewRegexInt()

	//Variable
	variableBuilder, _ := regexvariable.NewRegexVariable(EL, LL, GL)
	varalfabeticoBuilder, _ := regexvaralfabetico.NewRegexVariableAlfabetico()
	varenteroBuilder, _ := regexvarentero.NewRegexVariableEntero()
	varflotanteBuilder, _ := regexvarflotante.NewRegexVariableFlotante()
	varlogicoBuilder, _ := regexvarlogico.NewRegexVariableLogico()
	varrealBuilder, _ := regexvarreal.NewRegexVariableReal()
	vardefaultBuilder, _ := regexvardefault.NewRegexVariableDefault()

	//FunctionProto
	funcionProtoBuilder, _ := regexfuncionproto.NewRegexFuncionProto(EL, LL, GL)
	funcProtoDefault, _ := regexfunctionprotodefault.NewRegexFuncProtoDefault()
	funcProtoAlfabetico, _ := regexfunctionprotoalfabetico.NewRegexFuncProtoAlfabetico()
	funcProtoEntero, _ := regexfunctionprotoentero.NewRegexFuncProtoEntero()
	funcProtoReal, _ := regexfunctionprotoreal.NewRegexFuncProtoReal()
	funcProtoFlotante, _ := regexfunctionprotoflotante.NewRegexFuncProtoFlotante()
	funcProtoLogico, _ := regexfunctionprotologico.NewRegexFuncProtoLogico()

	return &CustomRegex{
		//Reserved
		RegexConstante: constanteBuilder,
		RegexVariable:  variableBuilder,
		//Variants
		RegexConstanteFloat:     constfloatBuilder,
		RegexConstanteInt:       constintBuilder,
		RegexVariableAlfabetico: varalfabeticoBuilder,
		RegexVariableEntero:     varenteroBuilder,
		RegexVariableFlotante:   varflotanteBuilder,
		RegexVariableLogico:     varlogicoBuilder,
		RegexVariableReal:       varrealBuilder,
		RegexVariableDefault:    vardefaultBuilder,
		//Proto
		RegexFuncionProto:            funcionProtoBuilder,
		RegexFunctionProtoAlfabetico: funcProtoAlfabetico,
		RegexFunctionProtoDefault:    funcProtoDefault,
		RegexFunctionProtoEntero:     funcProtoEntero,
		RegexFunctionProtoFlotante:   funcProtoFlotante,
		RegexFunctionProtoLogico:     funcProtoLogico,
		RegexFunctionProtoReal:       funcProtoReal,

		EL: EL,
		LL: LL,
		GL: GL,
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
