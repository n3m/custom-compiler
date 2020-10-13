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
			currentLine = strings.TrimSuffix(currentLine, ";")
			l.GL.Printf("%+v[VARIABLE] Alfabetico Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Alfabetico Found > %+v", currentLine)
			}

			return
		}
		if l.R.RegexVariableEntero.MatchVariableEntero(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			l.GL.Printf("%+v[VARIABLE] Entero Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Entero Found > %+v", currentLine)
			}
			return
		}
		if l.R.RegexVariableFlotante.MatchVariableFlotante(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			l.GL.Printf("%+v[VARIABLE] Flotante Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Flotante Found > %+v", currentLine)
			}
			return
		}
		if l.R.RegexVariableLogico.MatchVariableLogico(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			l.GL.Printf("%+v[VARIABLE] Logico Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Logico Found > %+v", currentLine)
			}
			return
		}
		if l.R.RegexVariableReal.MatchVariableReal(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			l.GL.Printf("%+v[VARIABLE] Real Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Real Found > %+v", currentLine)
			}
			return
		}

	}

}
