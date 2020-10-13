package regexvardefault

import (
	"regexp"
)

//RegexVarDefault ...
type RegexVarDefault struct {
	V1 *regexp.Regexp
}

//NewRegexVariableDefault ...
func NewRegexVariableDefault() (*RegexVarDefault, error) {
	// var moduleName string = "[regexint][NewRegexVariableDefault()]"

	return &RegexVarDefault{
		V1: regexp.MustCompile(`^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:[a-zA-Z]+;$`),
	}, nil
}

//MatchVariableDefault ...
func (r *RegexVarDefault) MatchVariableDefault(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
