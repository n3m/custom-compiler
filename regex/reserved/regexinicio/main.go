package regexinicio

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexInicio ...
type RegexInicio struct {
	Keyword string
	V1      *regexp.Regexp
	V2      *regexp.Regexp
	V3      *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexInicio ...
func NewRegexInicio(EL, LL, GL *log.Logger) (*RegexInicio, error) {
	var moduleName string = "[regexInicio][NewRegexInicio()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}
	return &RegexInicio{
		Keyword: "Inicio",
		V1:      regexp.MustCompile("^(?i)Inicio"),
		V2:      regexp.MustCompile("^(?i)Inici"),
		V3:      regexp.MustCompile("^(?i)Ini"),
		GL:      GL,
		EL:      EL,
		LL:      LL,
	}, nil
}

//StartsWithInicio ...
func (r *RegexInicio) StartsWithInicio(str string, lineIndex int64) bool {

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
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)
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
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)
				}
			}
		}
		return true
	}

	return false
}

//StartsWithInicioNoCheck ...
func (r *RegexInicio) StartsWithInicioNoCheck(str string) bool {
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

//r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword), str)

//LogError ...
//"# Linea | # Columna | Error | Descripcion | Linea del Error"
func (r *RegexInicio) LogError(lineIndex int64, columnIndex interface{}, err string, description string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, description, currentLine)
}
