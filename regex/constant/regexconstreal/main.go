package regexconstreal

import (
	"regexp"
)

//RegexReal ...
type RegexReal struct {
	V1 *regexp.Regexp
}

//NewRegexReal ...
func NewRegexReal() (*RegexReal, error) {
	// var moduleName string = "[regexReal][NewRegexReal()]"

	return &RegexReal{
		V1: regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*):=(\s*)(([0-9]+|\-[0-9]+)\.([0-9]+)|([0-9]+|\-[0-9]+)e[0-9]+);$`),
	}, nil
}

//MatchRealConstantDeclaration ...
func (r *RegexReal) MatchRealConstantDeclaration(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
