package regexprocedure

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexProcedure ...
type RegexProcedure struct {
	Keyword string
	V1      *regexp.Regexp
	V2      *regexp.Regexp
	V3      *regexp.Regexp
	V4End   *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexProcedure ...
func NewRegexProcedure(EL, LL, GL *log.Logger) (*RegexProcedure, error) {
	var moduleName string = "[regexProcedure][NewRegexProcedure()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}

	compiledV1 := regexp.MustCompile("^Procedimiento")
	compiledV2 := regexp.MustCompile("^(?i)Procedim")
	compiledV3 := regexp.MustCompile("^(?i)Proc")
	compiledV4End := regexp.MustCompile("[^;]$")
	return &RegexProcedure{
		Keyword: "Procedimiento",
		V1:      compiledV1,
		V2:      compiledV2,
		V3:      compiledV3,
		V4End:   compiledV4End,
		EL:      EL,
		LL:      LL,
		GL:      GL,
	}, nil
}

//StartsWithProcedure ...
func (r *RegexProcedure) StartsWithProcedure(str string) bool {
	if r.V1.MatchString(str) && r.V4End.MatchString(str) {
		return true
	}

	if r.V2.MatchString(str) && r.V4End.MatchString(str) {
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

	if r.V3.MatchString(str) && r.V4End.MatchString(str) {
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

//StartsWithProcedureNoCheck ...
func (r *RegexProcedure) StartsWithProcedureNoCheck(str string) bool {
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
