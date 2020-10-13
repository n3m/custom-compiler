package lexyc

import (
	"go-custom-compiler/models"
	"log"
	"strings"
)

//NextConstant ...
func (l *LexicalAnalyzer) NextConstant(currentLine string, debug bool) {
	funcName := "[NextConstant()] "
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
			l.GL.Printf("%+vSwitched to VARIABLEBLOCK while analyzing for CONSTANTBLOCK", funcName)

			if debug {
				log.Printf("Switched to VARIABLEBLOCK while analyzing for CONSTANTBLOCK")
			}
			return
		}

		if l.R.RegexConstanteFloat.MatchFloatConstantDeclaration(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			floatData := strings.Split(currentLine, ":=")
			l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.FLOTANTE, Key: floatData[0], Value: floatData[1]})
			l.GL.Printf("%+v[CONSTANT] Float Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[CONSTANT] Float Found > %+v", currentLine)
			}
			return
		}
		if l.R.RegexConstanteInt.MatchIntConstantDeclaration(currentLine) {
			currentLine = strings.TrimSuffix(currentLine, ";")
			floatData := strings.Split(currentLine, ":=")
			l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.ENTERO, Key: floatData[0], Value: floatData[1]})
			l.GL.Printf("%+v[CONSTANT] Int Found > %+v", funcName, currentLine)

			if debug {
				log.Printf("[CONSTANT] Int Found > %+v", currentLine)
			}
			return
		}

	}
}
