package regexfin

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexFin ...
type RegexFin struct {
	Keyword  string
	V1       *regexp.Regexp
	V2       *regexp.Regexp
	V3       *regexp.Regexp
	V4       *regexp.Regexp
	ENDSWITH *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexFin ...
func NewRegexFin(EL, LL, GL *log.Logger) (*RegexFin, error) {
	var moduleName string = "[RegexFin][NewRegexFin()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}
	return &RegexFin{
		Keyword:  "Fin",
		V1:       regexp.MustCompile(`^(\s*)(?i)Fin(\s*);$`),
		V2:       regexp.MustCompile(`^(\s*)(?i)Fin(\s*)$`),
		V3:       regexp.MustCompile(`^(\s*)(?i)Fi(\s*)$`),
		V4:       regexp.MustCompile(`^(\s*)(?i)F(\s*)$`),
		ENDSWITH: regexp.MustCompile(`;$`),

		GL: GL,
		EL: EL,
		LL: LL,
	}, nil
}

//StartsWithFin ...
func (r *RegexFin) StartsWithFin(str string, lineIndex int64) bool {

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

	if r.V4.MatchString(str) {
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

//StartsWithFinNoCheck ...
func (r *RegexFin) StartsWithFinNoCheck(str string) bool {
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

//LogError ...
//"# Linea | # Columna | Error | Descripcion | Linea del Error"
func (r *RegexFin) LogError(lineIndex int64, columnIndex interface{}, err string, description string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, description, currentLine)
}
