package regexfunctionprotoflotante

import (
	"regexp"
)

//RegexFuncProtoFlotante ...
type RegexFuncProtoFlotante struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexFuncProtoFlotante ...
func NewRegexFuncProtoFlotante() (*RegexFuncProtoFlotante, error) {
	// var moduleName string = "[regexint][NewRegexFuncProtoFlotante()]"

	return &RegexFuncProtoFlotante{
		V1:      regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):Flotante;$`),
		V2i:     regexp.MustCompile(`(?i)flotante`),
		Keyword: "Flotante",
	}, nil
}

//MatchFuncFlotante ...
func (r *RegexFuncProtoFlotante) MatchFuncFlotante(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}
	return false
}

//MatchFuncFlotanteCaseless ...
func (r *RegexFuncProtoFlotante) MatchFuncFlotanteCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
