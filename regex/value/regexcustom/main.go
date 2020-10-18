package regexcustom

import (
	"regexp"
)

//RegexCustom ...
type RegexCustom struct {
	Keyword string
	CTEENT  *regexp.Regexp
	CTEALFA *regexp.Regexp
	CTEREAL *regexp.Regexp
	CTELOG  *regexp.Regexp
	CTELOGi *regexp.Regexp
	ID      *regexp.Regexp
	OPARIT  *regexp.Regexp
	OPREL   *regexp.Regexp
	OPLOG   *regexp.Regexp
	OPLOGi  *regexp.Regexp
}

//NewRegexCustom ...
func NewRegexCustom() (*RegexCustom, error) {
	// var moduleName string = "[regexint][NewRegexCustom()]"

	return &RegexCustom{
		CTEALFA: regexp.MustCompile(`^(\")(\s*)([^"])*(\")$`),
		CTEENT:  regexp.MustCompile(`(0-9)+`),
		CTEREAL: regexp.MustCompile(`^(\s*)(-?)(((E|e)(-?)([0-9]*))|([0-9]+\.[0-9]+)|(([0-9]+)(E|e)(-?)([0-9]*)))$`),
		CTELOG:  regexp.MustCompile(`^(verdadero|falso)$`),
		CTELOGi: regexp.MustCompile(`^((?i)verdadero|(?i)falso)$`),
		ID:      regexp.MustCompile(`^(\s*)([a-zA-Z]+)([a-zA-Z0-9])*$`),
		OPARIT:  regexp.MustCompile(`^(\s*)(\+|\-|\*|\/|\%|\^)$`),
		OPREL:   regexp.MustCompile(`^(\s*)(\=|\<\>|\<|\>|\<\=|\>\=)$`),
		OPLOG:   regexp.MustCompile(`^(\s*)(y|o|no)$`),
		OPLOGi:  regexp.MustCompile(`^(\s*)((?i)y|(?i)o|(?i)no)$`),
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
func (r *RegexCustom) MatchCteLog(str string) bool {
	if r.CTELOG.MatchString(str) {
		return true
	}

	if r.CTELOGi.MatchString(str) {
		return true
		//TODO: Check error
	}
	return false
}
