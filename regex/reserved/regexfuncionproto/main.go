package regexfuncionproto

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexFuncionProto ...
type RegexFuncionProto struct {
	Keyword string
	V1      *regexp.Regexp
	V2      *regexp.Regexp
	V3      *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexFuncionProto ...
func NewRegexFuncionProto(EL, LL, GL *log.Logger) (*RegexFuncionProto, error) {
	var moduleName string = "[regexFuncionProto][NewRegexFuncionProto()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}

	compiledV1 := regexp.MustCompile("^Funcion")
	compiledV2 := regexp.MustCompile("^Funci")
	compiledV3 := regexp.MustCompile("^fun")
	return &RegexFuncionProto{
		Keyword: "Funcion",
		V1:      compiledV1,
		V2:      compiledV2,
		V3:      compiledV3,
		EL:      EL,
		LL:      LL,
		GL:      GL,
	}, nil
}

//StartsWithFuncionProto ...
func (r *RegexFuncionProto) StartsWithFuncionProto(str string) bool {
	if r.V1.MatchString(str) {
		return true
	}

	if r.V2.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		Keyword := strings.Split(r.Keyword, "")
		foundTypo := false
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					log.Printf("[ERR] Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.Keyword)
					r.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.Keyword)

				}
			}
		}
		return true
	}

	if r.V3.MatchString(str) {
		strData := strings.Split(str, " ")
		wrongWord := strData[0]
		Keyword := strings.Split(r.Keyword, "")
		foundTypo := false
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					log.Printf("[ERR] Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.Keyword)
					r.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.Keyword)

				}
			}
		}
		return true
	}

	return false
}

//StartsWithFuncionProtoNoCheck ...
func (r *RegexFuncionProto) StartsWithFuncionProtoNoCheck(str string) bool {
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
