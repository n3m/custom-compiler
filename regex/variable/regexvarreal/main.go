package regexvarreal

import (
	"regexp"
)

//RegexVarReal ...
type RegexVarReal struct {
	V1 *regexp.Regexp
}

//NewRegexVariableReal ...
func NewRegexVariableReal() (*RegexVarReal, error) {
	// var moduleName string = "[regexint][NewRegexVariableReal()]"

	return &RegexVarReal{
		V1: regexp.MustCompile(`^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:Real;$`),
	}, nil
}

//MatchVariableReal ...
func (r *RegexVarReal) MatchVariableReal(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
