package lexyc

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"

	"go-custom-compiler/helpers"
	"go-custom-compiler/models"
	"go-custom-compiler/regex"

	"github.com/DrN3MESiS/pprnt"
)

//LexicalAnalyzer ...
type LexicalAnalyzer struct {
	File *bufio.Scanner     //File
	R    *regex.CustomRegex //Regex Handler
	EL   *log.Logger        //Error Logger
	LL   *log.Logger        //Lex Logger
	GL   *log.Logger        //General Logger

	//TEST
	CurrentBlockType models.BlockType
	ParentBlockType  models.BlockType
	BlockQueue       []models.BlockType
	OpQueue          []models.TokenComp
	ConstantStorage  []models.Token
	VariableStorage  []models.Token
}

//NewLexicalAnalyzer ...
func NewLexicalAnalyzer(file *bufio.Scanner, ErrorLogger, LexLogger, GeneralLogger *log.Logger) (*LexicalAnalyzer, error) {
	var moduleName string = "[Lexyc][NewLexicalAnalyzer()]"

	if file == nil {
		GeneralLogger.Printf("[ERR]%+v file is not present", moduleName)
		return nil, fmt.Errorf("[ERR]%+v file is not present", moduleName)
	}
	if ErrorLogger == nil || LexLogger == nil || GeneralLogger == nil {
		GeneralLogger.Printf("[ERR]%+v Loggers are not present", moduleName)
		return nil, fmt.Errorf("[ERR]%+v Loggers are not present", moduleName)
	}
	R, err := regex.NewRegex(ErrorLogger, LexLogger, GeneralLogger)
	if err != nil {
		GeneralLogger.Printf("[ERR]%+v %+v", moduleName, err.Error())
		return nil, fmt.Errorf("[ERR]%+v %+v", moduleName, err.Error())
	}

	LexLogger.Println("----------------------------------------------")
	LexLogger.Println(helpers.IndentString(helpers.LEXINDENT, []string{"Lexema", "Token"}))
	LexLogger.Println("----------------------------------------------")
	ErrorLogger.Printf("=============================================================")
	ErrorLogger.Printf("# Linea | # Columna | Error | Descripcion | Linea del Error")
	ErrorLogger.Printf("=============================================================")

	return &LexicalAnalyzer{
		File: file,
		R:    R,
		EL:   ErrorLogger,
		LL:   LexLogger,
		GL:   GeneralLogger,

		ParentBlockType:  models.NULLBLOCK,
		BlockQueue:       []models.BlockType{},
		CurrentBlockType: models.NULLBLOCK,
		OpQueue:          []models.TokenComp{},
		ConstantStorage:  []models.Token{},
		VariableStorage:  []models.Token{},
	}, nil
}

//Analyze ...
func (l *LexicalAnalyzer) Analyze(debug bool) error {
	funcName := "[Analyze()]"
	var lineIndex int64 = 1
	for l.File.Scan() {
		currentLine := l.File.Text()

		if len(currentLine) == 0 {
			l.GL.Printf("%+v Skipped [Line: %+v]; Reason: Empty", funcName, lineIndex)
			lineIndex++

			continue
		}
		var LastBlockState models.BlockType
		LastBlockState = l.CurrentBlockType
		/* Type Validation */
		isComment, err := l.R.StartsWith("//", currentLine)
		if err != nil {
			l.GL.Printf("%+v[APP_ERR] %+v", funcName, err.Error())
			return fmt.Errorf("%+v[APP_ERR] %+v", funcName, err.Error())
		}
		if isComment {
			l.GL.Printf("%+vSkipping Comment at line %+v", funcName, lineIndex)
			if debug {
				log.Printf("Skipping Comment at line %+v", lineIndex)
			}
			lineIndex++

			continue
		}

		currentLine = strings.TrimSpace(currentLine)

		// log.Printf("BLOCK [Line:%+v]['%+v'] > %+v\n", lineIndex, currentLine, l.BlockQueue)
		log.Printf("BLOCK [Line:%+v] > %+v\n", lineIndex, l.BlockQueue)

		/* StartsWith */

		//Contante
		if l.R.RegexConstante.StartsWithConstante(currentLine, lineIndex) {
			l.CurrentBlockType = models.CONSTANTBLOCK
			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"constantes", helpers.PALABRARESERVADA}))
		}

		//Variable
		if l.R.RegexVariable.StartsWithVariable(currentLine, lineIndex) {
			l.CurrentBlockType = models.VARIABLEBLOCK
			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"variables", helpers.PALABRARESERVADA}))
		}

		//FunctionProto
		if l.R.RegexFuncionProto.StartsWithFuncionProto(currentLine, lineIndex) && l.ParentBlockType == models.NULLBLOCK {
			l.CurrentBlockType = models.FUNCTIONPROTOBLOCK
			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"funcion", helpers.PALABRARESERVADA}))
		}

		//ProcedureProto
		if l.R.RegexProcedureProto.StartsWithProcedureProto(currentLine, lineIndex) && l.ParentBlockType == models.NULLBLOCK {
			l.CurrentBlockType = models.PROCEDUREPROTOBLOCK
			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"procedimiento", helpers.PALABRARESERVADA}))
		}

		//Procedure
		if l.R.RegexProcedure.StartsWithProcedure(currentLine, lineIndex) {
			l.GL.Println()

			if len(l.BlockQueue) > 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to create new procedure without finalizing the last Function or Procedure", currentLine)
				l.BlockQueue = []models.BlockType{}
			}
			l.BlockQueue = append(l.BlockQueue, models.PROCEDUREBLOCK)
		}

		//Function
		if l.R.RegexFunction.StartsWithFunction(currentLine, lineIndex) {
			l.GL.Println()

			if len(l.BlockQueue) > 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to create new function without finalizing the last Function or Procedure", currentLine)
				l.BlockQueue = []models.BlockType{}
			}

			l.BlockQueue = append(l.BlockQueue, models.FUNCTIONBLOCK)
		}

		//Inicio
		if l.R.RegexInicio.StartsWithInicio(currentLine, lineIndex) {
			if len(l.BlockQueue) == 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to initialize something outside of a Block", currentLine)
			}

			switch l.BlockQueue[len(l.BlockQueue)-1] {
			case models.INITBLOCK:
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to initialize something when already initialized", currentLine)
				break
			case models.PROCEDUREBLOCK, models.FUNCTIONBLOCK, models.CUANDOBLOCK:
				l.GL.Printf("%+v Initialized a %+v [Line: %+v]", funcName, l.BlockQueue[len(l.BlockQueue)-1], lineIndex)
				l.BlockQueue = append(l.BlockQueue, models.INITBLOCK)
				break

			default:
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to initialize something non existent", currentLine)
				break
			}

			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"Inicio", helpers.PALABRARESERVADA}))
		}

		//FinDeFuncion
		if l.R.RegexFinFunction.StartsWithFinDeFuncion(currentLine, lineIndex) {
			if len(l.BlockQueue) == 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a FUNCTIONBLOCK outside of a FUNCTIONBLOCK", currentLine)
			}

			if l.BlockQueue[len(l.BlockQueue)-1] != models.INITBLOCK {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a FUNCTIONBLOCK that wasn't initialized", currentLine)
			}

			newArr, ok := helpers.RemoveFromQueue(l.BlockQueue, models.INITBLOCK)
			if ok {
				l.BlockQueue = newArr
			} else {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a FUNCTIONBLOCK that wasn't initialized", currentLine)
			}

			newArr, ok = helpers.RemoveFromQueue(l.BlockQueue, models.FUNCTIONBLOCK)
			if ok {
				l.BlockQueue = newArr
			} else {
				if helpers.QueueContainsBlock(l.BlockQueue, models.PROCEDUREBLOCK) {
					l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a FUNCTIONBLOCK:Inicio with a PROCEDUREBLOCK as parent", currentLine)
				} else {
					l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a FUNCTIONBLOCK outside of a FUNCTIONBLOCK", currentLine)
				}

			}
			l.GL.Println()
			l.LL.Print(helpers.IndentStringInLines(helpers.LEXINDENT, 2, []string{
				"Fin", helpers.PALABRARESERVADA,
				"de", helpers.PALABRARESERVADA,
				"funcion", helpers.PALABRARESERVADA,
				";", helpers.DELIMITADOR,
			}))
		}

		//FinDeProcedimiento
		if l.R.RegexFinProcedure.StartsWithFinDeProcedimiento(currentLine, lineIndex) {
			if len(l.BlockQueue) == 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a PROCEDUREBLOCK outside of a PROCEDUREBLOCK", currentLine)
			}

			newArr, ok := helpers.RemoveFromQueue(l.BlockQueue, models.INITBLOCK)
			if ok {
				l.BlockQueue = newArr
			} else {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a PROCEDUREBLOCK that wasn't initialized", currentLine)
			}

			newArr, ok = helpers.RemoveFromQueue(l.BlockQueue, models.PROCEDUREBLOCK)
			if ok {
				l.BlockQueue = newArr
			} else {
				if helpers.QueueContainsBlock(l.BlockQueue, models.FUNCTIONBLOCK) {
					l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a PROCEDUREBLOCK:Inicio with a FUNCTIONBLOCK as parent", currentLine)
				} else {
					l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a PROCEDUREBLOCK outside of a PROCEDUREBLOCK", currentLine)
				}
			}
			l.LL.Print(helpers.IndentStringInLines(helpers.LEXINDENT, 2, []string{
				"Fin", helpers.PALABRARESERVADA,
				"de", helpers.PALABRARESERVADA,
				"procedimiento", helpers.PALABRARESERVADA,
				";", helpers.DELIMITADOR,
			}))
		}

		//Fin
		if l.R.RegexFin.StartsWithFin(currentLine, lineIndex) {
			if !l.R.RegexIO.MatchPC(currentLine, lineIndex) {
				l.LogError(lineIndex, len(currentLine)-1, ";", "Missing ';'", currentLine)
			}

			newArr, ok := helpers.RemoveFromQueue(l.BlockQueue, models.INITBLOCK)
			if ok {
				l.BlockQueue = newArr
			} else {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a SOMETHING that wasn't initialized", currentLine)
			}

			switch l.BlockQueue[len(l.BlockQueue)-1] {
			case models.CUANDOBLOCK:
				newArr, ok = helpers.RemoveFromQueue(l.BlockQueue, models.CUANDOBLOCK)
				if ok {
					l.BlockQueue = newArr
				}
				break
			default:
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a SOMETHING:Inicio that didn't exist", currentLine)
				break
			}
			l.LL.Print(helpers.IndentStringInLines(helpers.LEXINDENT, 2, []string{
				"Fin", helpers.PALABRARESERVADA,
				";", helpers.DELIMITADOR,
			}))
		}

		//Repetir
		if l.R.RegexLoopRepetir.StartsWithRepetir(currentLine, lineIndex) {
			if len(l.BlockQueue) == 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to create a REPEATBLOCK outside of a BLOCK", currentLine)
			}

			l.BlockQueue = append(l.BlockQueue, models.REPEATBLOCK)
			l.GL.Printf("%+v Initialized a REPEATBLOCK [Line: %+v]", funcName, lineIndex)

			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"repetir", helpers.PALABRARESERVADA}))
		}

		//Hasta Que (Repetir)
		if l.R.RegexLoopHastaQue.StartsWithHastaQue(currentLine, lineIndex) {
			if len(l.BlockQueue) == 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to end a REPEATBLOCK outside of a BLOCK", currentLine)
			}

			/* Analyze Params */

			data := strings.Split(currentLine, " ")
			currentLine = ""
			for _, str := range data[2:] {
				currentLine += str + " "
			}

			l.OpQueue = []models.TokenComp{}

			stage1 := regexp.MustCompile(`^(\s*)\((.*)\);(\s*)$`)
			stage1v2 := regexp.MustCompile(`^(\s*)\((.*)\)(\s*)$`)
			if stage1.MatchString(currentLine) || stage1v2.MatchString(currentLine) {
				currentLine = strings.TrimPrefix(currentLine, "(")
				currentLine = strings.TrimSuffix(currentLine, ";")
				currentLine = strings.TrimSuffix(currentLine, ")")

				lineData := strings.Split(currentLine, " ")
				switch len(lineData) {
				case 0:
					l.LogError(lineIndex, "N/A", "N/A", "Instruction 'Hasta que' doesn't have params", currentLine)
					break
				case 2:
					l.LogError(lineIndex, "N/A", "N/A", "Instruction 'Hasta que' only has 2 params", currentLine)
					break
				default:
					for _, dat := range lineData {
						l.AnalyzeForItem(dat, lineIndex)
					}
					break
				}

				//TODO: Create Func to eval OPQueue

			} else {
				l.LogError(lineIndex, "N/A", "N/A", "Instruction 'Hasta que' doesn't have params", currentLine)
			}

			/* End Analyze Params*/

			if l.BlockQueue[len(l.BlockQueue)-1] == models.REPEATBLOCK {
				newArr, ok := helpers.RemoveFromQueue(l.BlockQueue, models.REPEATBLOCK)
				if ok {
					l.BlockQueue = newArr
				} else {
					l.LogErrorGeneral(lineIndex, "N/A", "N/A", "I tried to delete something that was inside the slice that I saw before trying to delete", currentLine)
				}
			} else {
				l.LogError(lineIndex, "N/A", "N/A", fmt.Sprintf("Attempted to end a REPEATBLOCK before finalizing a %+v", l.BlockQueue[len(l.BlockQueue)-1]), currentLine)
			}

			l.LL.Print(helpers.IndentStringInLines(helpers.LEXINDENT, 2, []string{
				"hasta", helpers.PALABRARESERVADA,
				"que", helpers.PALABRARESERVADA,
				"(", helpers.DELIMITADOR,
				")", helpers.DELIMITADOR,
				";", helpers.DELIMITADOR,
			}))
		}

		//ImprimeNL
		if l.R.RegexIO.MatchImprimenl(currentLine, lineIndex) {
			if !l.R.RegexIO.MatchPC(currentLine, lineIndex) {
				l.LogError(lineIndex, len(currentLine)-1, ";", "Missing ';'", currentLine)
			}
			currentLine = strings.TrimSuffix(currentLine, ";")
			currentLine = strings.TrimSuffix(currentLine, ")")

			data := strings.Split(currentLine, "(")
			currentLine = ""
			for _, str := range data[1:] {
				currentLine += str + " "
			}

			params := strings.Split(currentLine, ",")

			l.OpQueue = []models.TokenComp{}
			for _, str := range params {
				l.AnalyzeForItem(str, lineIndex)
			}

			if !l.ExpectNoNone() {
				l.LogError(lineIndex, "N/A", "N/A", "One of the parameters introduced is not valid", currentLine)
			}

			l.GL.Printf("%+v Found 'Imprimenl' instruction [Line: %+v]", funcName, lineIndex)

			l.LL.Print(helpers.IndentStringInLines(helpers.LEXINDENT, 2, []string{
				data[0], helpers.PALABRARESERVADA,
				"(", helpers.DELIMITADOR,
				")", helpers.DELIMITADOR,
				";", helpers.DELIMITADOR,
			}))
			//Imprime
		} else if l.R.RegexIO.MatchImprime(currentLine, lineIndex) {
			if !l.R.RegexIO.MatchPC(currentLine, lineIndex) {
				l.LogError(lineIndex, len(currentLine)-1, ";", "Missing ';'", currentLine)
			}
			currentLine = strings.TrimSuffix(currentLine, ";")
			currentLine = strings.TrimSuffix(currentLine, ")")

			data := strings.Split(currentLine, "(")
			currentLine = ""
			for _, str := range data[1:] {
				currentLine += str + " "
			}

			params := strings.Split(currentLine, ",")
			l.OpQueue = []models.TokenComp{}
			for _, str := range params {
				l.AnalyzeForItem(str, lineIndex)
			}

			if !l.ExpectNoNone() {
				l.LogError(lineIndex, "N/A", "N/A", "One of the parameters introduced is not valid", currentLine)
			}
			l.GL.Printf("%+v Found 'Imprime' instruction [Line: %+v]", funcName, lineIndex)

			l.LL.Print(helpers.IndentStringInLines(helpers.LEXINDENT, 2, []string{
				data[0], helpers.PALABRARESERVADA,
				"(", helpers.DELIMITADOR,
				")", helpers.DELIMITADOR,
				";", helpers.DELIMITADOR,
			}))
		}

		//Lee
		if l.R.RegexIO.MatchLee(currentLine, lineIndex) {
			if !l.R.RegexIO.MatchPC(currentLine, lineIndex) {
				l.LogError(lineIndex, len(currentLine)-1, ";", "Missing ';'", currentLine)
			}
			currentLine = strings.TrimSuffix(currentLine, ";")
			currentLine = strings.TrimSuffix(currentLine, ")")

			data := strings.Split(currentLine, "(")
			currentLine = ""
			for _, str := range data[1:] {
				currentLine += str + " "
			}
			params := strings.Split(currentLine, ",")
			l.OpQueue = []models.TokenComp{}
			for _, str := range params {
				l.AnalyzeForItem(str, lineIndex)
			}

			if !l.ExpectIdent(currentLine, lineIndex) {
				l.LogError(lineIndex, "N/A", "N/A", "Expected <Ident> in parameters", currentLine)
			}

			l.GL.Printf("%+v Found 'Lee' instruction [Line: %+v]", funcName, lineIndex)

			l.LL.Print(helpers.IndentStringInLines(helpers.LEXINDENT, 2, []string{
				data[0], helpers.PALABRARESERVADA,
				"(", helpers.DELIMITADOR,
				")", helpers.DELIMITADOR,
				";", helpers.DELIMITADOR,
			}))
		}

		//Cuando
		if l.R.RegexConditionCuando.StartsWithCuando(currentLine, lineIndex) {
			if len(l.BlockQueue) == 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to create a CUANDOBLOCK outside of a BLOCK", currentLine)
			}
			l.BlockQueue = append(l.BlockQueue, models.CUANDOBLOCK)

			//TODO: Get params

			l.GL.Printf("%+v Created a CUANDOBLOCK [Line: %+v]", funcName, lineIndex)

			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"cuando", helpers.PALABRARESERVADA}))
		}

		//Si
		if l.R.RegexConditionSi.StartsWithSi(currentLine, lineIndex) {
			if len(l.BlockQueue) == 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to create a 'Si' condition outside of a BLOCK", currentLine)
			}

			l.R.RegexConditionSi.ValidateCondition(currentLine, lineIndex)

			//TODO: Get Params

			l.GL.Printf("%+v Found 'Si' condition [Line: %+v]", funcName, lineIndex)

			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"si", helpers.PALABRARESERVADA}))
		}

		//Sino
		if l.R.RegexConditionSi.StartsWithSino(currentLine, lineIndex) {
			if len(l.BlockQueue) == 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to create a 'Si' condition outside of a BLOCK", currentLine)
			}

			l.R.RegexConditionSi.ValidateCondition(currentLine, lineIndex)

			//TODO: Get Params

			l.GL.Printf("%+v Found 'Sino' condition [Line: %+v]", funcName, lineIndex)

			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"sino", helpers.PALABRARESERVADA}))
		}

		//Switch: Sea
		if l.R.RegexConditionSwitch.StartsWithSea(currentLine, lineIndex) {
			if len(l.BlockQueue) == 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to create a 'Sea' instruction outside of a BLOCK", currentLine)
			}

			if l.BlockQueue[len(l.BlockQueue)-1] != models.INITBLOCK && l.BlockQueue[len(l.BlockQueue)-2] != models.CUANDOBLOCK {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to create a 'Sea' instruction outside of a CUANDOBLOCK", currentLine)
			}

			//TODO: Get Params

			l.GL.Printf("%+v Found 'Sea' instruction for CUANDOBLOCK [Line: %+v]", funcName, lineIndex)
			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"sea", helpers.PALABRARESERVADA}))
		}
		//Switch: Otro
		if l.R.RegexConditionSwitch.StartsWithOtro(currentLine, lineIndex) {
			if len(l.BlockQueue) == 0 {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to create a 'Otro' instruction outside of a BLOCK", currentLine)
			}
			if l.BlockQueue[len(l.BlockQueue)-1] != models.INITBLOCK && l.BlockQueue[len(l.BlockQueue)-2] != models.CUANDOBLOCK {
				l.LogError(lineIndex, "N/A", "N/A", "Attempted to create a 'Otro' instruction outside of a CUANDOBLOCK", currentLine)
			}

			//TODO: Get Params

			l.GL.Printf("%+v Found 'Otro' instruction for CUANDOBLOCK [Line: %+v]", funcName, lineIndex)
			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"otro", helpers.PALABRARESERVADA}))
		}

		//Regresa
		if l.R.RegexRegresa.MatchRegresa(currentLine, lineIndex) {
			if !l.R.RegexRegresa.MatchPC(currentLine, lineIndex) {
				l.LogError(lineIndex, len(currentLine)-1, ";", "Missing ';'", currentLine)
			}
			currentLine = strings.TrimSuffix(currentLine, ";")
			currentLine = strings.TrimSuffix(currentLine, ")")

			data := strings.Split(currentLine, "(")
			currentLine = ""
			for _, str := range data[1:] {
				currentLine += str + " "
			}
			// params := strings.Split(currentLine, ",")
			// l.OpQueue = []models.TokenComp{}
			// for _, str := range params {
			// 	l.AnalyzeForItem(str, lineIndex)
			// }

			// if !l.ExpectIdent(currentLine, lineIndex) {
			// 	l.LogError(lineIndex, "N/A", "N/A", "Expected <Ident> in parameters", currentLine)
			// }

			l.GL.Printf("%+v Found 'Regresa' instruction [Line: %+v]", funcName, lineIndex)

			l.LL.Print(helpers.IndentStringInLines(helpers.LEXINDENT, 2, []string{
				data[0], helpers.PALABRARESERVADA,
				"(", helpers.DELIMITADOR,
				")", helpers.DELIMITADOR,
				";", helpers.DELIMITADOR,
			}))
		}

		//Desde
		if l.R.RegexLoopDesde.StartsWithDesde(currentLine, lineIndex) {
			//TODO: Analyze
			l.GL.Printf("%+v Found 'Desde' instruction [Line: %+v]", funcName, lineIndex)

			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"desde", helpers.PALABRARESERVADA}))
		}

		//Logger
		l.RegisterBlockChange(LastBlockState, debug, funcName, lineIndex)

		/* Data Segregator */
		if l.CurrentBlockType == models.CONSTANTBLOCK {
			l.NextConstant(currentLine, lineIndex, debug)
		}

		if l.CurrentBlockType == models.VARIABLEBLOCK {
			l.NextVariable(currentLine, lineIndex, debug)
		}

		if l.CurrentBlockType == models.FUNCTIONPROTOBLOCK {
			l.NextFuncionProto(currentLine, lineIndex, debug)
		}

		if l.CurrentBlockType == models.PROCEDUREPROTOBLOCK {
			l.NextProcedureProto(currentLine, lineIndex, debug)
		}

		lineIndex++
	}

	return nil
}

//AnalyzeForItem ...
func (l *LexicalAnalyzer) AnalyzeForItem(str string, lineIndex int64) {
	str = strings.TrimSpace(str)
	if l.R.RegexCustom.MatchCteLog(str, lineIndex) {
		l.OpQueue = append(l.OpQueue, models.CTELOG)
		return
	}
	if l.R.RegexCustom.MatchCteEnt(str) {
		l.OpQueue = append(l.OpQueue, models.CTEENT)
		return
	}
	if l.R.RegexCustom.MatchCteAlfa(str) {
		l.OpQueue = append(l.OpQueue, models.CTEALFA)
		return
	}
	if l.R.RegexCustom.MatchCteReal(str) {
		l.OpQueue = append(l.OpQueue, models.CTEREAL)
		return
	}
	if l.R.RegexCustom.MatchOpArit(str) {
		l.OpQueue = append(l.OpQueue, models.OPARIT)
		return
	}
	if l.R.RegexCustom.MatchOpLog(str) {
		l.OpQueue = append(l.OpQueue, models.OPLOG)
		return
	}
	if l.R.RegexCustom.MatchOpRel(str) {
		l.OpQueue = append(l.OpQueue, models.OPREL)
		return
	}
	if l.R.RegexCustom.MatchIdent(str) {
		l.OpQueue = append(l.OpQueue, models.ID)
		return
	}

	l.OpQueue = append(l.OpQueue, models.NONE)
}

//LogError ...
//"# Linea | # Columna | Error | Descripcion | Linea del Error"
func (l *LexicalAnalyzer) LogError(lineIndex int64, columnIndex interface{}, err string, description string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	l.GL.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	l.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, description, currentLine)
}

//LogErrorGeneral ...
//"# Linea | # Columna | Error | Descripcion | Linea del Error"
func (l *LexicalAnalyzer) LogErrorGeneral(lineIndex int64, columnIndex interface{}, err string, description string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	l.GL.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
}

//RegisterBlockChange ...
func (l *LexicalAnalyzer) RegisterBlockChange(LastBlockState models.BlockType, debug bool, funcName string, lineIndex int64) {
	if LastBlockState != l.CurrentBlockType {
		l.GL.Printf("%+vSwitched to %+v [%+v]", funcName, l.CurrentBlockType, lineIndex)
		if debug {
			log.Printf("Switched to %+v [%+v]", l.CurrentBlockType, lineIndex)
		}
	}
}

//Print ...
func (l *LexicalAnalyzer) Print() {
	log.SetFlags(0)
	if len(l.ConstantStorage) > 0 {
		log.Print("Constants: ")
		pprnt.Print(l.ConstantStorage)
		log.Print("\n")
	} else {
		log.Println("Constants: 0")
	}

	if len(l.VariableStorage) > 0 {
		log.Print("Variables: ")
		pprnt.Print(l.VariableStorage)
		log.Print("\n")
	} else {
		log.Println("Variables: 0")
	}

	log.SetFlags(log.LstdFlags)
}
