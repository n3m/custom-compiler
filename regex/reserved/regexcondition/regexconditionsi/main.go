package regexconditionsi

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexConditionSi ...
type RegexConditionSi struct {
	Keyword         string
	KeywordValidate string
	V1              *regexp.Regexp
	V2              *regexp.Regexp
	VALIDATEV1      *regexp.Regexp
	VALIDATEV2      *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexConditionSi ...
func NewRegexConditionSi(EL, LL, GL *log.Logger) (*RegexConditionSi, error) {
	var moduleName string = "[RegexConditionSi][NewRegexConditionSi()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}

	compiledE1 := regexp.MustCompile(`^(\s*)(Si)(\s+)`)
	compiledE2 := regexp.MustCompile(`^(\s*)((?i)Si)(\s+)`)
	return &RegexConditionSi{
		Keyword:         "Si",
		KeywordValidate: "hacer",
		V1:              compiledE1,
		V2:              compiledE2,
		VALIDATEV1:      regexp.MustCompile(`(\s*)hacer(\s*)$`),
		VALIDATEV2:      regexp.MustCompile(`(\s*)(?i)hacer(\s*)$`),
		EL:              EL,
		LL:              LL,
		GL:              GL,
	}, nil
}

//StartsWithSi ...
func (r *RegexConditionSi) StartsWithSi(str string, lineIndex int64) bool {

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

	return false
}

//ValidateCondition ...
func (r *RegexConditionSi) ValidateCondition(str string, lineIndex int64) bool {

	if r.VALIDATEV1.MatchString(str) {
		return true
	}

	if r.VALIDATEV2.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		Keyword := strings.Split(r.KeywordValidate, "")
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

//StartsWithSiNoCheck ...
func (r *RegexConditionSi) StartsWithSiNoCheck(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	if r.V2.MatchString(str) {
		return true
	}

	return false

}

//r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)

//LogError ...
//"# Linea | # Columna | Error | Descripcion | Linea del Error"
func (r *RegexConditionSi) LogError(lineIndex int64, columnIndex interface{}, err string, description string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, description, currentLine)
}
