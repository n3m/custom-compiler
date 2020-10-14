package regexfunctionprotodefault

import (
	"regexp"
)

//RegexFuncProtoDefault ...
type RegexFuncProtoDefault struct {
	V1 *regexp.Regexp
}

//NewRegexFuncProtoDefault ...
func NewRegexFuncProtoDefault() (*RegexFuncProtoDefault, error) {
	// var moduleName string = "[regexint][NewRegexFuncProtoDefault()]"

	return &RegexFuncProtoDefault{
		V1: regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):[a-zA-Z]+;$`),
	}, nil
}

//MatchVarDefault ...
func (r *RegexFuncProtoDefault) MatchVarDefault(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}
