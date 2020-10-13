package regexvarflotante

import (
	"regexp"
)

//RegexVarFlotante ...
type RegexVarFlotante struct {
	V1 *regexp.Regexp
}

//NewRegexVariableFlotante ...
func NewRegexVariableFlotante() (*RegexVarFlotante, error) {
	// var moduleName string = "[regexint][NewRegexVariableFlotante()]"

	return &RegexVarFlotante{
		V1: regexp.MustCompile(`^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:Flotante;$`),
	}, nil
}

//MatchVariableFlotante ...
func (r *RegexVarFlotante) MatchVariableFlotante(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
