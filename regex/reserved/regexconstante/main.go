package regexconstante

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexConstante ...
type RegexConstante struct {
	keyword string
	V1      *regexp.Regexp
	V2      *regexp.Regexp
	V3      *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexConstante ...
func NewRegexConstante(EL, LL, GL *log.Logger) (*RegexConstante, error) {
	var moduleName string = "[regexconstante][NewRegexConstante()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}
	return &RegexConstante{
		keyword: "constantes",
		V1:      regexp.MustCompile("^constantes"),
		V2:      regexp.MustCompile("^consta"),
		V3:      regexp.MustCompile("^con"),
		GL:      GL,
		EL:      EL,
		LL:      LL,
	}, nil
}

//StartsWithConstante ...
func (r *RegexConstante) StartsWithConstante(str string) bool {

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
					r.GL.Printf("Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.keyword)
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
					r.GL.Printf("Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.keyword)
				}
			}
		}
		return true
	}

	return false
}

//StartsWithConstanteNoCheck ...
func (r *RegexConstante) StartsWithConstanteNoCheck(str string) bool {
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
