package regexprocedureprotoflotante

import (
	"regexp"
)

//RegexProcedureProtoFlotante ...
type RegexProcedureProtoFlotante struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexProcedureProtoFlotante ...
func NewRegexProcedureProtoFlotante() (*RegexProcedureProtoFlotante, error) {
	// var moduleName string = "[regexint][NewRegexProcedureProtoFlotante()]"

	return &RegexProcedureProtoFlotante{
		V1:      regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):Flotante;$`),
		V2i:     regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):(?i)Flotante;$`),
		Keyword: "Flotante",
	}, nil
}

//MatchProcedureFlotante ...
func (r *RegexProcedureProtoFlotante) MatchProcedureFlotante(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchProcedureFlotanteCaseless ...
func (r *RegexProcedureProtoFlotante) MatchProcedureFlotanteCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
