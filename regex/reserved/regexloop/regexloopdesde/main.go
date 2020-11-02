package regexloopdesde

import (
	"fmt"
	"log"
	"regexp"
)

//RegexLoopDesde ...
type RegexLoopDesde struct {
	Keyword string
	V1      *regexp.Regexp
	V2      *regexp.Regexp
	V3      *regexp.Regexp
	V4      *regexp.Regexp
	V5      *regexp.Regexp
	V6      *regexp.Regexp

	WTest1 *regexp.Regexp
	WTest2 *regexp.Regexp
	WTest3 *regexp.Regexp
	WTest4 *regexp.Regexp
	WTest5 *regexp.Regexp
	WTest6 *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexLoopDesde ...
func NewRegexLoopDesde(EL, LL, GL *log.Logger) (*RegexLoopDesde, error) {
	var moduleName string = "[RegexLoopDesde][NewRegexLoopDesde()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}

	return &RegexLoopDesde{
		Keyword: "Desde el valor de ",
		V1:      regexp.MustCompile(`^(\s*)((?i)Desde(\s)+el(\s)+valor(\s)+de(\s)+)([a-zA-Z]+[a-zA-Z0-9]*)(\s)*(\:\=)(\s)*(([a-zA-Z]+[a-zA-Z0-9]*)|([0-9]+)|(\(.+\))|(.+))(\s+)+hasta(\s+)(([a-zA-Z]+[a-zA-Z0-9]*)|([0-9]+)|(\(.+\)))(\s?)((decr|incr)(\s+)(([a-zA-Z]+[a-zA-Z0-9]*)|([0-9]+)))?`),
		V2:      regexp.MustCompile(`^(\s*)((?i)Desde(\s)+el(\s)+valor(\s)+de(\s)+)`),
		V3:      regexp.MustCompile(`^(\s*)((?i)Desde(\s)+el(\s)+valor(\s)+)`),
		V4:      regexp.MustCompile(`^(\s*)((?i)Desde(\s)+el(\s)+)`),
		V5:      regexp.MustCompile(`^(\s*)((?i)Desde(\s)+)`),
		V6:      regexp.MustCompile(`^(\s*)((?i)Des)`),

		WTest1: regexp.MustCompile(`((?i)Desde)`),
		WTest2: regexp.MustCompile(`((?i)el)`),
		WTest3: regexp.MustCompile(`((?i)valor)`),
		WTest4: regexp.MustCompile(`((?i)de)`),
		WTest5: regexp.MustCompile(`((?i)hasta)`),
		EL:     EL,
		LL:     LL,
		GL:     GL,
	}, nil
}

//StartsWithDesde ...
func (r *RegexLoopDesde) StartsWithDesde(str string, lineIndex int64) bool {
	if r.V1.MatchString(str) {
		return true
	}

	if r.V2.MatchString(str) {
		r.LogError(lineIndex, 0, str, fmt.Sprintf("Found invalid syntax. Correct syntax should be '%+v'", "hasta"), str)
		return true
	}
	if r.V3.MatchString(str) {
		r.LogError(lineIndex, 0, str, fmt.Sprintf("Found invalid syntax. Correct syntax should be '%+v'", "de"), str)
		return true
	}
	if r.V4.MatchString(str) {
		r.LogError(lineIndex, 0, str, fmt.Sprintf("Found invalid syntax. Correct syntax should be '%+v'", "valor"), str)
		return true
	}
	if r.V5.MatchString(str) {
		r.LogError(lineIndex, 0, str, fmt.Sprintf("Found invalid syntax. Correct syntax should be '%+v'", "el"), str)
		return true
	}
	if r.V6.MatchString(str) {
		r.LogError(lineIndex, 0, str, fmt.Sprintf("Found invalid syntax. Correct syntax should be '%+v'", "desde"), str)
		return true
	}

	return false
}

//StartsWithDesdeNoCheck ...
func (r *RegexLoopDesde) StartsWithDesdeNoCheck(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false

}

//r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)

//LogError ...
//"# Linea | # Columna | Error | Descripcion | Linea del Error"
func (r *RegexLoopDesde) LogError(lineIndex int64, columnIndex interface{}, err string, description string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, description, currentLine)
}
