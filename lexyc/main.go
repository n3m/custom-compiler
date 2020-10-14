package lexyc

import (
	"bufio"
	"fmt"
	"go-custom-compiler/models"
	"go-custom-compiler/regex"
	"log"

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
		// l.GL.Printf("%+vAnalyzing Line: '%+v'", funcName, currentLine)

		/* Type Validation */

		if l.R.RegexConstante.StartsWithConstante(currentLine) {
			l.CurrentBlockType = models.CONSTANTBLOCK
			l.GL.Printf("%+vSwitched to CONSTANTBLOCK", funcName)
			if debug {
				log.Printf("Switched to CONSTANTBLOCK")
			}
		}

		if l.R.RegexVariable.StartsWithVariable(currentLine) {
			l.CurrentBlockType = models.VARIABLEBLOCK
			l.GL.Printf("%+vSwitched to VARIABLEBLOCK", funcName)

			if debug {
				log.Printf("Switched to VARIABLEBLOCK")
			}
		}

		if l.R.RegexFuncionProto.StartsWithFuncionProto(currentLine) {
			l.CurrentBlockType = models.FUNCTIONPROTOBLOCK
			l.GL.Printf("%+vSwitched to FUNCTIONPROTOBLOCK", funcName)

			if debug {
				log.Printf("Switched to FUNCTIONPROTOBLOCK")
			}
		}

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
		lineIndex++
	}

	return nil
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
