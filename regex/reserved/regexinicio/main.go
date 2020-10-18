package regexinicio

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexInicio ...
type RegexInicio struct {
	Keyword string
	V1      *regexp.Regexp
	V2      *regexp.Regexp
	V3      *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexInicio ...
func NewRegexInicio(EL, LL, GL *log.Logger) (*RegexInicio, error) {
	var moduleName string = "[regexInicio][NewRegexInicio()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}
	return &RegexInicio{
		Keyword: "Inicio",
		V1:      regexp.MustCompile("^Inicio"),
		V2:      regexp.MustCompile("^(?i)Inic"),
		V3:      regexp.MustCompile("^(?i)In"),
		GL:      GL,
		EL:      EL,
		LL:      LL,
	}, nil
}

//StartsWithInicio ...
func (r *RegexInicio) StartsWithInicio(str string) bool {

	if r.V1.MatchString(str) {
		return true
	}

	if r.V2.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		Keyword := strings.Split(r.Keyword, "")

		foundTypo := false
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true

					log.Printf("[ERR] Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.Keyword)
					r.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.Keyword)
				}
			}
		}
		return true
	}

	if r.V3.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		Keyword := strings.Split(r.Keyword, "")
		foundTypo := false
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					log.Printf("[ERR] Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.Keyword)
					r.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.Keyword)
				}
			}
		}
		return true
	}

	return false
}

//StartsWithInicioNoCheck ...
func (r *RegexInicio) StartsWithInicioNoCheck(str string) bool {
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
