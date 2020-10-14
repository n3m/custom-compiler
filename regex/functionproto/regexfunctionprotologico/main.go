package regexfunctionprotologico

import (
	"regexp"
)

//RegexFuncProtoLogico ...
type RegexFuncProtoLogico struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexFuncProtoLogico ...
func NewRegexFuncProtoLogico() (*RegexFuncProtoLogico, error) {
	// var moduleName string = "[regexint][NewRegexFuncProtoLogico()]"

	return &RegexFuncProtoLogico{
		V1:      regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):Logico;$`),
		V2i:     regexp.MustCompile(`(?i)logico`),
		Keyword: "Logico",
	}, nil
}

//MatchFuncLogico ...
func (r *RegexFuncProtoLogico) MatchFuncLogico(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchFuncLogicoCaseless ...
func (r *RegexFuncProtoLogico) MatchFuncLogicoCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
