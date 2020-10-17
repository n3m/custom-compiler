package regexfunctionprotoreal

import (
	"regexp"
)

//RegexFuncProtoReal ...
type RegexFuncProtoReal struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexFuncProtoReal ...
func NewRegexFuncProtoReal() (*RegexFuncProtoReal, error) {
	// var moduleName string = "[regexint][NewRegexFuncProtoReal()]"

	return &RegexFuncProtoReal{
		V1:      regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):Real;$`),
		V2i:     regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):(?i)real;$`),
		Keyword: "Real",
	}, nil
}

//MatchFuncReal ...
func (r *RegexFuncProtoReal) MatchFuncReal(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchFuncRealCaseless ...
func (r *RegexFuncProtoReal) MatchFuncRealCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
