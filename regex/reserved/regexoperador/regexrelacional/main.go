package regexoprelacional

import (
	"regexp"
)

//RegexOpRelacional ...
type RegexOpRelacional struct {
	V1 *regexp.Regexp
}

//NewRegexOpRelacional ...
func NewRegexOpRelacional() (*RegexOpRelacional, error) {
	// var moduleName string = "[regexOpRelacional][NewRegexOpRelacional()]"

	return &RegexOpRelacional{
		V1: regexp.MustCompile(`(?m)(=|<>|<|>|<=|>=)`),
	}, nil
}

//MatchOpRelacional ...
func (r *RegexOpRelacional) MatchOpRelacional(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}

//GroupsOpRelacional ...
func (r *RegexOpRelacional) GroupsOpRelacional(str string) []string {
	groups := []string{}

	if !r.MatchOpRelacional(str) {
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
