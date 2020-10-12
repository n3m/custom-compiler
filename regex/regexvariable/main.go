package regexvariable

import "regexp"

//RegexVariable ...
type RegexVariable struct {
	V1 *regexp.Regexp
	V2 *regexp.Regexp
	V3 *regexp.Regexp
}

//NewRegexVariable ...
func NewRegexVariable() (*RegexVariable, error) {
	compiledV1 := regexp.MustCompile("^variables")
	compiledV2 := regexp.MustCompile("^varia")
	compiledV3 := regexp.MustCompile("^var")
	return &RegexVariable{
		V1: compiledV1,
		V2: compiledV2,
		V3: compiledV3,
	}, nil
}

//StartsWithVariable ...
func (r *RegexVariable) StartsWithVariable(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	if r.V2.MatchString(str) {
		return true
	}

	if r.V3.MatchString(str) {
		return true
	}

	return false

}
