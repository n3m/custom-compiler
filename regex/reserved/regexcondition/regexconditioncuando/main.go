package regexconditioncuando

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexConditionCuando ...
type RegexConditionCuando struct {
	Keyword string
	V1      *regexp.Regexp
	V2      *regexp.Regexp
	V3      *regexp.Regexp
	V4      *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexConditionCuando ...
func NewRegexConditionCuando(EL, LL, GL *log.Logger) (*RegexConditionCuando, error) {
	var moduleName string = "[RegexConditionCuando][NewRegexConditionCuando()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}

	compiledE1 := regexp.MustCompile(`^(\s*)(Cuando)(\s+)`)
	compiledE2 := regexp.MustCompile(`^(\s*)((?i)Cuando)(\s+)`)
	compiledE3 := regexp.MustCompile(`^(\s*)((?i)Cuan)(\s+)`)
	compiledE4 := regexp.MustCompile(`^(\s*)((?i)Cu)(\s+)`)
	return &RegexConditionCuando{
		Keyword: "Cuando",
		V1:      compiledE1,
		V2:      compiledE2,
		V3:      compiledE3,
		V4:      compiledE4,
		EL:      EL,
		LL:      LL,
		GL:      GL,
	}, nil
}

//StartsWithCuando ...
func (r *RegexConditionCuando) StartsWithCuando(str string, lineIndex int64) bool {

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
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)
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
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)
				}
			}
		}
		return true
	}

	if r.V4.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		Keyword := strings.Split(r.Keyword, "")
		foundTypo := false
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)
				}
			}
		}
		return true
	}

	return false
}

//StartsWithCuandoNoCheck ...
func (r *RegexConditionCuando) StartsWithCuandoNoCheck(str string) bool {
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

//r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)

//LogError ...
//"# Linea | # Columna | Error | Descripcion | Linea del Error"
func (r *RegexConditionCuando) LogError(lineIndex int64, columnIndex interface{}, err string, description string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, description, currentLine)
}
