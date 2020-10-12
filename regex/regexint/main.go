package regexint

import (
	"regexp"
)

//RegexInt ...
type RegexInt struct {
	V1 *regexp.Regexp
}

//NewRegexInt ...
func NewRegexInt() (*RegexInt, error) {
	return &RegexInt{
		V1: regexp.MustCompile(`^\w[^\s]*\s*:=\s*\d+\s*(\s*\+\s*\d+)*`),
	}, nil
}

//MatchIntConstantDeclaration ...
func (r *RegexInt) MatchIntConstantDeclaration(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}
