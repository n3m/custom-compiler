package regexloophastaque

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexLoopHastaQue ...
type RegexLoopHastaQue struct {
	Keyword string
	V1      *regexp.Regexp
	V2      *regexp.Regexp
	V3      *regexp.Regexp
	V4      *regexp.Regexp
	V5      *regexp.Regexp
	V6      *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexLoopHastaQue ...
func NewRegexLoopHastaQue(EL, LL, GL *log.Logger) (*RegexLoopHastaQue, error) {
	var moduleName string = "[RegexLoopHastaQue][NewRegexLoopHastaQue()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}

	compiledE1 := regexp.MustCompile(`^(\s*)(Hasta(\s)+que)`)
	compiledE2 := regexp.MustCompile(`^(\s*)((?i)hasta(\s)+(?i)que)`)
	compiledE3 := regexp.MustCompile(`^(\s*)((?i)hasta(\s)+(?i)q)`)
	compiledE4 := regexp.MustCompile(`^(\s*)(?i)hasta(\s)+`)
	compiledE5 := regexp.MustCompile(`^(\s*)(?i)hast`)
	compiledE6 := regexp.MustCompile(`^(\s*)(?i)ha`)
	return &RegexLoopHastaQue{
		Keyword: "Hasta que",
		V1:      compiledE1,
		V2:      compiledE2,
		V3:      compiledE3,
		V4:      compiledE4,
		V5:      compiledE5,
		V6:      compiledE6,
		EL:      EL,
		LL:      LL,
		GL:      GL,
	}, nil
}

//StartsWithHastaQue ...
func (r *RegexLoopHastaQue) StartsWithHastaQue(str string, lineIndex int64) bool {
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
					//"# Linea | # Columna | Error | Descripcion | Linea del Error"
					r.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, i, wrongWord, r.Keyword, str)
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
					//"# Linea | # Columna | Error | Descripcion | Linea del Error"
					r.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, i, wrongWord, r.Keyword, str)
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
					log.Printf("[ERR] Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.Keyword)
					r.GL.Printf("[ERR] Found typo in '%+v' declaration at [%+v]. Correct syntax should be '%+v'", wrongWord, i, r.Keyword)
					//"# Linea | # Columna | Error | Descripcion | Linea del Error"
					r.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, i, wrongWord, r.Keyword, str)
				}
			}
		}
		return true
	}

	if r.V5.MatchString(str) {
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
					//"# Linea | # Columna | Error | Descripcion | Linea del Error"
					r.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, i, wrongWord, r.Keyword, str)
				}
			}
		}
		return true
	}

	if r.V6.MatchString(str) {
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
					//"# Linea | # Columna | Error | Descripcion | Linea del Error"
					r.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, i, wrongWord, r.Keyword, str)
				}
			}
		}
		return true
	}

	return false
}

//StartsWithHastaQueNoCheck ...
func (r *RegexLoopHastaQue) StartsWithHastaQueNoCheck(str string) bool {
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
