package lexyc

import (
	"bufio"
	"fmt"
	"go-custom-compiler/regex"
	"log"

	"github.com/DrN3MESiS/pprnt"
)

//BlockType ...
type BlockType int

const (
	//DEFAULTBLOCK ...
	DEFAULTBLOCK BlockType = iota
	//CONSTANTBLOCK ...
	CONSTANTBLOCK
	//VARIABLEBLOCK ...
	VARIABLEBLOCK
)

//LexicalAnalyzer ...
type LexicalAnalyzer struct {
	File *bufio.Scanner     //File
	R    *regex.CustomRegex //Regex Handler
	EL   *log.Logger        //Error Logger
	LL   *log.Logger        //Lex Logger
	GL   *log.Logger        //General Logger

	//TEST
	CurrentBlockType BlockType
	FloatConstants   map[string]interface{}
	IntConstants     map[string]interface{}
	StringConstants  map[string]interface{}
	FloatVariables   map[string]interface{}
	IntVariables     map[string]interface{}
	StringVariables  map[string]interface{}
}

//NewLexicalAnalyzer ...
func NewLexicalAnalyzer(file *bufio.Scanner, ErrorLogger, LexLogger, GeneralLogger *log.Logger) (*LexicalAnalyzer, error) {
	var moduleName string = "[Lexyc][NewLexicalAnalyzer()]"

	if file == nil {
		return nil, fmt.Errorf("[ERROR]%+v file is not present", moduleName)
	}
	if ErrorLogger == nil || LexLogger == nil || GeneralLogger == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers are not present", moduleName)
	}
	R, err := regex.NewRegex(ErrorLogger, LexLogger, GeneralLogger)
	if err != nil {
		return nil, fmt.Errorf("[ERROR]%+v %+v", moduleName, err)
	}

	return &LexicalAnalyzer{
		File: file,
		R:    R,
		EL:   ErrorLogger,
		LL:   LexLogger,
		GL:   GeneralLogger,

		CurrentBlockType: DEFAULTBLOCK,
		FloatConstants:   make(map[string]interface{}),
		IntConstants:     make(map[string]interface{}),
		StringConstants:  make(map[string]interface{}),
		FloatVariables:   make(map[string]interface{}),
		IntVariables:     make(map[string]interface{}),
		StringVariables:  make(map[string]interface{}),
	}, nil
}

//Analyze ...
func (l *LexicalAnalyzer) Analyze(debug bool) error {
	funcName := "[Analyze()] "
	for l.File.Scan() {
		currentLine := l.File.Text()

		/* Type Validation */

		if l.R.RegexConstante.StartsWithConstante(currentLine) {
			l.CurrentBlockType = CONSTANTBLOCK
			l.GL.Printf("%+vSwitched to CONSTANTBLOCK", funcName)
			if debug {
				log.Printf("Switched to CONSTANTBLOCK")
			}
		}

		if l.R.RegexVariable.StartsWithVariable(currentLine) {
			l.CurrentBlockType = VARIABLEBLOCK
			l.GL.Printf("%+vSwitched to VARIABLEBLOCK", funcName)

			if debug {
				log.Printf("Switched to VARIABLEBLOCK")
			}
		}

		/* Data Segregator */

		if l.CurrentBlockType == CONSTANTBLOCK {
			l.NextConstant(currentLine, debug)
		}

		if l.CurrentBlockType == VARIABLEBLOCK {
			l.NextVariable(currentLine, debug)
		}

	}

	return nil
}

//Print ...
func (l *LexicalAnalyzer) Print() {
	log.SetFlags(0)
	if len(l.FloatConstants) > 0 {
		log.Print("Float Constants: ")
		pprnt.Print(l.FloatConstants)
	} else {
		log.Println("Float Constants: 0")
	}

	if len(l.IntConstants) > 0 {
		log.Print("Int Constants: ")
		pprnt.Print(l.IntConstants)
	} else {
		log.Println("Int Constants: 0")
	}

	if len(l.StringConstants) > 0 {
		log.Print("String Constants: ")
		pprnt.Print(l.StringConstants)
	} else {
		log.Println("String Constants: 0")
	}

	if len(l.FloatVariables) > 0 {
		log.Print("Float Variables: ")
		pprnt.Print(l.FloatVariables)
	} else {
		log.Println("Float Variables: 0")
	}

	if len(l.IntVariables) > 0 {
		log.Print("Int Variables: ")
		pprnt.Print(l.IntVariables)
	} else {
		log.Println("Int Variables: 0")
	}

	if len(l.StringVariables) > 0 {
		log.Print("String Variables: ")
		pprnt.Print(l.StringVariables)
	} else {
		log.Println("String Variables: 0")
	}

	log.SetFlags(log.LstdFlags)
}
