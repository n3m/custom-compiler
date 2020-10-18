package lexyc

import (
	"bufio"
	"fmt"
	"log"

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
	ConstantStorage  []models.Token
	VariableStorage  []models.Token
}

//NewLexicalAnalyzer ...
func NewLexicalAnalyzer(file *bufio.Scanner, ErrorLogger, LexLogger, GeneralLogger *log.Logger) (*LexicalAnalyzer, error) {
	var moduleName string = "[Lexyc][NewLexicalAnalyzer()]"

	if file == nil {
		GeneralLogger.Printf("[ERROR]%+v file is not present", moduleName)
		return nil, fmt.Errorf("[ERROR]%+v file is not present", moduleName)
	}
	if ErrorLogger == nil || LexLogger == nil || GeneralLogger == nil {
		GeneralLogger.Printf("[ERROR]%+v Loggers are not present", moduleName)
		return nil, fmt.Errorf("[ERROR]%+v Loggers are not present", moduleName)
	}
	R, err := regex.NewRegex(ErrorLogger, LexLogger, GeneralLogger)
	if err != nil {
		GeneralLogger.Printf("[ERROR]%+v %+v", moduleName, err.Error())
		return nil, fmt.Errorf("[ERROR]%+v %+v", moduleName, err.Error())
	}

	LexLogger.Println("----------------------------------------------")
	LexLogger.Println(helpers.IndentString(helpers.LEXINDENT, []string{"Lexema", "Token"}))
	LexLogger.Println("----------------------------------------------")

	return &LexicalAnalyzer{
		File: file,
		R:    R,
		EL:   ErrorLogger,
		LL:   LexLogger,
		GL:   GeneralLogger,

		CurrentBlockType: models.DEFAULTBLOCK,
		ConstantStorage:  []models.Token{},
		VariableStorage:  []models.Token{},
	}, nil
}

//Analyze ...
func (l *LexicalAnalyzer) Analyze(debug bool) error {
	funcName := "[Analyze()] "
	var lineIndex int64 = 1
	for l.File.Scan() {
		currentLine := l.File.Text()
		if len(currentLine) == 0 {
			lineIndex++

			continue
		}
		var LastBlockState models.BlockType
		LastBlockState = l.CurrentBlockType
		/* Type Validation */
		if isComment, err := l.R.StartsWith("//", currentLine); err != nil {
			l.GL.Printf("%+v[APP_ERR] %+v", funcName, err.Error())
			return fmt.Errorf("%+v[APP_ERR] %+v", funcName, err.Error())
		} else {
			if isComment {
				l.GL.Printf("%+vSkipping Comment at line %+v", funcName, lineIndex)
				log.Printf("Skipping Comment at line %+v", lineIndex)
				lineIndex++

				continue
			}
		}

		if l.R.RegexConstante.StartsWithConstante(currentLine) {
			l.CurrentBlockType = models.CONSTANTBLOCK
			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"constantes", helpers.PALABRARESERVADA}))
		}

		if l.R.RegexVariable.StartsWithVariable(currentLine) {
			l.CurrentBlockType = models.VARIABLEBLOCK
			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"variables", helpers.PALABRARESERVADA}))
		}

		if l.R.RegexFuncionProto.StartsWithFuncionProto(currentLine) {
			l.CurrentBlockType = models.FUNCTIONPROTOBLOCK
			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"funcion", helpers.PALABRARESERVADA}))
		}

		if l.R.RegexProcedure.StartsWithProcedure(currentLine) {
			l.CurrentBlockType = models.PROCEDUREBLOCK
			l.LL.Println(helpers.IndentString(helpers.LEXINDENT, []string{"procedimiento", helpers.PALABRARESERVADA}))
		}
		if l.R.RegexFunction.StartsWithFunction(currentLine) {
			l.CurrentBlockType = models.FUNCTIONBLOCK
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
