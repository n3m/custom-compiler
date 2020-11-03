package regexregresa

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexRegresa ...
type RegexRegresa struct {
	Keyword string

	Regresa   *regexp.Regexp
	RegresaV2 *regexp.Regexp
	ENDSWITH  *regexp.Regexp
	V3        *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexRegresa ...
func NewRegexRegresa(EL, LL, GL *log.Logger) (*RegexRegresa, error) {
	// var moduleName string = "[regexint][NewRegexRegresa()]"

	return &RegexRegresa{
		Keyword:   "regresa",
		Regresa:   regexp.MustCompile(`^((?i)regresa)(\s*)\(.+\)(\s*);$`),
		RegresaV2: regexp.MustCompile(`^((?i)regresa)(\s*)\(.+\)(\s*)`),
		ENDSWITH:  regexp.MustCompile(`;$`),
		V3:        regexp.MustCompile(`(?m)[rR]egresa\((.*)\)`),

		GL: GL,
		EL: EL,
		LL: LL,
	}, nil
}

//MatchPC ...
func (r *RegexRegresa) MatchPC(str string, lineIndex int64) bool {
	if r.ENDSWITH.MatchString(str) {
		return true
	}

	return false
}

//MatchRegresa ...
func (r *RegexRegresa) MatchRegresa(str string, lineIndex int64) bool {
	if r.Regresa.MatchString(str) {
		return true
	}

	if r.RegresaV2.MatchString(str) {
		strData := strings.Split(str, "(")
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

//GroupsRegresa ...
func (r *RegexRegresa) GroupsRegresa(str string) []string {
	groups := []string{}

	if !r.V3.MatchString(str) {
		return groups
	}

	matched := r.V3.FindAllStringSubmatch(str, -1)
	for _, m := range matched {
		for _, group := range m[1:] {
			if group != "" {
				groups = append(groups, group)
			}
		}
	}

	return groups
}

//LogError ...
//"# Linea | # Columna | Error | DescripcRegresan | Linea del Error"
func (r *RegexRegresa) LogError(lineIndex int64, columnIndex interface{}, err string, descriptRegresan string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", descriptRegresan, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", descriptRegresan, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, descriptRegresan, currentLine)
}
