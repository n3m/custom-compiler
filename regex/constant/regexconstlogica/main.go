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
		V1: regexp.MustCompile(`^\w[^\s]*\s*:=\s*\d+\s*(\s*\+\s*\d+)*;`),
	}, nil
}

//MatchLogicaConstantDeclaration ...
func (r *RegexLogica) MatchLogicaConstantDeclaration(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
