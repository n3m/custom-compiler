package regexvaralfabetico

import (
	"regexp"
)

//RegexVarAlfabetico ...
type RegexVarAlfabetico struct {
	V1 *regexp.Regexp
}

//NewRegexVariableAlfabetico ...
func NewRegexVariableAlfabetico() (*RegexVarAlfabetico, error) {
	// var moduleName string = "[regexint][NewRegexVariableAlfabetico()]"

	return &RegexVarAlfabetico{
		V1: regexp.MustCompile(`//^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:Alfabetico;$`),
	}, nil
}

//MatchVariableAlfabetico ...
func (r *RegexVarAlfabetico) MatchVariableAlfabetico(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
