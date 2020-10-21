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

	compiledE1 := regexp.MustCompile(`^(\s*)((?i)Desde(\s)+el(\s)+valor(\s)+de(\s)+)([a-zA-Z]+[a-zA-Z0-9]*)(\s)*(\:\=)(\s)*(([a-zA-Z]+[a-zA-Z0-9]*)|([0-9]+)|(\(.+\))|(.+))(\s+)+hasta(\s+)(([a-zA-Z]+[a-zA-Z0-9]*)|([0-9]+)|(\(.+\)))(\s?)((decr|incr)(\s+)(([a-zA-Z]+[a-zA-Z0-9]*)|([0-9]+)))?`)

	return &RegexLoopDesde{
		Keyword: "Desde el valor de ",
		V1:      compiledE1,
		EL:      EL,
		LL:      LL,
		GL:      GL,
	}, nil
}

//StartsWithDesde ...
func (r *RegexLoopDesde) StartsWithDesde(str string, lineIndex int64) bool {
	if r.V1.MatchString(str) {
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
