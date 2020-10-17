package lexyc

import (
	"go-custom-compiler/models"
	"log"
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
			floatData := strings.Split(currentLine, ":=")
			l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.ENTERO, Key: floatData[0], Value: floatData[1]})
			l.GL.Printf("%+v[CONSTANT] Entero Found > %+v", funcName, currentLine)

			if debug {
				log.Printf("[CONSTANT] Entero Found > %+v", currentLine)
			}
			return
		}
		if l.R.RegexConstanteReal.MatchRealConstantDeclaration(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			floatData := strings.Split(currentLine, ":=")
			l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.REAL, Key: floatData[0], Value: floatData[1]})
			l.GL.Printf("%+v[CONSTANT] Real Found > %+v", funcName, currentLine)

			if debug {
				log.Printf("[CONSTANT] Real Found > %+v", currentLine)
			}
			return
		}
		if l.R.RegexConstanteAlfabetica.MatchAlfabeticaConstantDeclaration(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			floatData := strings.Split(currentLine, ":=")
			l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.ALFABETICO, Key: floatData[0], Value: floatData[1]})
			l.GL.Printf("%+v[CONSTANT] Alfabetico Found > %+v", funcName, currentLine)

			if debug {
				log.Printf("[CONSTANT] Alfabetico Found > %+v", currentLine)
			}
			return
		}
		if l.R.RegexConstanteLogica.MatchLogicaConstantDeclaration(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			floatData := strings.Split(currentLine, ":=")
			l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.LOGICO, Key: floatData[0], Value: floatData[1]})
			l.GL.Printf("%+v[CONSTANT] Logico Found > %+v", funcName, currentLine)

			if debug {
				log.Printf("[CONSTANT] Logico Found > %+v", currentLine)
			}
			return
		}

		l.GL.Printf("%+v Did not found any type of match on Line[%+v]! ", funcName, lineIndex)

	}
}
