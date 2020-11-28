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

//DoesTheTokenExistsInFunctionsStorage ...
func (l *LexicalAnalyzer) DoesTheTokenExistsInFunctionsStorage(token *models.Token) bool {
	for _, each := range l.FunctionStorage {
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

//RetrieveLocalVariableIfExists ...
func (l *LexicalAnalyzer) RetrieveLocalVariableIfExists(token *models.Token, function *models.TokenFunc) *models.Token {
	for _, each := range function.Vars {

		if each.Key == token.Key {
			return each
		}
	}
	return nil
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

//RetrieveGlobalConstantIfExists ...
func (l *LexicalAnalyzer) RetrieveGlobalConstantIfExists(token *models.Token) *models.Token {
	for _, each := range l.ConstantStorage {
		if each.Key == token.Key {

			return each
		}
	}
	return nil
}

//RetrieveFunctionOrProcedureIfExists ...
func (l *LexicalAnalyzer) RetrieveFunctionOrProcedureIfExists(token *models.Token) *models.TokenFunc {
	for _, each := range l.FunctionStorage {

		if each.Key == token.Key {
			return each
		}
	}
	return nil
}
