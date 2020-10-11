package variable

import "regexp"

//RegexVariable ...
type RegexVariable struct {
}

//NewRegexVariable ...
func NewRegexVariable() (*RegexVariable, error) {
	return &RegexVariable{}, nil
}

//StartsWithVariable ...
func (r *RegexVariable) StartsWithVariable(str string) bool {
	compiledV1 := regexp.MustCompile("^variables")
	if compiledV1.MatchString(str) {
		return true
	}

	compiledV2 := regexp.MustCompile("^varia")
	if compiledV2.MatchString(str) {
		return true
	}

	compiledV3 := regexp.MustCompile("^var")
	if compiledV3.MatchString(str) {
		return true
	}

	return false

}
