package regexvarflotante

import (
	"regexp"
)

//RegexVarFlotante ...
type RegexVarFlotante struct {
	keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexVariableFlotante ...
func NewRegexVariableFlotante() (*RegexVarFlotante, error) {
	// var moduleName string = "[regexint][NewRegexVariableFlotante()]"

	return &RegexVarFlotante{
		V1:      regexp.MustCompile(`^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:Flotante;$`),
		V2i:     regexp.MustCompile(`(?i)flotante`),
		keyword: "Flotante",
	}, nil
}

//MatchVariableFlotante ...
func (r *RegexVarFlotante) MatchVariableFlotante(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchVariableFlotanteCaseless ...
func (r *RegexVarFlotante) MatchVariableFlotanteCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
