package regexconstante

import (
	"fmt"
	"log"
	"regexp"
)

//RegexConstante ...
type RegexConstante struct {
	V1 *regexp.Regexp
	V2 *regexp.Regexp
	V3 *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
}

//NewRegexConstante ...
func NewRegexConstante(EL *log.Logger, LL *log.Logger) (*RegexConstante, error) {
	if EL == nil || LL == nil {
		return nil, fmt.Errorf("EL or LL loggers came empty")
	}
	return &RegexConstante{
		V1: regexp.MustCompile("^constantes"),
		V2: regexp.MustCompile("^const"),
		V3: regexp.MustCompile("^co"),
	}, nil
}

//StartsWithConstante ...
func (r *RegexConstante) StartsWithConstante(str string) bool {
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
