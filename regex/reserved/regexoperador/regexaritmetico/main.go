package regexoparitmetico

import (
	"regexp"
)

//RegexOpAritmetico ...
type RegexOpAritmetico struct {
	V1 *regexp.Regexp
}

//NewRegexOpAritmetico ...
func NewRegexOpAritmetico() (*RegexOpAritmetico, error) {
	// var moduleName string = "[regexOpAritmetico][NewRegexOpAritmetico()]"

	return &RegexOpAritmetico{
		V1: regexp.MustCompile(`(?m)(\+|\-|\*|\/|\%|\^)`),
	}, nil
}

//MatchOpAritmetico ...
func (r *RegexOpAritmetico) MatchOpAritmetico(str string) bool {
	return r.V1.MatchString(str)

}

//GroupsOpAritmetico ...
func (r *RegexOpAritmetico) GroupsOpAritmetico(str string) []string {
	groups := []string{}

	if !r.MatchOpAritmetico(str) {
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
