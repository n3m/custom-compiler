package lexyc

import (
	"go-custom-compiler/models"
	"strings"
)

//NextFuncionProto ...
func (l *LexicalAnalyzer) NextFuncionProto(currentLine string, lineIndex int64, debug bool) {
	funcName := "[NextFuncionProto()] "
	// var moduleName string = "[constants.go][NextFuncionProto()]"

	if l.CurrentBlockType == models.FUNCTIONBLOCK {
		if l.R.RegexFuncionProto.StartsWithFuncionProtoNoCheck(currentLine) {
			data := strings.Split(currentLine, " ")
			currentLine = ""
			for _, str := range data[1:] {
				currentLine += str + " "
			}
		}
		currentLine = strings.TrimSpace(currentLine)

		// if l.R.RegexConstanteFloat.MatchFloatConstantDeclaration(currentLine) {
		// 	currentLine = strings.TrimSuffix(currentLine, ";")
		// 	floatData := strings.Split(currentLine, ":=")
		// 	l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.FLOTANTE, Key: floatData[0], Value: floatData[1]})
		// 	l.GL.Printf("%+v[CONSTANT] Float Found > %+v", funcName, currentLine)
		// 	if debug {
		// 		log.Printf("[CONSTANT] Float Found > %+v", currentLine)
		// 	}
		// 	return
		// }
		// if l.R.RegexConstanteInt.MatchIntConstantDeclaration(currentLine) {
		// 	currentLine = strings.TrimSuffix(currentLine, ";")
		// 	floatData := strings.Split(currentLine, ":=")
		// 	l.ConstantStorage = append(l.ConstantStorage, models.Token{Type: models.ENTERO, Key: floatData[0], Value: floatData[1]})
		// 	l.GL.Printf("%+v[CONSTANT] Int Found > %+v", funcName, currentLine)

		// 	if debug {
		// 		log.Printf("[CONSTANT] Int Found > %+v", currentLine)
		// 	}
		// 	return
		// }

		l.GL.Printf("%+v Did not found any type of match on Line[%+v]! ", funcName, lineIndex)

	}
}
