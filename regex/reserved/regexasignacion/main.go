package regexasignacion

import (
	"log"
	"regexp"
)

//RegexAsignacion ...
type RegexAsignacion struct {
	Keyword string

	Test     *regexp.Regexp
	ENDSWITH *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexAsignacion ...
func NewRegexAsignacion(EL, LL, GL *log.Logger) (*RegexAsignacion, error) {
	// var moduleName string = "[regexint][NewRegexAsignacion()]"

	return &RegexAsignacion{
		Test:     regexp.MustCompile(`^([a-zA-Z]+([a-zA-Z0-9]*))((\[([a-zA-Z0-9]+)\])*)(\s*):=(\s*)(.*)`),
		ENDSWITH: regexp.MustCompile(`;$`),

		GL: GL,
		EL: EL,
		LL: LL,
	}, nil
}

//MatchPC ...
func (r *RegexAsignacion) MatchPC(str string, lineIndex int64) bool {
	if r.ENDSWITH.MatchString(str) {
		return true
	}

	return false
}

//MatchAsignacion ...
func (r *RegexAsignacion) MatchAsignacion(str string, lineIndex int64) bool {
	if r.Test.MatchString(str) {
		return true
	}

	return false
}

//LogError ...
//"# Linea | # Columna | Error | DescripcAsignacionn | Linea del Error"
func (r *RegexAsignacion) LogError(lineIndex int64, columnIndex interface{}, err string, descriptAsignacionn string, currentLine string) {
	//log.Printf("[ERR] %+v [Line: %+v]", descriptAsignacionn, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", descriptAsignacionn, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, descriptAsignacionn, currentLine)
}
