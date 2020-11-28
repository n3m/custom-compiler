package object

import (
	"fmt"
	"log"
	"strings"

	"go-custom-compiler/lexyc"
	"go-custom-compiler/models"
)

//CodeGenerator ...
type CodeGenerator struct {
	LA  *lexyc.LexicalAnalyzer //Lexical Analyzer
	OCL *log.Logger            //Object Code Logger
}

//NewCodeGenerator ...
func NewCodeGenerator(lexicalAnalyzer *lexyc.LexicalAnalyzer, objectCodeLogger *log.Logger) (*CodeGenerator, error) {
	var moduleName string = "[Object][NewCodeGenerator()]"

	if lexicalAnalyzer == nil {
		return nil, fmt.Errorf("[ERR]%+v lexical analyzer is not present", moduleName)
	}

	lexicalAnalyzer.GL.Printf("Started the  Code Generator")

	return &CodeGenerator{
		LA:  lexicalAnalyzer,
		OCL: objectCodeLogger,
	}, nil
}

//Generate ...
func (cG *CodeGenerator) Generate() error {

	for _, globalVariable := range cG.LA.VariableStorage {
		cG.printToken(globalVariable, "V", []string{})
	}
	for _, globalConstant := range cG.LA.ConstantStorage {
		cG.printToken(globalConstant, "V", []string{})
	}
	for _, function := range cG.LA.FunctionStorage {
		cG.printFunctionToken(function)
	}

	return nil
}

func (cG *CodeGenerator) printToken(token *models.Token, tokenType string, definition []string) {
	tokenProp := []string{}
	tokenProp = append(tokenProp, token.Key)
	tokenProp = append(tokenProp, definition...)

	tokenProp = append(tokenProp, tokenType)

	tokenProp = append(tokenProp, string(token.Type[0]))

	for i := 0; i < 2; i++ {
		dimension := "0"
		if i < len(token.Dimensions) {
			dimension = fmt.Sprintf("%v", token.Dimensions[i])
		}
		tokenProp = append(tokenProp, dimension)
	}

	tokenProp = append(tokenProp, "#,")

	cG.OCL.Println(strings.Join(tokenProp, ","))
	return
}

func (cG *CodeGenerator) printFunctionToken(token *models.TokenFunc) {
	tokenProp := []string{}
	tokenProp = append(tokenProp, token.Key)

	tokenType := "P"
	functionType := "I"
	if token.Type != "" {
		tokenType = "F"
		functionType = string(token.Type[0])
	}
	tokenProp = append(tokenProp, tokenType)
	tokenProp = append(tokenProp, functionType)

	tokenDefinition := "0" //TODO
	tokenProp = append(tokenProp, tokenDefinition)
	tokenProp = append(tokenProp, "0,#,")

	cG.OCL.Println(strings.Join(tokenProp, ","))

	for _, funcParam := range token.Params {
		cG.printToken(funcParam, "V", []string{"I", "I", "0", "0"})
	}
	for _, funcVar := range token.Vars {
		cG.printToken(funcVar, "V", []string{"I", "I", "0", "0"})
	}

	return
}
