package regexconstreal

import (
	"regexp"
)

//RegexConstReal ...
type RegexConstReal struct {
	V1 *regexp.Regexp
	V2 *regexp.Regexp
}

//NewRegexConstReal ...
func NewRegexConstReal() (*RegexConstReal, error) {
	// var moduleName string = "[regexConstReal][NewRegexConstReal()]"

	return &RegexConstReal{
		V1: regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*):=(\s*)(([0-9]+|\-[0-9]+)\.([0-9]+)|([0-9]+|\-[0-9]+)e[0-9]+);$`),
		V2: regexp.MustCompile(`(([0-9]+|\-[0-9]+)\.([0-9]+)|([0-9]+|\-[0-9]+)e[0-9]+)`),
	}, nil
}

//MatchRealConstantDeclaration ...
func (r *RegexConstReal) MatchRealConstantDeclaration(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}

//MatchRealConstant ...
func (r *RegexConstReal) MatchRealConstant(str string) bool {
	if r.V2.MatchString(str) {
		return true
	}

	return false
}
