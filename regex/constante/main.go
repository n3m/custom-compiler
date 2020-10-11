package constante

import "regexp"

//RegexConstante ...
type RegexConstante struct {
}

//NewRegexConstante ...
func NewRegexConstante() (*RegexConstante, error) {
	return &RegexConstante{}, nil
}

//StartsWithConstante ...
func (r *RegexConstante) StartsWithConstante(str string) bool {
	compiledV1 := regexp.MustCompile("^constantes")
	if compiledV1.MatchString(str) {
		return true
	}

	compiledV2 := regexp.MustCompile("^const")
	if compiledV2.MatchString(str) {
		return true
	}

	compiledV3 := regexp.MustCompile("^co")
	if compiledV3.MatchString(str) {
		return true
	}

	return false

}
