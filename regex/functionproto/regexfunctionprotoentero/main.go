package regexfunctionprotoentero

import (
	"regexp"
)

//RegexFuncProtoEntero ...
type RegexFuncProtoEntero struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexFuncProtoEntero ...
func NewRegexFuncProtoEntero() (*RegexFuncProtoEntero, error) {
	// var moduleName string = "[regexint][NewRegexFuncProtoEntero()]"

	return &RegexFuncProtoEntero{
		V1:      regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):Entero;$`),
		V2i:     regexp.MustCompile(`(?i)entero`),
		Keyword: "Entero",
	}, nil
}

//MatchFuncEntero ...
func (r *RegexFuncProtoEntero) MatchFuncEntero(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchFuncEnteroCaseless ...
func (r *RegexFuncProtoEntero) MatchFuncEnteroCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
