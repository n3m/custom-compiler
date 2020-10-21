package helpers

import "regexp"

const (
	//ARRAYREGEXP ...
	ARRAYREGEXP string = `(?m)(\w+)([\w*])*`
)

// GetGroupMatches ...
func GetGroupMatches(line, pattern string) []string {
	groups := []string{}

	match, _ := regexp.MatchString(pattern, line)
	if match {

		expression := regexp.MustCompile(pattern)

		matched := expression.FindAllStringSubmatch(line, -1)

		for _, m := range matched {
			groups = append(groups, m[1])
		}
	}
	return groups
}
