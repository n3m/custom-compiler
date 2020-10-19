package regexvarreal

import (
	"regexp"
)

//RegexVarReal ...
type RegexVarReal struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexVariableReal ...
func NewRegexVariableReal() (*RegexVarReal, error) {
	// var moduleName string = "[regexint][NewRegexVariableReal()]"

	return &RegexVarReal{
		V1:      regexp.MustCompile(`^[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*(\[[a-zA-Z0-9]+[a-zA-Z0-9]*\])*)*:(?i)Real;$`),
		V2i:     regexp.MustCompile(`(?i)real`),
		Keyword: "Real",
	}, nil
}

//MatchVariableReal ...
func (r *RegexVarReal) MatchVariableReal(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchVariableRealCaseless ...
func (r *RegexVarReal) MatchVariableRealCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
