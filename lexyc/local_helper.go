package lexyc

import (
	"go-custom-compiler/models"
)

//DoesTheTokenExistsInGlobalVariables ...
func (l *LexicalAnalyzer) DoesTheTokenExistsInGlobalVariables(token *models.Token) bool {
	for _, each := range l.VariableStorage {
		if each.Key == token.Key {

			return true
		}
	}
	return false
}

//DoesTheTokenExistsInGlobalConstants ...
func (l *LexicalAnalyzer) DoesTheTokenExistsInGlobalConstants(token *models.Token) bool {
	for _, each := range l.ConstantStorage {
		if each.Key == token.Key {
			return true
		}
	}
	return false
}

//DoesTheTokenExistsInLocalVariables ...
func (l *LexicalAnalyzer) DoesTheTokenExistsInLocalVariables(token *models.Token, function *models.TokenFunc) bool {
	for _, each := range function.Vars {

		if each.Key == token.Key {
			return true
		}
	}
	return false
}
