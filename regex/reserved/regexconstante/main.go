package regexconstante

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexConstante ...
type RegexConstante struct {
	Keyword string
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
		Keyword: "constantes",
		V1:      regexp.MustCompile("^constantes"),
		V2:      regexp.MustCompile("^(?i)consta"),
		V3:      regexp.MustCompile("^(?i)con"),
		GL:      GL,
		EL:      EL,
		LL:      LL,
	}, nil
}

//StartsWithConstante ...
func (r *RegexConstante) StartsWithConstante(str string, lineIndex int64) bool {

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
					//"# Linea | # Columna | Error | Descripcion | Linea del Error"
					r.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, i, wrongWord, r.Keyword, str)
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
					//"# Linea | # Columna | Error | Descripcion | Linea del Error"
					r.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, i, wrongWord, r.Keyword, str)
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
