package regexconstalfabetica

import (
	"regexp"
)

//RegexConstAlfabetica ...
type RegexConstAlfabetica struct {
	V1 *regexp.Regexp
}

//NewRegexConstAlfabetica ...
func NewRegexConstAlfabetica() (*RegexConstAlfabetica, error) {
	// var moduleName string = "[regexConstAlfabetica][NewRegexConstAlfabetica()]"

	return &RegexConstAlfabetica{
		V1: regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*):=(\s*)(\"(\w)*\");$`),
	}, nil
}

//MatchAlfabeticaConstantDeclaration ...
func (r *RegexConstAlfabetica) MatchAlfabeticaConstantDeclaration(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
