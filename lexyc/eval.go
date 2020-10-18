package lexyc

import "go-custom-compiler/models"

//ExpectIdent ...
func (l *LexicalAnalyzer) ExpectIdent(currentLine string, lineIndex int64) bool {
	if len(l.OpQueue) != 1 {
		l.LogError(lineIndex, "N/A", "N/A", "Number of Parameters is greater or less than expected", currentLine)
		return false
	}

	if len(l.OpQueue) > 0 {
		if l.OpQueue[0] != models.ID {
			l.LogError(lineIndex, "N/A", "N/A", "Expected parameter of type ID", currentLine)
			return false
		}

		return true
	}

	return false

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
