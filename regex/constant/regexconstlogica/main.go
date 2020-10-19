package regexconstlogica

import (
	"regexp"
)

//RegexConstLogica ...
type RegexConstLogica struct {
	V1  *regexp.Regexp
	V2i *regexp.Regexp
}

//NewRegexConstLogica ...
func NewRegexConstLogica() (*RegexConstLogica, error) {
	// var moduleName string = "[regexConstLogica][NewRegexConstLogica()]"

	return &RegexConstLogica{
		V1:  regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*):=(\s*)((?i)verdadero|(?i)falso);$`),
		V2i: regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*):=(\s*)((?i)verdadero|(?i)falso);$`),
	}, nil
}

//MatchLogicaConstantDeclaration ...
func (r *RegexConstLogica) MatchLogicaConstantDeclaration(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
