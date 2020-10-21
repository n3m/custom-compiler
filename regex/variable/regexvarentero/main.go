package regexvarentero

import (
	"regexp"
)

//RegexVarEntero ...
type RegexVarEntero struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexVariableEntero ...
func NewRegexVariableEntero() (*RegexVarEntero, error) {
	// var moduleName string = "[regexint][NewRegexVariableEntero()]"

	return &RegexVarEntero{
		V1:      regexp.MustCompile(`^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:(?i)Entero;$`),
		V2i:     regexp.MustCompile(`(?i)entero`),
		Keyword: "Entero",
	}, nil
}

//MatchVariableEntero ...
func (r *RegexVarEntero) MatchVariableEntero(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchVariableEnteroCaseless ...
func (r *RegexVarEntero) MatchVariableEnteroCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
