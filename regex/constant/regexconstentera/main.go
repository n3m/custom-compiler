package regexconstentera

import (
	"regexp"
)

//RegexConstEntera ...
type RegexConstEntera struct {
	V1 *regexp.Regexp
	V2 *regexp.Regexp
}

//NewRegexConstEntera ...
func NewRegexConstEntera() (*RegexConstEntera, error) {
	// var moduleName string = "[regexConstEntera][NewRegexConstEntera()]"

	return &RegexConstEntera{
		V1: regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*):=(\s*)([0-9]+|\-[0-9]+)(;)?$`),
		V2: regexp.MustCompile(`([0-9]+|\-[0-9]+)`),
	}, nil
}

//MatchEnteraConstantDeclaration ...
func (r *RegexConstEntera) MatchEnteraConstantDeclaration(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}

//MatchEnteraConstant ...
func (r *RegexConstEntera) MatchEnteraConstant(str string) bool {
	if r.V2.MatchString(str) {
		return true
	}

	return false

}
