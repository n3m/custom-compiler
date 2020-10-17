package regexprocedureprotoreal

import (
	"regexp"
)

//RegexProcedureProtoReal ...
type RegexProcedureProtoReal struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexProcedureProtoReal ...
func NewRegexProcedureProtoReal() (*RegexProcedureProtoReal, error) {
	// var moduleName string = "[regexint][NewRegexProcedureProtoReal()]"

	return &RegexProcedureProtoReal{
		V1:      regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*)(\()(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*)((\s*),(\s*)([a-zA-Z]+[a-zA-Z0-9]*))*:(\s*)Real(\s*)\)(\s*);$`),
		V2i:     regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*)(\()(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*)((\s*),(\s*)([a-zA-Z]+[a-zA-Z0-9]*))*:(\s*)(?i)Real(\s*)\)(\s*);$`),
		Keyword: "Real",
	}, nil
}

//MatchProcedureReal ...
func (r *RegexProcedureProtoReal) MatchProcedureReal(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchProcedureRealCaseless ...
func (r *RegexProcedureProtoReal) MatchProcedureRealCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
