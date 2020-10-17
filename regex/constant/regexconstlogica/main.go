package regexconstlogica

import (
	"regexp"
)

//RegexLogica ...
type RegexLogica struct {
	V1 *regexp.Regexp
}

//NewRegexLogica ...
func NewRegexLogica() (*RegexLogica, error) {
	// var moduleName string = "[regexLogica][NewRegexLogica()]"

	return &RegexLogica{
		V1: regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*):=(\s*)(verdadero|falso);$`),
	}, nil
}

//MatchLogicaConstantDeclaration ...
func (r *RegexLogica) MatchLogicaConstantDeclaration(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
