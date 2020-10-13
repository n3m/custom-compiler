package lexyc

import (
	"go-custom-compiler/models"
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
			_, variableData := getVariablesFromString(currentLine)
			for _, name := range variableData {
				l.VariableStorage = append(l.VariableStorage, models.Token{Type: models.ALFABETICO, Key: name})
			}
			l.GL.Printf("%+v[VARIABLE] Alfabetico Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Alfabetico Found > %+v", currentLine)
			}

			return
		}

		if l.R.RegexVariableEntero.MatchVariableEntero(currentLine) {
			_, variableData := getVariablesFromString(currentLine)
			for _, name := range variableData {
				l.VariableStorage = append(l.VariableStorage, models.Token{Type: models.ENTERO, Key: name})
			}
			l.GL.Printf("%+v[VARIABLE] Alfabetico Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Alfabetico Found > %+v", currentLine)
			}

			l.GL.Printf("%+v[VARIABLE] Entero Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Entero Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexVariableFlotante.MatchVariableFlotante(currentLine) {
			_, variableData := getVariablesFromString(currentLine)
			for _, name := range variableData {
				l.VariableStorage = append(l.VariableStorage, models.Token{Type: models.FLOTANTE, Key: name})
			}
			l.GL.Printf("%+v[VARIABLE] Alfabetico Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Alfabetico Found > %+v", currentLine)
			}

			l.GL.Printf("%+v[VARIABLE] Flotante Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Flotante Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexVariableLogico.MatchVariableLogico(currentLine) {
			_, variableData := getVariablesFromString(currentLine)
			for _, name := range variableData {
				l.VariableStorage = append(l.VariableStorage, models.Token{Type: models.LOGICO, Key: name})
			}
			l.GL.Printf("%+v[VARIABLE] Alfabetico Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Alfabetico Found > %+v", currentLine)
			}

			l.GL.Printf("%+v[VARIABLE] Logico Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Logico Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexVariableReal.MatchVariableReal(currentLine) {
			_, variableData := getVariablesFromString(currentLine)
			for _, name := range variableData {
				l.VariableStorage = append(l.VariableStorage, models.Token{Type: models.REAL, Key: name})
			}
			l.GL.Printf("%+v[VARIABLE] Alfabetico Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Alfabetico Found > %+v", currentLine)
			}

			l.GL.Printf("%+v[VARIABLE] Real Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Real Found > %+v", currentLine)
			}
			return
		}

	}

}

func getVariablesFromString(currentLine string) (string, []string) {
	currentLine = strings.TrimSuffix(currentLine, ";")
	currentLine = strings.TrimSuffix(currentLine, " ")
	lineData := strings.Split(currentLine, ":")
	varType := lineData[1]
	variables := lineData[0]
	variableData := strings.Split(variables, ",")
	return varType, variableData
}
