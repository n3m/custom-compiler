package regexprograma

import (
	"fmt"
	"log"
	"regexp"
)

//RegexReserved ...
type RegexReserved struct {
	ALL *regexp.Regexp

	EL *log.Logger
	LL *log.Logger
	GL *log.Logger
}

//NewRegexReserved ...
func NewRegexReserved(EL, LL, GL *log.Logger) (*RegexReserved, error) {
	var moduleName string = "[RegexReserved][NewRegexReserved()]"

	if EL == nil || LL == nil || GL == nil {
		return nil, fmt.Errorf("[ERROR]%+v Loggers came empty", moduleName)
	}
	return &RegexReserved{
		ALL: regexp.MustCompile(`(?m)^([cC]onstantes|[vV]ariables|[rR]eal|[aA]lfabetico|[lL]ogico|[eE]ntero|` +
			`[fF]uncion|[iI]nicio|[fF]in|[dD]e|[pP]rocedimiento|[rR]egresa|[sS]i|[hH]acer|[sS]ino|[cC]uando|[eE]l|[vV]alor|` +
			`[sS]ea|[oO]tro|[dD]esde|[hH]asta|[iI]ncr|[dD]ecr|[rR]epetir|[qQ]ue|[mM]ientras|[sS]e|[cC]umpla|[cC]ontinua|` +
			`[iI]nterrumpe|[lL]impia|[lL]ee|[iI]mprimenl|[iI]mprime|[yY]|[oO]|[nN]o)$`),

		GL: GL,
		EL: EL,
		LL: LL,
	}, nil
}

//IsReserved ...
func (r *RegexReserved) IsReserved(str string) bool {
	if r.ALL.MatchString(str) {
		return true
	}

	return false
}
