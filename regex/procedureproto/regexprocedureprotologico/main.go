package regexprocedureprotologico

import (
	"regexp"
)

//RegexProcedureProtoLogico ...
type RegexProcedureProtoLogico struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexProcedureProtoLogico ...
func NewRegexProcedureProtoLogico() (*RegexProcedureProtoLogico, error) {
	// var moduleName string = "[regexint][NewRegexProcedureProtoLogico()]"

	return &RegexProcedureProtoLogico{
		V1:      regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):Logico;$`),
		V2i:     regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):(?i)Logico;$`),
		Keyword: "Logico",
	}, nil
}

//MatchProcedureLogico ...
func (r *RegexProcedureProtoLogico) MatchProcedureLogico(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchProcedureLogicoCaseless ...
func (r *RegexProcedureProtoLogico) MatchProcedureLogicoCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
