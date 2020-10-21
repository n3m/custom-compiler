package regexprocedureprotoentero

import (
	"regexp"
)

//RegexProcedureProtoEntero ...
type RegexProcedureProtoEntero struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexProcedureProtoEntero ...
func NewRegexProcedureProtoEntero() (*RegexProcedureProtoEntero, error) {
	// var moduleName string = "[regexint][NewRegexProcedureProtoEntero()]"

	return &RegexProcedureProtoEntero{
		V1:      regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*)(\()(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*)((\s*),(\s*)([a-zA-Z]+[a-zA-Z0-9]*))*:(\s*)(?i)Entero(\s*)\)(\s*);$`),
		V2i:     regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*)(\()(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*)((\s*),(\s*)([a-zA-Z]+[a-zA-Z0-9]*))*:(\s*)(?i)Entero(\s*)\)(\s*);$`),
		Keyword: "Entero",
	}, nil
}

//MatchProcedureEntero ...
func (r *RegexProcedureProtoEntero) MatchProcedureEntero(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchProcedureEnteroCaseless ...
func (r *RegexProcedureProtoEntero) MatchProcedureEnteroCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
