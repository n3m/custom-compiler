package regexio

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexIO ...
type RegexIO struct {
	KeywordImprimenl string
	KeywordImprime   string
	KeywordLee       string
	IMPRIMENL        *regexp.Regexp
	IMPRIMENLi       *regexp.Regexp
	IMPRIME          *regexp.Regexp
	IMPRIMEi         *regexp.Regexp
	LEE              *regexp.Regexp
	LEEi             *regexp.Regexp
	ENDSWITH         *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexIO ...
func NewRegexIO(EL, LL, GL *log.Logger) (*RegexIO, error) {
	// var moduleName string = "[regexint][NewRegexIO()]"

	return &RegexIO{
		KeywordImprime:   "Imprime",
		KeywordImprimenl: "Imprimenl",
		KeywordLee:       "Lee",
		IMPRIMENL:        regexp.MustCompile(`^((?i)Imprimenl)`),
		IMPRIMENLi:       regexp.MustCompile(`^((?i)Imprimenl)`),
		IMPRIME:          regexp.MustCompile(`^((?i)Imprime)`),
		IMPRIMEi:         regexp.MustCompile(`^((?i)Imprime)`),
		LEE:              regexp.MustCompile(`^((?i)Lee)(\()`),
		LEEi:             regexp.MustCompile(`^((?i)Lee)(\()`),
		ENDSWITH:         regexp.MustCompile(`;$`),

		GL: GL,
		EL: EL,
		LL: LL,
	}, nil
}

//MatchPC ...
func (r *RegexIO) MatchPC(str string, lineIndex int64) bool {
	if r.ENDSWITH.MatchString(str) {
		return true
	}

	return false
}

//MatchImprimenl ...
func (r *RegexIO) MatchImprimenl(str string, lineIndex int64) bool {
	if r.IMPRIMENL.MatchString(str) {
		return true
	}

	if r.IMPRIMENLi.MatchString(str) {
		strData := splitAtCharRespectingQuotes(str, '(')
		wrongWord := strData[0]
		Keyword := strings.Split(r.KeywordImprimenl, "")
		foundTypo := false
		if len(wrongWord) > len(r.KeywordImprimenl) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordImprimenl), str)
			return true
		}
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordImprimenl), str)
				}
			}
		}
		return true
	}
	return false
}

//MatchImprime ...
func (r *RegexIO) MatchImprime(str string, lineIndex int64) bool {
	if r.IMPRIME.MatchString(str) {
		return true
	}
	if r.IMPRIMEi.MatchString(str) {
		strData := splitAtCharRespectingQuotes(str, '(')
		wrongWord := strData[0]
		Keyword := strings.Split(r.KeywordImprime, "")
		foundTypo := false
		if len(wrongWord) > len(r.KeywordImprime) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordImprime), str)
			return true
		}
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordImprime), str)
				}
			}
		}
		return true
	}
	return false
}

//MatchLee ...
func (r *RegexIO) MatchLee(str string, lineIndex int64) bool {
	if r.LEE.MatchString(str) {
		return true
	}
	if r.LEEi.MatchString(str) {
		strData := splitAtCharRespectingQuotes(str, '(')
		wrongWord := strData[0]
		Keyword := strings.Split(r.KeywordLee, "")
		foundTypo := false
		if len(wrongWord) > len(r.KeywordLee) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordLee), str)
			return true
		}
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordLee), str)
				}
			}
		}
		return true
	}
	return false
}

//LogError ...
//"# Linea | # Columna | Error | Descripcion | Linea del Error"
func (r *RegexIO) LogError(lineIndex int64, columnIndex interface{}, err string, description string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, description, currentLine)
}

func splitAtCharRespectingQuotes(s string, char byte) []string {
	res := []string{}
	var beg int
	var inString bool

	for i := 0; i < len(s); i++ {
		if s[i] == char && !inString {
			res = append(res, s[beg:i])
			beg = i + 1
		} else if s[i] == '"' {
			if !inString {
				inString = true
			} else if i > 0 && s[i-1] != '\\' {
				inString = false
			}
		}
	}
	return append(res, s[beg:])
}
