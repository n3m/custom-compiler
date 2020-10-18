package regexlooprepetir

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexLoopRepetir ...
type RegexLoopRepetir struct {
	Keyword string
	V1      *regexp.Regexp
	V2      *regexp.Regexp
	V3      *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexLoopRepetir ...
func NewRegexLoopRepetir(EL, LL, GL *log.Logger) (*RegexLoopRepetir, error) {
	var moduleName string = "[RegexLoopRepetir][NewRegexLoopRepetir()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}

	compiledV1 := regexp.MustCompile("^Repetir")
	compiledV2 := regexp.MustCompile("^(?i)Repet")
	compiledV3 := regexp.MustCompile("^(?i)Rep")

	return &RegexLoopRepetir{
		Keyword: "Repetir",
		V1:      compiledV1,
		V2:      compiledV2,
		V3:      compiledV3,
		EL:      EL,
		LL:      LL,
		GL:      GL,
	}, nil
}

//StartsWithRepetir ...
func (r *RegexLoopRepetir) StartsWithRepetir(str string, lineIndex int64) bool {
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

//StartsWithRepetirNoCheck ...
func (r *RegexLoopRepetir) StartsWithRepetirNoCheck(str string) bool {
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
