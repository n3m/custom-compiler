package lexyc

import (
	"go-custom-compiler/models"
	"log"
	"strings"
)

//NextProcedureProto ...
func (l *LexicalAnalyzer) NextProcedureProto(currentLine string, lineIndex int64, debug bool) {
	funcName := "[NextProcedureProto()] "
	// var moduleName string = "[regexfunctionproto][NextProcedureProto()]"

	if l.CurrentBlockType == models.PROCEDUREPROTOBLOCK {
		if l.R.RegexProcedureProto.StartsWithProcedureProtoNoCheck(currentLine) {
			data := strings.Split(currentLine, " ")
			currentLine = ""
			for _, str := range data[1:] {
				currentLine += str + " "
			}
		}
		currentLine = strings.TrimSpace(currentLine)

		if l.R.RegexProcedureProtoEntero.MatchProcedureEntero(currentLine) {
			l.GL.Printf("%+v[PROCEDURE PROTO] Entero Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[PROCEDURE PROTO] Entero Found > %+v", currentLine)
			}
			return
		}
		if l.R.RegexProcedureProtoLogico.MatchProcedureLogico(currentLine) {

			l.GL.Printf("%+v[PROCEDURE PROTO] Logico Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[PROCEDURE PROTO] Logico Found > %+v", currentLine)
			}
			return
		}
		if l.R.RegexProcedureProtoReal.MatchProcedureReal(currentLine) {

			l.GL.Printf("%+v[PROCEDURE PROTO] Real Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[PROCEDURE PROTO] Real Found > %+v", currentLine)
			}
			return
		}
		if l.R.RegexProcedureProtoAlfabetico.MatchProcedureAlfabetico(currentLine) {

			l.GL.Printf("%+v[PROCEDURE PROTO] Alfabetico Found > %+v", funcName, currentLine)
			if debug {
				log.Printf("[PROCEDURE PROTO] Alfabetico Found > %+v", currentLine)
			}
			return
		}

		l.GL.Printf("%+v Did not found any type of match on Line[%+v]! ", funcName, lineIndex)

	}
}
