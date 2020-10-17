package regexconstdefault

import (
	"regexp"
)

//RegexConstDefault ...
type RegexConstDefault struct {
	V1 *regexp.Regexp
}

//NewRegexVariableDefault ...
func NewRegexVariableDefault() (*RegexConstDefault, error) {
	// var moduleName string = "[regexint][NewRegexVariableDefault()]"

	return &RegexConstDefault{
		V1: regexp.MustCompile(`^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:[a-zA-Z]+;$`),
	}, nil
}

//MatchVariableDefault ...
func (r *RegexConstDefault) MatchVariableDefault(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}
