package regexconstfloat

import (
	"regexp"
)

//RegexFloat ...
type RegexFloat struct {
	V1 *regexp.Regexp
}

//NewRegexFloat ...
func NewRegexFloat() (*RegexFloat, error) {
	// var moduleName string = "[regexfloat][NewRegexFloat()]"

	return &RegexFloat{
		V1: regexp.MustCompile(`^\w[^\s]*\s*:=\s*\d+\.\d+\s*(\s*\+\s*\d+\.\d+)*`),
	}, nil
}

//MatchFloatConstantDeclaration ...
func (r *RegexFloat) MatchFloatConstantDeclaration(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
