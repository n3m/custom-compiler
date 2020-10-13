package lexyc

import (
	"log"
	"strings"
)

//NextVariable ...
func (l *LexicalAnalyzer) NextVariable(currentLine string, debug bool) {
	// var moduleName string = "[variables.go][NextVariable()]"
	funcName := "[NextVariable()] "

	if l.CurrentBlockType == VARIABLEBLOCK {
		if l.R.RegexVariable.StartsWithVariableNoCheck(currentLine) {
			data := strings.Split(currentLine, " ")
			currentLine = ""
			for _, str := range data[1:] {
				currentLine += str + " "
			}
		}
		currentLine = strings.TrimSpace(currentLine)

		if l.R.RegexConstante.StartsWithConstante(currentLine) {
			l.CurrentBlockType = CONSTANTBLOCK
			l.GL.Printf("%+vSwitched to CONSTANTBLOCK while analyzing for VARIABLEBLOCK", funcName)

			if debug {
				log.Printf("Switched to CONSTANTBLOCK while analyzing for VARIABLEBLOCK")
			}
			return
		}

		log.Printf("%+v > %+v", funcName, currentLine)
		//^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:[a-zA-Z]+;$

		// if l.R.RegexFloat.MatchFloatConstantDeclaration(currentLine) {
		// 	currentLine = strings.TrimSuffix(currentLine, ";")
		// 	floatData := strings.Split(currentLine, ":=")
		// 	l.FloatConstants[floatData[0]] = floatData[1]
		// 	l.GL.Printf("%+v[CONSTANT] Float Found > %+v", funcName, currentLine)
		// 	if debug {
		// 		log.Printf("[CONSTANT] Float Found > %+v", currentLine)
		// 	}
		// 	return
		// }
		// if l.R.RegexInt.MatchIntConstantDeclaration(currentLine) {
		// 	currentLine = strings.TrimSuffix(currentLine, ";")
		// 	floatData := strings.Split(currentLine, ":=")
		// 	l.IntConstants[floatData[0]] = floatData[1]
		// 	l.GL.Printf("%+v[CONSTANT] Int Found > %+v", funcName, currentLine)

		// 	if debug {
		// 		log.Printf("[CONSTANT] Int Found > %+v", currentLine)
		// 	}
		// 	return
		// }

	}

}
