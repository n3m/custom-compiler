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
	KeywordV3       string
	V1              *regexp.Regexp
	V1i             *regexp.Regexp
	VALIDATEV1      *regexp.Regexp
	VALIDATEV2      *regexp.Regexp
	V3              *regexp.Regexp
	V3i             *regexp.Regexp

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

	return &RegexConditionSi{
		Keyword:         "Si",
		KeywordValidate: "hacer",
		KeywordV3:       "sino",
		V1:              regexp.MustCompile(`^(\s*)((?i)Si)(\s+)`),
		V1i:             regexp.MustCompile(`^(\s*)((?i)Si)(\s+)`),
		VALIDATEV1:      regexp.MustCompile(`(\s*)(?i)hacer(\s*)$`),
		VALIDATEV2:      regexp.MustCompile(`(\s*)(?i)hacer(\s*)$`),
		V3:              regexp.MustCompile(`^(\s*)((?i)sino)(\s*)`),
		V3i:             regexp.MustCompile(`^(\s*)((?i)sino)(\s*)`),
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

	if r.V1i.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		Keyword := strings.Split(r.Keyword, "")
		foundTypo := false
		if len(wrongWord) > len(r.Keyword) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)
			return true
		}
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

//StartsWithSino ...
func (r *RegexConditionSi) StartsWithSino(str string, lineIndex int64) bool {
	if r.V3.MatchString(str) {
		return true
	}

	if r.V3i.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		Keyword := strings.Split(r.KeywordV3, "")
		foundTypo := false
		if len(wrongWord) > len(r.KeywordV3) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)
			return true
		}
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
		if len(wrongWord) > len(r.KeywordValidate) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)
			return true
		}
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

	if r.V1i.MatchString(str) {
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
