package lexyc

import (
	"log"
	"strings"
)

//NextConstant ...
func (l *LexicalAnalyzer) NextConstant(currentLine string, debug bool) {
	// var moduleName string = "[constants.go][NextConstant()]"

	if l.CurrentBlockType == CONSTANTBLOCK {
		if l.R.RegexConstante.StartsWithConstanteNoCheck(currentLine) {
			data := strings.Split(currentLine, " ")
			currentLine = ""
			for _, str := range data[1:] {
				currentLine += str + " "
			}
		}
		currentLine = strings.TrimSpace(currentLine)

		if l.R.RegexVariable.StartsWithVariable(currentLine) {
			l.CurrentBlockType = VARIABLEBLOCK
			if debug {
				log.Printf("Switched to VARIABLEBLOCK while analyzing for Constant")
			}
			return
		}

		if l.R.RegexFloat.MatchFloatConstantDeclaration(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			floatData := strings.Split(currentLine, ":=")
			l.FloatConstants[floatData[0]] = floatData[1]
			if debug {
				log.Printf("[CONSTANT] Float Found %+v", currentLine)
			}
			return
		}
		if l.R.RegexInt.MatchIntConstantDeclaration(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			floatData := strings.Split(currentLine, ":=")
			l.IntConstants[floatData[0]] = floatData[1]
			if debug {
				log.Printf("[CONSTANT] Int Found %+v", currentLine)
			}
			return
		}

	}
}
