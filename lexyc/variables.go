package lexyc

import (
	"log"
	"strings"
)

//^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:[a-zA-Z]+;$

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

		if l.R.RegexVariableAlfabetico.MatchVariableAlfabetico(currentLine) {
			return
		}
		if l.R.RegexVariableEntero.MatchVariableEntero(currentLine) {
			return
		}
		if l.R.RegexVariableFlotante.MatchVariableFlotante(currentLine) {
			return
		}
		if l.R.RegexVariableLogico.MatchVariableLogico(currentLine) {
			return
		}
		if l.R.RegexVariableReal.MatchVariableReal(currentLine) {
			return
		}

	}

}
