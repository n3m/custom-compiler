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
			// l.GL.Printf("%+v[VARIABLE] Default Found > %+v", funcName, currentLine)
			// if debug {
			// 	log.Printf("[VARIABLE] Default Found > %+v", currentLine)
			// }

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
								if debug {
									log.Printf("Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableAlfabetico.Keyword)
								}
								l.GL.Printf("Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableAlfabetico.Keyword)
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
								if debug {
									log.Printf("Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableEntero.Keyword)
								}
								l.GL.Printf("Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableEntero.Keyword)
							}
						}
					}
				}
			}

			if l.R.RegexVariableFlotante.MatchVariableFlotanteCaseless(typeOfData) {
				for _, name := range variableData {
					l.VariableStorage = append(l.VariableStorage, models.Token{Type: models.FLOTANTE, Key: name})
				}

				l.GL.Printf("%+v[VARIABLE] Flotante Found > %+v", funcName, currentLine)
				if debug {
					log.Printf("[VARIABLE] Flotante Found > %+v", currentLine)
				}

				foundTypo := false
				keyData := strings.Split(l.R.RegexVariableFlotante.Keyword, "")
				for i, char := range typeOfData {
					if i < len(keyData)-1 {
						if !foundTypo {
							if string(char) != keyData[i] {
								foundTypo = true
								if debug {
									log.Printf("Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableFlotante.Keyword)
								}
								l.GL.Printf("Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableFlotante.Keyword)
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
								if debug {
									log.Printf("Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableLogico.Keyword)
								}
								l.GL.Printf("Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableLogico.Keyword)
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
								if debug {
									log.Printf("Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableReal.Keyword)
								}
								l.GL.Printf("Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", typeOfData, i, lineIndex, l.R.RegexVariableReal.Keyword)
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
