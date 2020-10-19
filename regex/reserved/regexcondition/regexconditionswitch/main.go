package regexconditionswitch

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexConditionSwitch ...
type RegexConditionSwitch struct {
	KeywordV1 string
	KeywordV2 string
	V1        *regexp.Regexp
	V1i       *regexp.Regexp
	V2        *regexp.Regexp
	V2i       *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexConditionSwitch ...
func NewRegexConditionSwitch(EL, LL, GL *log.Logger) (*RegexConditionSwitch, error) {
	var moduleName string = "[RegexConditionSwitch][NewRegexConditionSwitch()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}

	return &RegexConditionSwitch{
		KeywordV1: "Sea",
		KeywordV2: "Otro",
		V1:        regexp.MustCompile(`^(\s*)(Sea)(\s+)`),
		V1i:       regexp.MustCompile(`^(\s*)((?i)Sea)(\s+)`),
		V2:        regexp.MustCompile(`^(\s*)(Otro)(\s+)`),
		V2i:       regexp.MustCompile(`^(\s*)((?i)Otro)(\s+)`),
		EL:        EL,
		LL:        LL,
		GL:        GL,
	}, nil
}

//StartsWithSea ...
func (r *RegexConditionSwitch) StartsWithSea(str string, lineIndex int64) bool {

	if r.V1.MatchString(str) {
		return true
	}

	if r.V2.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		Keyword := strings.Split(r.KeywordV1, "")
		foundTypo := false
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordV1), str)
				}
			}
		}
		return true
	}

	return false
}

//StartsWithOtro ...
func (r *RegexConditionSwitch) StartsWithOtro(str string, lineIndex int64) bool {

	if r.V1.MatchString(str) {
		return true
	}

	if r.V2.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		Keyword := strings.Split(r.KeywordV2, "")
		foundTypo := false
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordV2), str)
				}
			}
		}
		return true
	}

	return false
}

//StartsWithSeaNoCheck ...
func (r *RegexConditionSwitch) StartsWithSeaNoCheck(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	if r.V2.MatchString(str) {
		return true
	}

	return false
}

//StartsWithOtroNoCheck ...
func (r *RegexConditionSwitch) StartsWithOtroNoCheck(str string) bool {
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
func (r *RegexConditionSwitch) LogError(lineIndex int64, columnIndex interface{}, err string, description string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, description, currentLine)
}
