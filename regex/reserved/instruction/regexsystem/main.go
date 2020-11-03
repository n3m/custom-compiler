package regexsystem

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexSystem ...
type RegexSystem struct {
	Keyword1 string
	W1       *regexp.Regexp
	W1_1     *regexp.Regexp
	W1_2     *regexp.Regexp

	Keyword2 string
	W2       *regexp.Regexp
	W2_1     *regexp.Regexp
	W2_2     *regexp.Regexp

	ENDSWITH *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexSystem ...
func NewRegexSystem(EL, LL, GL *log.Logger) (*RegexSystem, error) {
	// var moduleName string = "[regexint][NewRegexSystem()]"

	return &RegexSystem{
		ENDSWITH: regexp.MustCompile(`;$`),

		Keyword1: "interrumpe",
		W1:       regexp.MustCompile(`^((?i)interrumpe)(\s*);$`),
		W1_1:     regexp.MustCompile(`^((?i)interru)(\s*)`),
		W1_2:     regexp.MustCompile(`^((?i)int)(\s*)`),

		Keyword2: "limpia",
		W2:       regexp.MustCompile(`^((?i)limpia)(\s*);$`),
		W2_1:     regexp.MustCompile(`^((?i)limp)(\s*)`),
		W2_2:     regexp.MustCompile(`^((?i)li)(\s*)`),

		GL: GL,
		EL: EL,
		LL: LL,
	}, nil
}

//MatchPC ...
func (r *RegexSystem) MatchPC(str string, lineIndex int64) bool {
	if r.ENDSWITH.MatchString(str) {
		return true
	}

	return false
}

//MatchInterrumpe ...
func (r *RegexSystem) MatchInterrumpe(str string, lineIndex int64) bool {
	if r.W1.MatchString(str) {
		return true
	}

	if r.W1_1.MatchString(str) {
		wrongWord := str
		Keyword := strings.Split(r.Keyword1, "")
		foundTypo := false
		if len(wrongWord) > len(r.Keyword1) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword1), str)
			return true
		}
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword1), str)
				}
			}
		}
		return true
	}

	if r.W1_2.MatchString(str) {
		wrongWord := str
		Keyword := strings.Split(r.Keyword1, "")
		foundTypo := false
		if len(wrongWord) > len(r.Keyword1) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword1), str)
			return true
		}
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword1), str)
				}
			}
		}
		return true
	}
	return false
}

//MatchLimpia ...
func (r *RegexSystem) MatchLimpia(str string, lineIndex int64) bool {
	if r.W2.MatchString(str) {
		return true
	}

	if r.W2_1.MatchString(str) {
		wrongWord := str
		Keyword := strings.Split(r.Keyword2, "")
		foundTypo := false
		if len(wrongWord) > len(r.Keyword2) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword2), str)
			return true
		}
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword2), str)
				}
			}
		}
		return true
	}

	if r.W2_2.MatchString(str) {
		wrongWord := str
		Keyword := strings.Split(r.Keyword2, "")
		foundTypo := false
		if len(wrongWord) > len(r.Keyword2) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword2), str)
			return true
		}
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.Keyword2), str)
				}
			}
		}
		return true
	}
	return false
}

//LogError ...
//"# Linea | # Columna | Error | DescripcSystemn | Linea del Error"
func (r *RegexSystem) LogError(lineIndex int64, columnIndex interface{}, err string, descriptSystemn string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", descriptSystemn, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", descriptSystemn, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, descriptSystemn, currentLine)
}
