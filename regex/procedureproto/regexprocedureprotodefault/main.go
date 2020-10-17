package regexprocedureprotodefault

import (
	"regexp"
)

//RegexProcedureProtoDefault ...
type RegexProcedureProtoDefault struct {
	V1 *regexp.Regexp
}

//NewRegexProcedureProtoDefault ...
func NewRegexProcedureProtoDefault() (*RegexProcedureProtoDefault, error) {
	// var moduleName string = "[regexint][NewRegexProcedureProtoDefault()]"

	return &RegexProcedureProtoDefault{
		V1: regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):[a-zA-Z]+;$`),
	}, nil
}

//MatchProcedureDefault ...
func (r *RegexProcedureProtoDefault) MatchProcedureDefault(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}
