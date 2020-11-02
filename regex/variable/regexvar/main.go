package regexconstentera

import (
	"regexp"
)

//RegexVar ...
type RegexVar struct {
	V1 *regexp.Regexp
}

//NewRegexVar ...
func NewRegexVar() (*RegexVar, error) {
	// var moduleName string = "[regexVar][NewRegexVar()]"

	return &RegexVar{
		V1: regexp.MustCompile(`(?m)(\w+)(\[\w*\])?(\[\w*\])?`),
	}, nil
}

//MatchVar ...
func (r *RegexVar) MatchVar(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false
}

//GroupsVar ...
func (r *RegexVar) GroupsVar(str string) []string {
	groups := []string{}

	if !r.MatchVar(str) {
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
