package regexcustom

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//RegexCustom ...
type RegexCustom struct {
	KeywordLOG1 string
	KeywordLOG2 string
	CTEENT      *regexp.Regexp
	CTEALFA     *regexp.Regexp
	CTEREAL     *regexp.Regexp
	CTELOG      *regexp.Regexp
	CTELOGi     *regexp.Regexp
	ID          *regexp.Regexp
	OPARIT      *regexp.Regexp
	OPREL       *regexp.Regexp
	OPLOG       *regexp.Regexp
	OPLOGi      *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexCustom ...
func NewRegexCustom(EL, LL, GL *log.Logger) (*RegexCustom, error) {
	// var moduleName string = "[regexint][NewRegexCustom()]"

	return &RegexCustom{
		KeywordLOG1: "verdadero",
		KeywordLOG2: "falso",
		CTEALFA:     regexp.MustCompile(`^(\")(\s*)([^"])*(\")$`),
		CTEENT:      regexp.MustCompile(`^((-?)[0-9]+)$`),
		CTEREAL:     regexp.MustCompile(`^(\s*)(-?)(((E|e)(-?)([0-9]*))|([0-9]+\.[0-9]+)|(([0-9]+)(E|e)(-?)([0-9]*)))$`),
		CTELOG:      regexp.MustCompile(`^((?i)verdadero|(?i)falso)$`),
		CTELOGi:     regexp.MustCompile(`^((?i)verdadero|(?i)falso)$`),
		ID:          regexp.MustCompile(`^(\s*)([a-zA-Z]+)([a-zA-Z0-9]*)$`),
		OPARIT:      regexp.MustCompile(`^(\s*)(\+|\-|\*|\/|\%|\^)$`),
		OPREL:       regexp.MustCompile(`^(\s*)(\=|\<\>|\<|\>|\<\=|\>\=)$`),
		OPLOG:       regexp.MustCompile(`^(\s*)(y|o|no)$`),
		OPLOGi:      regexp.MustCompile(`^(\s*)((?i)y|(?i)o|(?i)no)$`),

		GL: GL,
		EL: EL,
		LL: LL,
	}, nil
}

//MatchOpLog ...
func (r *RegexCustom) MatchOpLog(str string) bool {
	if r.OPLOG.MatchString(str) {
		return true
	}
	if r.OPLOGi.MatchString(str) {
		return true
	}
	return false
}

//MatchOpRel ...
func (r *RegexCustom) MatchOpRel(str string) bool {
	if r.OPREL.MatchString(str) {
		return true
	}
	return false
}

//MatchOpArit ...
func (r *RegexCustom) MatchOpArit(str string) bool {
	if r.OPARIT.MatchString(str) {
		return true
	}
	return false
}

//MatchIdent ...
func (r *RegexCustom) MatchIdent(str string) bool {
	if r.ID.MatchString(str) {
		return true
	}
	return false
}

//MatchCteAlfa ...
func (r *RegexCustom) MatchCteAlfa(str string) bool {
	if r.CTEALFA.MatchString(str) {
		return true
	}
	return false
}

//MatchCteEnt ...
func (r *RegexCustom) MatchCteEnt(str string) bool {
	if r.CTEENT.MatchString(str) {
		return true
	}
	return false
}

//MatchCteReal ...
func (r *RegexCustom) MatchCteReal(str string) bool {
	if r.CTEREAL.MatchString(str) {
		return true
	}
	return false
}

//MatchCteLog ...
func (r *RegexCustom) MatchCteLog(str string, lineIndex int64) bool {
	if r.CTELOG.MatchString(str) {
		return true
	}

	if r.CTELOGi.MatchString(str) {
		wrongWord := str
		Keyword := strings.Split(r.KeywordLOG1, "")

		foundTypo := false
		if len(wrongWord) > len(r.KeywordLOG1) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordLOG1), str)
			return true
		}
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					if i == 0 {
						break
					} else {
						foundTypo = true
						r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordLOG1), str)
					}
				}
			}
		}

		Keyword = strings.Split(r.KeywordLOG2, "")
		if len(wrongWord) > len(r.KeywordLOG2) {
			r.LogError(lineIndex, 0, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordLOG2), str)
			return true
		}
		for i, char := range wrongWord {
			if !foundTypo {
				if string(char) != Keyword[i] {
					foundTypo = true
					r.LogError(lineIndex, i, wrongWord, fmt.Sprintf("Found typo in '%+v' declaration. Correct syntax should be '%+v'", wrongWord, r.KeywordLOG2), str)
				}
			}
		}
		return true

	}
	return false
}

//LogError ...
//"# Linea | # Columna | Error | Descripcion | Linea del Error"
func (r *RegexCustom) LogError(lineIndex int64, columnIndex interface{}, err string, description string, currentLine string) {
	log.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.GL.Printf("[ERR] %+v [Line: %+v]", description, lineIndex)
	r.EL.Printf("%+v\t|\t%+v\t|\t%+v\t|\t%+v\t|\t%+v", lineIndex, columnIndex, err, description, currentLine)
}
