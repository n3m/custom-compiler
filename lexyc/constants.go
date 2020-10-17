package lexyc

import (
	"go-custom-compiler/models"
	"log"
	"regexp"
	"strings"
)

//NextConstant ...
func (l *LexicalAnalyzer) NextConstant(currentLine string, lineIndex int64, debug bool) {
	funcName := "[NextConstant()] "
	// var moduleName string = "[constants.go][NextConstant()]"

	if l.CurrentBlockType == models.CONSTANTBLOCK {
		if l.R.RegexConstante.StartsWithConstanteNoCheck(currentLine) {
			data := strings.Split(currentLine, " ")
			currentLine = ""
			for _, str := range data[1:] {
				currentLine += str + " "
			}
		}
		currentLine = strings.TrimSpace(currentLine)

		if l.R.RegexConstanteEntera.MatchEnteraConstantDeclaration(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			constantData := strings.Split(currentLine, ":=")
			l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.ENTERO, Key: constantData[0], Value: constantData[1]})
			l.GL.Printf("%+v[CONSTANT] Entero Found > %+v", funcName, currentLine)

			if debug {
				log.Printf("[CONSTANT] Entero Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexConstanteReal.MatchRealConstantDeclaration(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			constantData := strings.Split(currentLine, ":=")
			l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.REAL, Key: constantData[0], Value: constantData[1]})
			l.GL.Printf("%+v[CONSTANT] Real Found > %+v", funcName, currentLine)

			if debug {
				log.Printf("[CONSTANT] Real Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexConstanteAlfabetica.MatchAlfabeticaConstantDeclaration(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			constantData := strings.Split(currentLine, ":=")
			l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.ALFABETICO, Key: constantData[0], Value: constantData[1]})
			l.GL.Printf("%+v[CONSTANT] Alfabetico Found > %+v", funcName, currentLine)

			if debug {
				log.Printf("[CONSTANT] Alfabetico Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexConstanteLogica.MatchLogicaConstantDeclaration(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			constantData := strings.Split(currentLine, ":=")
			l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.LOGICO, Key: constantData[0], Value: constantData[1]})
			l.GL.Printf("%+v[CONSTANT] Logico Found > %+v", funcName, currentLine)

			if debug {
				log.Printf("[CONSTANT] Logico Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexConstanteDefault.MatchVariableDefault(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			constantData := strings.Split(currentLine, ":=")
			value := constantData[1]
			l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.LOGICO, Key: constantData[0], Value: value})

			// regexEntero := regexp.MustCompile(`([0-9]+|\-[0-9]+)`)
			// regexReal := regexp.MustCompile(`(([0-9]+|\-[0-9]+)\.([0-9]+)|([0-9]+|\-[0-9]+)e[0-9]+)`)
			regexLogico := regexp.MustCompile(`((?i)verdadero|(?i)falso))`)
			// regexAlfabetico := regexp.MustCompile(`(\s*)((\"(\w)*\")`)

			if regexLogico.MatchString(value) {
				foundTypo := false
				keyword := "verdadero"
				keyData := strings.Split(keyword, "")
				for i, char := range value {
					if i < len(keyData)-1 {
						if !foundTypo {
							if string(char) != keyData[i] {
								foundTypo = true
								if debug {
									log.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", value, i, lineIndex, keyword)
								}
								l.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", value, i, lineIndex, keyword)
							}
						}
					}
				}
				if foundTypo {
					return
				}
				keyword = "falso"
				for i, char := range value {
					if i < len(keyData)-1 {
						if !foundTypo {
							if string(char) != keyData[i] {
								foundTypo = true
								if debug {
									log.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", value, i, lineIndex, keyword)
								}
								l.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", value, i, lineIndex, keyword)
							}
						}
					}
				}
				if foundTypo {
					return
				}
			}

		}

		l.GL.Printf("%+v Did not found any type of match on Line[%+v]! ", funcName, lineIndex)

	}
}
