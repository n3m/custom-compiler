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

		if l.R.RegexFunctionProtoDefault.MatchFuncDefault(currentLine) {
			funcType, _, _, _ := getDataFromFunctionProto(currentLine)

			if l.R.RegexFunctionProtoAlfabetico.MatchFuncAlfabeticoCaseless(currentLine) {
				l.GL.Printf("%+v[FUNC PROTO] Alfabetico Funcion Found > %+v", funcName, currentLine)
				if debug {
					log.Printf("[FUNC PROTO] Alfabetico Funcion Found > %+v", currentLine)
				}
				foundTypo := false
				keyData := strings.Split(l.R.RegexFunctionProtoAlfabetico.Keyword, "")
				for i, char := range funcType {
					if i < len(keyData)-1 {
						if !foundTypo {
							if string(char) != keyData[i] {
								foundTypo = true

								log.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", funcType, i, lineIndex, l.R.RegexFunctionProtoAlfabetico.Keyword)

								l.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", funcType, i, lineIndex, l.R.RegexFunctionProtoAlfabetico.Keyword)
							}
						}
					}
				}
				return
			}

			if l.R.RegexFunctionProtoEntero.MatchFuncEnteroCaseless(currentLine) {
				l.GL.Printf("%+v[FUNC PROTO] Entero Funcion Found > %+v", funcName, currentLine)
				if debug {
					log.Printf("[FUNC PROTO] Entero Funcion Found > %+v", currentLine)
				}
				foundTypo := false
				keyData := strings.Split(l.R.RegexFunctionProtoEntero.Keyword, "")
				for i, char := range funcType {
					if i < len(keyData)-1 {
						if !foundTypo {
							if string(char) != keyData[i] {
								foundTypo = true

								log.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", funcType, i, lineIndex, l.R.RegexFunctionProtoEntero.Keyword)

								l.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", funcType, i, lineIndex, l.R.RegexFunctionProtoEntero.Keyword)
							}
						}
					}
				}
				return
			}

			if l.R.RegexFunctionProtoReal.MatchFuncRealCaseless(currentLine) {
				l.GL.Printf("%+v[FUNC PROTO] Real Funcion Found > %+v", funcName, currentLine)
				if debug {
					log.Printf("[FUNC PROTO] Real Funcion Found > %+v", currentLine)
				}
				foundTypo := false
				keyData := strings.Split(l.R.RegexFunctionProtoReal.Keyword, "")
				for i, char := range funcType {
					if i < len(keyData)-1 {
						if !foundTypo {
							if string(char) != keyData[i] {
								foundTypo = true
								log.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", funcType, i, lineIndex, l.R.RegexFunctionProtoReal.Keyword)
								l.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", funcType, i, lineIndex, l.R.RegexFunctionProtoReal.Keyword)
							}
						}
					}
				}
				return
			}

			if l.R.RegexFunctionProtoLogico.MatchFuncLogicoCaseless(currentLine) {
				l.GL.Printf("%+v[FUNC PROTO] Logico Funcion Found > %+v", funcName, currentLine)
				if debug {
					log.Printf("[FUNC PROTO] Logico Funcion Found > %+v", currentLine)
				}
				foundTypo := false
				keyData := strings.Split(l.R.RegexFunctionProtoLogico.Keyword, "")
				for i, char := range funcType {
					if i < len(keyData)-1 {
						if !foundTypo {
							if string(char) != keyData[i] {
								foundTypo = true

								log.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", funcType, i, lineIndex, l.R.RegexFunctionProtoLogico.Keyword)

								l.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v] on line [%+v]. Correct syntax should be '%+v'", funcType, i, lineIndex, l.R.RegexFunctionProtoLogico.Keyword)
							}
						}
					}
				}
				return
			}

			return
		}

		l.GL.Printf("%+v Did not found any type of match on Line[%+v]! ", funcName, lineIndex)

	}
}

func getDataFromFunctionProto(currentLine string) (string, string, string, string) {
	currentLine = strings.TrimSuffix(currentLine, ";")
	currentLine = strings.TrimSuffix(currentLine, " ")
	lineData := strings.Split(currentLine, "):")
	funcType := lineData[1]

	funcData := lineData[0]
	funcDataV2 := strings.Split(funcData, "(")
	funcName := funcDataV2[0]
	funcParams := funcDataV2[1]
	paramsData := strings.Split(funcParams, ":")
	funcParamType := paramsData[1]
	funcParamName := paramsData[0]

	return funcType, funcName, funcParamType, funcParamName
}
