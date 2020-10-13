package regexvariable

import (
	"log"
	"regexp"
	"strings"
)

//RegexVariable ...
type RegexVariable struct {
	keyword string
	V1      *regexp.Regexp
	V2      *regexp.Regexp
	V3      *regexp.Regexp
}

//NewRegexVariable ...
func NewRegexVariable() (*RegexVariable, error) {
	compiledV1 := regexp.MustCompile("^variables")
	compiledV2 := regexp.MustCompile("^varia")
	compiledV3 := regexp.MustCompile("^var")
	return &RegexVariable{
		keyword: "variables",
		V1:      compiledV1,
		V2:      compiledV2,
		V3:      compiledV3,
	}, nil
}

//StartsWithVariable ...
func (r *RegexVariable) StartsWithVariable(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	if r.V2.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		keyword := strings.Split(r.keyword, "")
		foundTypo := false
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != keyword[i] {
					foundTypo = true
					log.Printf("Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.keyword)
				}
			}
		}
		return true
	}

	if r.V3.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		keyword := strings.Split(r.keyword, "")
		foundTypo := false
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != keyword[i] {
					foundTypo = true
					log.Printf("Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.keyword)
				}
			}
		}
		return true
	}

	return false
}

//StartsWithVariableNoCheck ...
func (r *RegexVariable) StartsWithVariableNoCheck(str string) bool {
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
