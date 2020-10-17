package regexconstalfabetica

import (
	"regexp"
)

//RegexAlfabetica ...
type RegexAlfabetica struct {
	V1 *regexp.Regexp
}

//NewRegexAlfabetica ...
func NewRegexAlfabetica() (*RegexAlfabetica, error) {
	// var moduleName string = "[regexAlfabetica][NewRegexAlfabetica()]"

	return &RegexAlfabetica{
		V1: regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*):=(\s*)(\"(\w)*\");$`),
	}, nil
}

//MatchAlfabeticaConstantDeclaration ...
func (r *RegexAlfabetica) MatchAlfabeticaConstantDeclaration(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
