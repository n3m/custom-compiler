package lexyc

import (
	"go-custom-compiler/models"
	"log"
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

//RetrieveGlobalVarIfExists ...
func (l *LexicalAnalyzer) RetrieveGlobalVarIfExists(token *models.Token) *models.Token {
	for _, each := range l.VariableStorage {
		if each.Key == token.Key {

			return each
		}
	}
	return nil
}

//DoesTheTokenExistsInGlobalConstants ...
func (l *LexicalAnalyzer) DoesTheTokenExistsInGlobalConstants(token *models.Token) bool {
	for _, each := range l.ConstantStorage {
		if each.Key == token.Key {
			log.Printf("\t\t\t TEST EGC > '%+v' == '%+v' :: true", each.Key, token.Key)
			return true
		}
		log.Printf("\t\t\t TEST EGC > '%+v' == '%+v' :: false", each.Key, token.Key)
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

//RetrieveLocalVariableIfExists ...
func (l *LexicalAnalyzer) RetrieveLocalVariableIfExists(token *models.Token, function *models.TokenFunc) *models.Token {
	for _, each := range function.Vars {

		if each.Key == token.Key {
			return each
		}
	}
	return nil
}

//IsAssignmentDataTypeCorrect ...
func (l *LexicalAnalyzer) IsAssignmentDataTypeCorrect() bool {
	return true
}
