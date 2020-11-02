package regex

import (
	"fmt"
	"go-custom-compiler/regex/constant/regexconstalfabetica"
	"go-custom-compiler/regex/constant/regexconstdefault"
	"go-custom-compiler/regex/constant/regexconstentera"
	"go-custom-compiler/regex/constant/regexconstlogica"
	"go-custom-compiler/regex/constant/regexconstreal"
	"go-custom-compiler/regex/functionproto/regexfunctionprotoalfabetico"
	"go-custom-compiler/regex/functionproto/regexfunctionprotodefault"
	"go-custom-compiler/regex/functionproto/regexfunctionprotoentero"
	"go-custom-compiler/regex/functionproto/regexfunctionprotologico"
	"go-custom-compiler/regex/functionproto/regexfunctionprotoreal"
	"go-custom-compiler/regex/procedureproto/regexprocedureprotoalfabetico"
	"go-custom-compiler/regex/procedureproto/regexprocedureprotodefault"
	"go-custom-compiler/regex/procedureproto/regexprocedureprotoentero"
	"go-custom-compiler/regex/procedureproto/regexprocedureprotologico"
	"go-custom-compiler/regex/procedureproto/regexprocedureprotoreal"
	"go-custom-compiler/regex/reserved/instruction/regexcustomfunction"
	_regeximprime "go-custom-compiler/regex/reserved/instruction/regeximprime"
	"go-custom-compiler/regex/reserved/instruction/regexio"
	"go-custom-compiler/regex/reserved/instruction/regexlee"
	"go-custom-compiler/regex/reserved/instruction/regexregresa"
	"go-custom-compiler/regex/reserved/instruction/regexsystem"
	"go-custom-compiler/regex/reserved/regexasignacion"
	"go-custom-compiler/regex/reserved/regexcondition/regexconditioncuando"
	"go-custom-compiler/regex/reserved/regexcondition/regexconditionsi"
	"go-custom-compiler/regex/reserved/regexcondition/regexconditionswitch"
	"go-custom-compiler/regex/reserved/regexconstante"
	"go-custom-compiler/regex/reserved/regexcustom"
	"go-custom-compiler/regex/reserved/regexfin/regexfin"
	"go-custom-compiler/regex/reserved/regexfin/regexfinfunction"
	"go-custom-compiler/regex/reserved/regexfin/regexfinprocedure"
	"go-custom-compiler/regex/reserved/regexfin/regexfinprograma"
	"go-custom-compiler/regex/reserved/regexfuncionproto"
	"go-custom-compiler/regex/reserved/regexfunction"
	"go-custom-compiler/regex/reserved/regexinicio"
	"go-custom-compiler/regex/reserved/regexloop/regexloopdesde"
	"go-custom-compiler/regex/reserved/regexloop/regexloophastaque"
	"go-custom-compiler/regex/reserved/regexloop/regexlooprepetir"
	_regexaritmetico "go-custom-compiler/regex/reserved/regexoperador/regexaritmetico"
	_regexlogico "go-custom-compiler/regex/reserved/regexoperador/regexlogico"
	_regexrelacional "go-custom-compiler/regex/reserved/regexoperador/regexrelacional"
	"go-custom-compiler/regex/reserved/regexprocedure"
	"go-custom-compiler/regex/reserved/regexprocedureproto"
	"go-custom-compiler/regex/reserved/regexprograma"
	"go-custom-compiler/regex/reserved/regexvariable"
	_regexvar "go-custom-compiler/regex/variable/regexvar"
	"go-custom-compiler/regex/variable/regexvaralfabetico"
	"go-custom-compiler/regex/variable/regexvardefault"
	"go-custom-compiler/regex/variable/regexvarentero"
	"go-custom-compiler/regex/variable/regexvarlogico"
	"go-custom-compiler/regex/variable/regexvarreal"
	"log"
	"regexp"
)

//CustomRegex ...
type CustomRegex struct {
	//Constant
	RegexConstante           *regexconstante.RegexConstante
	RegexConstanteDefault    *regexconstdefault.RegexConstDefault
	RegexConstanteEntera     *regexconstentera.RegexConstEntera
	RegexConstanteReal       *regexconstreal.RegexConstReal
	RegexConstanteLogica     *regexconstlogica.RegexConstLogica
	RegexConstanteAlfabetica *regexconstalfabetica.RegexConstAlfabetica
	//Variable
	RegexVariable           *regexvariable.RegexVariable
	RegexVariableAlfabetico *regexvaralfabetico.RegexVarAlfabetico
	RegexVariableEntero     *regexvarentero.RegexVarEntero
	RegexVariableLogico     *regexvarlogico.RegexVarLogico
	RegexVariableReal       *regexvarreal.RegexVarReal
	RegexVariableDefault    *regexvardefault.RegexVarDefault
	RegexVar                *_regexvar.RegexVar
	//Funcion Proto
	RegexFuncionProto            *regexfuncionproto.RegexFuncionProto
	RegexFunctionProtoDefault    *regexfunctionprotodefault.RegexFuncProtoDefault
	RegexFunctionProtoAlfabetico *regexfunctionprotoalfabetico.RegexFuncProtoAlfabetico
	RegexFunctionProtoEntero     *regexfunctionprotoentero.RegexFuncProtoEntero
	RegexFunctionProtoReal       *regexfunctionprotoreal.RegexFuncProtoReal
	RegexFunctionProtoLogico     *regexfunctionprotologico.RegexFuncProtoLogico
	//Procimiento Proto
	RegexProcedureProto           *regexprocedureproto.RegexProcedureProto
	RegexProcedureProtoDefault    *regexprocedureprotodefault.RegexProcedureProtoDefault
	RegexProcedureProtoAlfabetico *regexprocedureprotoalfabetico.RegexProcedureProtoAlfabetico
	RegexProcedureProtoEntero     *regexprocedureprotoentero.RegexProcedureProtoEntero
	RegexProcedureProtoReal       *regexprocedureprotoreal.RegexProcedureProtoReal
	RegexProcedureProtoLogico     *regexprocedureprotologico.RegexProcedureProtoLogico
	//Procedimiento
	RegexProcedure *regexprocedure.RegexProcedure
	//Funcion
	RegexFunction *regexfunction.RegexFunction
	//Inicio
	RegexInicio *regexinicio.RegexInicio
	//Fin
	RegexFinProcedure *regexfinprocedure.RegexFinProcedure
	RegexFinFunction  *regexfinfunction.RegexFinFunction
	RegexFin          *regexfin.RegexFin
	//Loop
	RegexLoopRepetir  *regexlooprepetir.RegexLoopRepetir
	RegexLoopHastaQue *regexloophastaque.RegexLoopHastaQue
	RegexLoopDesde    *regexloopdesde.RegexLoopDesde
	//Custom
	RegexCustom         *regexcustom.RegexCustom
	RegexCustomFunction *regexcustomfunction.RegexCustomFunction
	//Instructions
	RegexIO         *regexio.RegexIO
	RegexRegresa    *regexregresa.RegexRegresa
	RegexImprime    *_regeximprime.RegexImprime
	RegexLee        *regexlee.RegexLee
	RegexSystem     *regexsystem.RegexSystem
	RegexAsignacion *regexasignacion.RegexAsignacion
	//Conditions
	RegexConditionCuando *regexconditioncuando.RegexConditionCuando
	RegexConditionSi     *regexconditionsi.RegexConditionSi
	RegexConditionSwitch *regexconditionswitch.RegexConditionSwitch
	//Operators
	RegexOperatorRelacional *_regexrelacional.RegexOpRelacional
	RegexOperatorAritmetico *_regexaritmetico.RegexOpAritmetico
	RegexOperatorLogico     *_regexlogico.RegexOpLogico
	//Programa
	RegexPrograma    *regexprograma.RegexPrograma
	RegexFinPrograma *regexfinprograma.RegexFinPrograma

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
	constdefaultBuilder, _ := regexconstdefault.NewRegexConstDefault()
	constenteraBuilder, _ := regexconstentera.NewRegexConstEntera()
	constrealBuilder, _ := regexconstreal.NewRegexConstReal()
	constlogicaBuilder, _ := regexconstlogica.NewRegexConstLogica()
	constalfabeticaBuilder, _ := regexconstalfabetica.NewRegexConstAlfabetica()

	//Variable
	variableBuilder, _ := regexvariable.NewRegexVariable(EL, LL, GL)
	varalfabeticoBuilder, _ := regexvaralfabetico.NewRegexVariableAlfabetico()
	varenteroBuilder, _ := regexvarentero.NewRegexVariableEntero()
	varlogicoBuilder, _ := regexvarlogico.NewRegexVariableLogico()
	varrealBuilder, _ := regexvarreal.NewRegexVariableReal()
	vardefaultBuilder, _ := regexvardefault.NewRegexVariableDefault()
	varBuilder, _ := _regexvar.NewRegexVar()

	//FunctionProto
	funcionProtoBuilder, _ := regexfuncionproto.NewRegexFuncionProto(EL, LL, GL)
	funcProtoDefault, _ := regexfunctionprotodefault.NewRegexFuncProtoDefault()
	funcProtoAlfabetico, _ := regexfunctionprotoalfabetico.NewRegexFuncProtoAlfabetico()
	funcProtoEntero, _ := regexfunctionprotoentero.NewRegexFuncProtoEntero()
	funcProtoReal, _ := regexfunctionprotoreal.NewRegexFuncProtoReal()
	funcProtoLogico, _ := regexfunctionprotologico.NewRegexFuncProtoLogico()

	//ProcedureProto
	procedureProtoBuilder, _ := regexprocedureproto.NewRegexProcedureProto(EL, LL, GL)
	procedureProtoDefault, _ := regexprocedureprotodefault.NewRegexProcedureProtoDefault()
	procedureProtoAlfabetico, _ := regexprocedureprotoalfabetico.NewRegexProcedureProtoAlfabetico()
	procedureProtoEntero, _ := regexprocedureprotoentero.NewRegexProcedureProtoEntero()
	procedureProtoReal, _ := regexprocedureprotoreal.NewRegexProcedureProtoReal()
	procedureProtoLogico, _ := regexprocedureprotologico.NewRegexProcedureProtoLogico()

	//Procedure
	procedureBuilder, _ := regexprocedure.NewRegexProcedure(EL, LL, GL)

	//Function
	functionBuilder, _ := regexfunction.NewRegexFunction(EL, LL, GL)

	//Inicio
	inicioBuilder, _ := regexinicio.NewRegexInicio(EL, LL, GL)

	//Fin
	finProcedureBuilder, _ := regexfinprocedure.NewRegexFinProcedure(EL, LL, GL)
	finFunctionBuilder, _ := regexfinfunction.NewRegexFinFunction(EL, LL, GL)
	finBuilder, _ := regexfin.NewRegexFin(EL, LL, GL)

	//Loop
	repetirLoopBuilder, _ := regexlooprepetir.NewRegexLoopRepetir(EL, LL, GL)
	hastaqueLoopBuilder, _ := regexloophastaque.NewRegexLoopHastaQue(EL, LL, GL)
	desdeLoopBuilder, _ := regexloopdesde.NewRegexLoopDesde(EL, LL, GL)

	//Custom
	customBuilder, _ := regexcustom.NewRegexCustom(EL, LL, GL)
	customFunctionsBuilder, _ := regexcustomfunction.NewRegexCustomFunction(EL, LL, GL)
	//Instructions
	ioBuilder, _ := regexio.NewRegexIO(EL, LL, GL)
	regresaBuilder, _ := regexregresa.NewRegexRegresa(EL, LL, GL)
	conditionCuandoBuilder, _ := regexconditioncuando.NewRegexConditionCuando(EL, LL, GL)
	conditionSiBuilder, _ := regexconditionsi.NewRegexConditionSi(EL, LL, GL)
	conditionSwitchBuilder, _ := regexconditionswitch.NewRegexConditionSwitch(EL, LL, GL)
	imprimeBuilder, _ := _regeximprime.NewRegexImprime()
	leeBuilder, _ := regexlee.NewRegexLee()
	//Operators
	relacionalBuilder, _ := _regexrelacional.NewRegexOpRelacional()
	aritmeticoBuilder, _ := _regexaritmetico.NewRegexOpAritmetico()
	logicoBuilder, _ := _regexlogico.NewRegexOpLogico()
	systemBuilder, _ := regexsystem.NewRegexSystem(EL, LL, GL)
	asignacionBuilder, _ := regexasignacion.NewRegexAsignacion(EL, LL, GL)
	//Programa
	finprogramaBuilder, _ := regexfinprograma.NewRegexFinPrograma(EL, LL, GL)
	programaBuilder, _ := regexprograma.NewRegexPrograma(EL, LL, GL)

	return &CustomRegex{
		//Reserved
		RegexConstante: constanteBuilder,
		RegexVariable:  variableBuilder,
		//Constante
		RegexConstanteDefault:    constdefaultBuilder,
		RegexConstanteAlfabetica: constalfabeticaBuilder,
		RegexConstanteEntera:     constenteraBuilder,
		RegexConstanteLogica:     constlogicaBuilder,
		RegexConstanteReal:       constrealBuilder,
		//Variable
		RegexVariableAlfabetico: varalfabeticoBuilder,
		RegexVariableEntero:     varenteroBuilder,
		RegexVariableLogico:     varlogicoBuilder,
		RegexVariableReal:       varrealBuilder,
		RegexVariableDefault:    vardefaultBuilder,
		RegexVar:                varBuilder,
		//Function Proto
		RegexFuncionProto:            funcionProtoBuilder,
		RegexFunctionProtoAlfabetico: funcProtoAlfabetico,
		RegexFunctionProtoDefault:    funcProtoDefault,
		RegexFunctionProtoEntero:     funcProtoEntero,
		RegexFunctionProtoLogico:     funcProtoLogico,
		RegexFunctionProtoReal:       funcProtoReal,
		//Procedure Proto
		RegexProcedureProto:           procedureProtoBuilder,
		RegexProcedureProtoAlfabetico: procedureProtoAlfabetico,
		RegexProcedureProtoDefault:    procedureProtoDefault,
		RegexProcedureProtoEntero:     procedureProtoEntero,
		RegexProcedureProtoLogico:     procedureProtoLogico,
		RegexProcedureProtoReal:       procedureProtoReal,
		//Procedure
		RegexProcedure: procedureBuilder,
		//Function
		RegexFunction: functionBuilder,
		//Inicio
		RegexInicio: inicioBuilder,
		//Fin
		RegexFinFunction:  finFunctionBuilder,
		RegexFinProcedure: finProcedureBuilder,
		RegexFin:          finBuilder,
		//Loop
		RegexLoopRepetir:  repetirLoopBuilder,
		RegexLoopHastaQue: hastaqueLoopBuilder,
		RegexLoopDesde:    desdeLoopBuilder,
		//Custom
		RegexCustom:         customBuilder,
		RegexCustomFunction: customFunctionsBuilder,
		//Instructions
		RegexIO:         ioBuilder,
		RegexRegresa:    regresaBuilder,
		RegexImprime:    imprimeBuilder,
		RegexLee:        leeBuilder,
		RegexSystem:     systemBuilder,
		RegexAsignacion: asignacionBuilder,
		//Conditions
		RegexConditionCuando: conditionCuandoBuilder,
		RegexConditionSi:     conditionSiBuilder,
		RegexConditionSwitch: conditionSwitchBuilder,
		//Operators
		RegexOperatorRelacional: relacionalBuilder,
		RegexOperatorAritmetico: aritmeticoBuilder,
		RegexOperatorLogico:     logicoBuilder,
		//Programa
		RegexFinPrograma: finprogramaBuilder,
		RegexPrograma:    programaBuilder,

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
