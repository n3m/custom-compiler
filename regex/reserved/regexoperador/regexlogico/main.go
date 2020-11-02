package regexoplogico

import (
	"regexp"
)

//RegexOpLogico ...
type RegexOpLogico struct {
	V1 *regexp.Regexp
}

//NewRegexOpLogico ...
func NewRegexOpLogico() (*RegexOpLogico, error) {
	// var moduleName string = "[regexOpLogico][NewRegexOpLogico()]"

	return &RegexOpLogico{
		V1: regexp.MustCompile(`(?m)\s(y|o|no)\s`),
	}, nil
}

//MatchOpLogico ...
func (r *RegexOpLogico) MatchOpLogico(str string) bool {
	return r.V1.MatchString(str)

}

//GroupsOpLogico ...
func (r *RegexOpLogico) GroupsOpLogico(str string) []string {
	groups := []string{}

	if !r.MatchOpLogico(str) {
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
