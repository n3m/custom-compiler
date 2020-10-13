package regexvarlogico

import (
	"regexp"
)

//RegexVarLogico ...
type RegexVarLogico struct {
	V1 *regexp.Regexp
}

//NewRegexVariableLogico ...
func NewRegexVariableLogico() (*RegexVarLogico, error) {
	// var moduleName string = "[regexint][NewRegexVariableLogico()]"

	return &RegexVarLogico{
		V1: regexp.MustCompile(`//^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:Logico;$`),
	}, nil
}

//MatchVariableLogico ...
func (r *RegexVarLogico) MatchVariableLogico(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
