package lexyc

import (
	"go-custom-compiler/models"
	"log"
	"strings"
)

//NextFuncionProto ...
func (l *LexicalAnalyzer) NextFuncionProto(currentLine string, lineIndex int64, debug bool) {
	funcName := "[NextFuncionProto()] "
	// var moduleName string = "[regexfunctionproto][NextFuncionProto()]"

	if l.CurrentBlockType == models.FUNCTIONPROTOBLOCK {
		if l.R.RegexFuncionProto.StartsWithFuncionProtoNoCheck(currentLine) {
			data := strings.Split(currentLine, " ")
			currentLine = ""
			for _, str := range data[1:] {
				currentLine += str + " "
			}
		}
		currentLine = strings.TrimSpace(currentLine)

		if l.R.RegexFunctionProtoAlfabetico.MatchFuncAlfabetico(currentLine) {
			l.GL.Printf("%+v[FUNC PROTO] Alfabetico Funcion Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[FUNC PROTO] Alfabetico Funcion Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexFunctionProtoEntero.MatchFuncEntero(currentLine) {
			l.GL.Printf("%+v[FUNC PROTO] Entero Funcion Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[FUNC PROTO] Entero Funcion Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexFunctionProtoReal.MatchFuncReal(currentLine) {
			l.GL.Printf("%+v[FUNC PROTO] Real Funcion Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[FUNC PROTO] Real Funcion Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexFunctionProtoLogico.MatchFuncLogico(currentLine) {
			l.GL.Printf("%+v[FUNC PROTO] Logico Funcion Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[FUNC PROTO] Logico Funcion Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexFunctionProtoFlotante.MatchFuncFlotante(currentLine) {
			l.GL.Printf("%+v[FUNC PROTO] Flotante Funcion Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[FUNC PROTO] Flotante Funcion Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexFunctionProtoReal.MatchFuncReal(currentLine) {
			l.GL.Printf("%+v[FUNC PROTO] Real Funcion Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[FUNC PROTO] Real Funcion Found > %+v", currentLine)
			}
			return
		}

		if l.R.RegexFunctionProtoDefault.MatchFuncDefault(currentLine) {
			l.GL.Printf("%+v[FUNC PROTO] Default Funcion Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[FUNC PROTO] Default Funcion Found > %+v", currentLine)
			}
			return
		}

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
