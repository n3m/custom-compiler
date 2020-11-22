package regexfunctionprotoalfabetico

import (
	"regexp"
)

//RegexFuncProtoAlfabetico ...
type RegexFuncProtoAlfabetico struct {
	Keyword string
	V1      *regexp.Regexp
	V2i     *regexp.Regexp
}

//NewRegexFuncProtoAlfabetico ...
func NewRegexFuncProtoAlfabetico() (*RegexFuncProtoAlfabetico, error) {
	// var moduleName string = "[regexint][NewRegexFuncProtoAlfabetico()]"

	return &RegexFuncProtoAlfabetico{
		V1:      regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):(?i)Alfabetico;$`),
		V2i:     regexp.MustCompile(`\s*[a-zA-Z]+[a-zA-Z0-9]*\(([a-zA-Z]+[a-zA-Z0-9]*(\s*,\s*[a-zA-Z]+[a-zA-Z0-9]*)*):[a-zA-Z]+\):(?i)alfabetico;$`),
		Keyword: "Alfabetico",
	}, nil
}

//MatchFuncAlfabetico ...
func (r *RegexFuncProtoAlfabetico) MatchFuncAlfabetico(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	return false
}

//MatchFuncAlfabeticoCaseless ...
func (r *RegexFuncProtoAlfabetico) MatchFuncAlfabeticoCaseless(str string) bool {
	if r.V2i.MatchString(str) {
		return true
	}
	return false
}
