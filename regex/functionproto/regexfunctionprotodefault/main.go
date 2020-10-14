package regexfunctionprotodefault

import (
	"regexp"
)

//RegexFuncProtoDefault ...
type RegexFuncProtoDefault struct {
	V1 *regexp.Regexp
}

//NewRegexFuncProtoiableDefault ...
func NewRegexFuncProtoiableDefault() (*RegexFuncProtoDefault, error) {
	// var moduleName string = "[regexint][NewRegexFuncProtoiableDefault()]"

	return &RegexFuncProtoDefault{
		V1: regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):[a-zA-Z]+;$`),
	}, nil
}

//MatchVariableDefault ...
func (r *RegexFuncProtoDefault) MatchVariableDefault(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}
