package lexyc

import (
	"go-custom-compiler/models"
	"log"
	"strings"
)

//^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:[a-zA-Z]+;$

//NextVariable ...
func (l *LexicalAnalyzer) NextVariable(currentLine string, lineIndex int64, debug bool) {
	// var moduleName string = "[variables.go][NextVariable()]"
	funcName := "[NextVariable()] "

	if l.CurrentBlockType == models.VARIABLEBLOCK {
		if l.R.RegexVariable.StartsWithVariableNoCheck(currentLine) {
			data := strings.Split(currentLine, " ")
			currentLine = ""
			for _, str := range data[1:] {
				currentLine += str + " "
			}
		}
		currentLine = strings.TrimSpace(currentLine)

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

			l.GL.Printf("%+v[VARIABLE] Entero Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Entero Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexVariableLogico.MatchVariableLogico(currentLine) {
			_, variableData := getVariablesFromString(currentLine)
			for _, name := range variableData {
				l.VariableStorage = append(l.VariableStorage, models.Token{Type: models.LOGICO, Key: name})
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

			l.GL.Printf("%+v[VARIABLE] Real Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[VARIABLE] Real Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexVariableDefault.MatchVariableDefault(currentLine) {
			typeOfData, variableData := getVariablesFromString(currentLine)

			if l.R.RegexVariableAlfabetico.MatchVariableAlfabeticoCaseless(typeOfData) {
				for _, name := range variableData {
					l.VariableStorage = append(l.VariableStorage, models.Token{Type: models.ALFABETICO, Key: name})
				}

				l.GL.Printf("%+v[VARIABLE] Alfabetico Found > %+v", funcName, currentLine)
				if debug {
					log.Printf("[VARIABLE] Alfabetico Found > %+v", currentLine)
				}

				foundTypo := false
				keyData := strings.Split(l.R.RegexVariableAlfabetico.Keyword, "")
				for i, char := range typeOfData {
					if i < len(keyData)-1 {
						if !foundTypo {
							if string(char) != keyData[i] {
								foundTypo = true
								log.Printf("[ERR] Found typo in '%+v' declaration at [%+v][Line: %+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableAlfabetico.Keyword)
								l.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v][Line: %+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableAlfabetico.Keyword)
								//"# Linea | # Columna | Error | Descripcion | Linea del Error"
								l.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, i, typeOfData, l.R.RegexVariableAlfabetico.Keyword, currentLine)
							}
						}
					}
				}
			}

			if l.R.RegexVariableEntero.MatchVariableEnteroCaseless(typeOfData) {
				for _, name := range variableData {
					l.VariableStorage = append(l.VariableStorage, models.Token{Type: models.ENTERO, Key: name})
				}

				l.GL.Printf("%+v[VARIABLE] Entero Found > %+v", funcName, currentLine)
				if debug {
					log.Printf("[VARIABLE] Entero Found > %+v", currentLine)
				}

				foundTypo := false
				keyData := strings.Split(l.R.RegexVariableEntero.Keyword, "")
				for i, char := range typeOfData {
					if i < len(keyData)-1 {
						if !foundTypo {
							if string(char) != keyData[i] {
								foundTypo = true
								log.Printf("[ERR] Found typo in '%+v' declaration at [%+v][Line: %+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableEntero.Keyword)
								l.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v][Line: %+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableEntero.Keyword)
								//"# Linea | # Columna | Error | Descripcion | Linea del Error"
								l.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, i, typeOfData, l.R.RegexVariableEntero.Keyword, currentLine)
							}
						}
					}
				}
			}

			if l.R.RegexVariableLogico.MatchVariableLogicoCaseless(typeOfData) {
				for _, name := range variableData {
					l.VariableStorage = append(l.VariableStorage, models.Token{Type: models.LOGICO, Key: name})
				}

				l.GL.Printf("%+v[VARIABLE] Logico Found > %+v", funcName, currentLine)
				if debug {
					log.Printf("[VARIABLE] Logico Found > %+v", currentLine)
				}

				foundTypo := false
				keyData := strings.Split(l.R.RegexVariableLogico.Keyword, "")
				for i, char := range typeOfData {
					if i < len(keyData)-1 {
						if !foundTypo {
							if string(char) != keyData[i] {
								foundTypo = true
								log.Printf("[ERR] Found typo in '%+v' declaration at [%+v][Line: %+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableLogico.Keyword)
								l.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v][Line: %+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableLogico.Keyword)
								//"# Linea | # Columna | Error | Descripcion | Linea del Error"
								l.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, i, typeOfData, l.R.RegexVariableLogico.Keyword, currentLine)
							}
						}
					}
				}
			}

			if l.R.RegexVariableReal.MatchVariableRealCaseless(typeOfData) {
				for _, name := range variableData {
					l.VariableStorage = append(l.VariableStorage, models.Token{Type: models.REAL, Key: name})
				}

				l.GL.Printf("%+v[VARIABLE] Real Found > %+v", funcName, currentLine)
				if debug {
					log.Printf("[VARIABLE] Real Found > %+v", currentLine)
				}

				foundTypo := false
				keyData := strings.Split(l.R.RegexVariableReal.Keyword, "")
				for i, char := range typeOfData {
					if i < len(keyData)-1 {
						if !foundTypo {
							if string(char) != keyData[i] {
								foundTypo = true
								log.Printf("[ERR] Found typo in '%+v' declaration at [%+v][Line: %+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableReal.Keyword)
								l.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v][Line: %+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableReal.Keyword)
								//"# Linea | # Columna | Error | Descripcion | Linea del Error"
								l.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, i, typeOfData, l.R.RegexVariableReal.Keyword, currentLine)
							}
						}
					}
				}
			}

			return
		}

		l.GL.Printf("%+v Did not found any type of match on Line[%+v]! ", funcName, lineIndex)

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
