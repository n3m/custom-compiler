package helpers

import (
	"regexp"
)

const (
	//ARRAYREGEXP ...
	ARRAYREGEXP string = `(?m)(\w+)([\w*])*`
	//TIPOREGEXP ...
	TIPOREGEXP string = `(?m)([aA]lfabetico|[eE]ntero|[lL]ogico|[rR]eal)`
	//PROCEDIMIENTOREGEXP ...
	PROCEDIMIENTOREGEXP string = `(?m)([pP]rocedimiento) ([a-zA-z]\w*)\((.*)\)`
	//SIREGEXP ...
	SIREGEXP string = `(?m)Si\s?\((.+)\)\s?hacer`
	//HASTAQUEREGEXP ...
	HASTAQUEREGEXP string = `(?m)hasta\sque\s\((.*)\);`
)

var (
	//VARIABLESREGEXP ...
	VARIABLESREGEXP string = ARRAYREGEXP + `(,.*)*:` + TIPOREGEXP[4:]
	//FUNCIONREGEXP ...
	FUNCIONREGEXP string = `(?m)([fF]uncion) ([a-zA-z]\w*)\((.*)\):` + TIPOREGEXP[4:]
)

// GetGroupMatches returns the groups matched by the pattern giving in the line
func GetGroupMatches(line, pattern string) []string {
	groups := []string{}

	match, _ := regexp.MatchString(pattern, line)
	if match {
		expression := regexp.MustCompile(pattern)
		matched := expression.FindAllStringSubmatch(line, -1)

		for _, m := range matched {
			for _, group := range m[1:] {
				if group != "" {
					groups = append(groups, group)
				}
			}
		}
	}
	return groups
}
