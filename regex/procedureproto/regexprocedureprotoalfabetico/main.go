package regexprocedureprotoalfabetico

import (
	"regexp"
)

//RegexProcedureProtoAlfabetico ...
type RegexProcedureProtoAlfabetico struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexProcedureProtoAlfabetico ...
func NewRegexProcedureProtoAlfabetico() (*RegexProcedureProtoAlfabetico, error) {
	// var moduleName string = "[regexint][NewRegexProcedureProtoAlfabetico()]"

	return &RegexProcedureProtoAlfabetico{
		V1:      regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):Alfabetico;$`),
		V2i:     regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):(?i)alfabetico;$`),
		Keyword: "Alfabetico",
	}, nil
}

//MatchProcedureAlfabetico ...
func (r *RegexProcedureProtoAlfabetico) MatchProcedureAlfabetico(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchProcedureAlfabeticoCaseless ...
func (r *RegexProcedureProtoAlfabetico) MatchProcedureAlfabeticoCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
