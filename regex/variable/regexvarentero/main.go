package regexvarentero

import (
	"regexp"
)

//RegexVarEntero ...
type RegexVarEntero struct {
	V1 *regexp.Regexp
}

//NewRegexVariableEntero ...
func NewRegexVariableEntero() (*RegexVarEntero, error) {
	// var moduleName string = "[regexint][NewRegexVariableEntero()]"

	return &RegexVarEntero{
		V1: regexp.MustCompile(`^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:Entero;$`),
	}, nil
}

//MatchVariableEntero ...
func (r *RegexVarEntero) MatchVariableEntero(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
