package regeximprime

import (
	"regexp"
)

//RegexImprime ...
type RegexImprime struct {
	V1 *regexp.Regexp
}

//NewRegexImprime ...
func NewRegexImprime() (*RegexImprime, error) {
	// var moduleName string = "[regexImprime][NewRegexImprime()]"

	return &RegexImprime{
		V1: regexp.MustCompile(`(?m)[iI]mprime(nl)?\((.*)\)`),
	}, nil
}

//MatchImprime ...
func (r *RegexImprime) MatchImprime(str string) bool {
	return r.V1.MatchString(str)
}

//GroupsImprime ...
func (r *RegexImprime) GroupsImprime(str string) []string {
	groups := []string{}

	if !r.MatchImprime(str) {
		return groups
	}

	matched := r.V1.FindAllStringSubmatch(str, -1)
	for _, m := range matched {
		for _, group := range m[1:] {
			if group != "" {
				groups = append(groups, group)
			}
		}
	}

	return groups
}
