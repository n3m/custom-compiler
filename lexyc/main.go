package lexyc

import (
	"bufio"
	"fmt"
	"go-custom-compiler/regex"
	"log"
	"strings"
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
	File *bufio.Scanner
	R    *regex.CustomRegex
	EL   *log.Logger
	LL   *log.Logger

	//TEST
	CurrentBlockType BlockType
	Constants        map[string]interface{}
	Variables        map[string]interface{}
}

//NewLexicalAnalyzer ...
func NewLexicalAnalyzer(file *bufio.Scanner, ErrorLogger, LexLogger *log.Logger) (*LexicalAnalyzer, error) {
	if file == nil {
		return nil, fmt.Errorf("[ERROR] file is not present")
	}
	R, err := regex.NewRegex()
	if err != nil {
		return nil, fmt.Errorf("[ERROR] %+v", err)
	}
	if ErrorLogger == nil || LexLogger == nil {
		return nil, fmt.Errorf("[ERROR] Loggers are not present")
	}

	return &LexicalAnalyzer{
		File: file,
		R:    R,
		EL:   ErrorLogger,
		LL:   LexLogger,

		CurrentBlockType: DEFAULTBLOCK,
		Constants:        make(map[string]interface{}),
		Variables:        make(map[string]interface{}),
	}, nil
}

//Analyze ...
func (l *LexicalAnalyzer) Analyze() error {
	for l.File.Scan() {
		currentLine := l.File.Text()

		if l.CurrentBlockType == DEFAULTBLOCK {
			if l.R.RegexConstante.StartsWithConstante(currentLine) {
				// log.Printf("Inicia con constante! > %+v", currentLine)
				l.CurrentBlockType = CONSTANTBLOCK
			}

			if l.R.RegexVariable.StartsWithVariable(currentLine) {
				// log.Printf("Inicia con Variable! > %+v", currentLine)
				l.CurrentBlockType = VARIABLEBLOCK
			}
		}

		if l.CurrentBlockType == CONSTANTBLOCK {
			if l.R.RegexConstante.StartsWithConstante(currentLine) {
				data := strings.Split(currentLine, " ")
				currentLine = ""
				for _, str := range data[1:] {
					currentLine += str + " "
				}
			}
			currentLine = strings.TrimSpace(currentLine)

			if l.R.RegexVariable.StartsWithVariable(currentLine) {
				l.CurrentBlockType = VARIABLEBLOCK
			}

			if l.CurrentBlockType == CONSTANTBLOCK && l.R.RegexFloat.MatchFloatConstantDeclaration(currentLine) {
				log.Printf("[CONSTANT] Float Found %+v", currentLine)
				continue
				//ADD TO FLOAT CONSTANTS
			}
			if l.CurrentBlockType == CONSTANTBLOCK && l.R.RegexInt.MatchIntConstantDeclaration(currentLine) {
				log.Printf("[CONSTANT] Int Found %+v", currentLine)
				continue
				//ADD TO INT CONSTANTS
			}

		}

		if l.CurrentBlockType == VARIABLEBLOCK {

		}

	}

	return nil
}

//NextConstant ...
func (l *LexicalAnalyzer) NextConstant() {

}
