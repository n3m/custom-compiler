package regexvaralfabetico

import (
	"regexp"
)

//RegexAlfabetico ...
type RegexAlfabetico struct {
	V1 *regexp.Regexp
}

//NewRegexVariableAlfabetico ...
func NewRegexVariableAlfabetico() (*RegexAlfabetico, error) {
	// var moduleName string = "[regexint][NewRegexVariableAlfabetico()]"

	return &RegexAlfabetico{
		V1: regexp.MustCompile(`//^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:Alfabetico;$`),
	}, nil
}

//MatchVariableAlfabetico ...
func (r *RegexAlfabetico) MatchVariableAlfabetico(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
