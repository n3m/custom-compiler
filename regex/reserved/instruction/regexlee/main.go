package regexlee

import (
	"regexp"
)

//RegexLee ...
type RegexLee struct {
	V1 *regexp.Regexp
}

//NewRegexLee ...
func NewRegexLee() (*RegexLee, error) {
	// var moduleName string = "[regexLee][NewRegexLee()]"

	return &RegexLee{
		V1: regexp.MustCompile(`(?m)[lL]ee\((.*)\)`),
	}, nil
}

//MatchLee ...
func (r *RegexLee) MatchLee(str string) bool {
	return r.V1.MatchString(str)
}

//GroupsLee ...
func (r *RegexLee) GroupsLee(str string) []string {
	groups := []string{}

	if !r.MatchLee(str) {
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
