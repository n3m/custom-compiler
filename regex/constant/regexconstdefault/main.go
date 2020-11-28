package regexconstdefault

import (
	"regexp"
)

//RegexConstDefault ...
type RegexConstDefault struct {
	V1 *regexp.Regexp
}

//NewRegexConstDefault ...
func NewRegexConstDefault() (*RegexConstDefault, error) {
	// var moduleName string = "[regexint][NewRegexConstDefault()]"

	return &RegexConstDefault{
		V1: regexp.MustCompile(`^(\s*)([a-zA-Z]+[a-zA-Z0-9]*)(\s*):=(\s*)((\"(\w)*\")|([0-9]+|\-[0-9]+)|(([0-9]+|\-[0-9]+)\.([0-9]+)|([0-9]+|\-[0-9]+)e[0-9]+)|((?i)verdadero|(?i)falso));$`),
	}, nil
}

//MatchConstantDefault ...
func (r *RegexConstDefault) MatchConstantDefault(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}
