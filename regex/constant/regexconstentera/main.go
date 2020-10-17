package regexconstentera

import (
	"regexp"
)

//RegexEntera ...
type RegexEntera struct {
	V1 *regexp.Regexp
}

//NewRegexEntera ...
func NewRegexEntera() (*RegexEntera, error) {
	// var moduleName string = "[regexEntera][NewRegexEntera()]"

	return &RegexEntera{
		V1: regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*):=(\s*)([0-9]+|\-[0-9]+);$`),
	}, nil
}

//MatchEnteraConstantDeclaration ...
func (r *RegexEntera) MatchEnteraConstantDeclaration(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
