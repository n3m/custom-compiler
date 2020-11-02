package lexyc

import (
	"go-custom-compiler/helpers"
	"go-custom-compiler/models"
)

//ExpectIdent ...
func (l *LexicalAnalyzer) ExpectIdent(currentLine string, lineIndex int64) bool {
	if len(l.OpQueue) < 1 {
		l.LogError(lineIndex, "N/A", "UNEXPECTED", "There are no parameters", currentLine)
		return false
	}

	noIDs := -1
	noBrackets := 0
	for _, item := range l.OpQueue {
		if item == models.ID {
			noIDs++
		} else if item == models.BRACK {
			noBrackets++
		} else {
			l.LogError(lineIndex, "N/A", "UNEXPECTED", "Expected "+helpers.IDENTIFICADOR, currentLine)
			return false
		}
	}

	if noBrackets != noIDs*2 {
		l.LogError(lineIndex, "N/A", "UNEXPECTED", "Mismatched number of brackets", currentLine)
		return false
	}
	return true
}

//ExpectCondition ...
func (l *LexicalAnalyzer) ExpectCondition(currentLine string, lineIndex int64) bool {
	if len(l.OpQueue) >= 3 {
		l.LogError(lineIndex, "N/A", "N/A", "Number of Parameters is less than expected", currentLine)
		return false
	}

	return false
}

//ExpectNoNone ...
func (l *LexicalAnalyzer) ExpectNoNone() bool {
	for _, t := range l.OpQueue {
		if t == models.NONE {
			return false
		}
	}

	return true
}
