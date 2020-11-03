package regexfunction

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexFunction ...
type RegexFunction struct {
	Keyword string
	V1      *regexp.Regexp
	V2      *regexp.Regexp
	V3      *regexp.Regexp
	V4End   *regexp.Regexp
	Call    *regexp.Regexp
	Call2   *regexp.Regexp
	CallEnd *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexFunction ...
func NewRegexFunction(EL, LL, GL *log.Logger) (*RegexFunction, error) {
	var moduleName string = "[regexFunction][NewRegexFunction()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}

	compiledV1 := regexp.MustCompile("^(?i)Funcion")
	compiledV2 := regexp.MustCompile("^(?i)Func")
	compiledV3 := regexp.MustCompile("^(?i)Fu")
	compiledV4End := regexp.MustCompile("[^;]$")
	compiledCall := regexp.MustCompile(`.*\(.*\)`)
	compiledCall2 := regexp.MustCompile(`.*\(.*`)
	compiledCallEnd := regexp.MustCompile(`.*\)`)
	return &RegexFunction{
		Keyword: "Funcion",
		V1:      compiledV1,
		V2:      compiledV2,
		V3:      compiledV3,
		V4End:   compiledV4End,
		Call:    compiledCall,
		Call2:   compiledCall2,
		CallEnd: compiledCallEnd,
		EL:      EL,
		LL:      LL,
		GL:      GL,
	}, nil
}

//StartsWithFunction ...
func (r *RegexFunction) StartsWithFunction(str string, lineIndex int64) bool {
	if r.V1.MatchString(str) && r.V4End.MatchString(str) {
		return true
	}

	if r.V2.MatchString(str) && r.V4End.MatchString(str) {
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

	if r.V3.MatchString(str) && r.V4End.MatchString(str) {
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

//StartsWithFunctionNoCheck ...
func (r *RegexFunction) StartsWithFunctionNoCheck(str string) bool {
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

//MatchFunctionCall ...
func (r *RegexFunction) MatchFunctionCall(str string) bool {
	return r.Call.MatchString(str)
}

//MatchFunctionCall2 ...
func (r *RegexFunction) MatchFunctionCall2(str string) bool {
	return r.Call2.MatchString(str)
}

//MatchFunctionCallEnd ...
func (r *RegexFunction) MatchFunctionCallEnd(str string) bool {
	return r.CallEnd.MatchString(str)
}

//r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)

//LogError ...
//"# Linea | # Columna | Error | Descripcion | Linea del Error"
func (r *RegexFunction) LogError(lineIndex int64, columnIndex interface{}, err string, description string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, description, currentLine)
}
