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
	LA   *lexyc.LexicalAnalyzer //Lexical Analyzer
	OCL  *log.Logger            //Object Code Logger
	Vars map[string]string
}

//NewCodeGenerator ...
func NewCodeGenerator(lexicalAnalyzer *lexyc.LexicalAnalyzer, objectCodeLogger *log.Logger) (*CodeGenerator, error) {
	var moduleName string = "[Object][NewCodeGenerator()]"

	if lexicalAnalyzer == nil {
		return nil, fmt.Errorf("[ERR]%+v lexical analyzer is not present", moduleName)
	}

	lexicalAnalyzer.GL.Printf("Started the  Code Generator")

	return &CodeGenerator{
		LA:   lexicalAnalyzer,
		OCL:  objectCodeLogger,
		Vars: make(map[string]string),
	}, nil
}

//Generate ...
func (cG *CodeGenerator) Generate() error {

	for _, globalVariable := range cG.LA.VariableStorage {
		cG.printToken(globalVariable, "V", "")
	}
	for _, globalConstant := range cG.LA.ConstantStorage {
		cG.printToken(globalConstant, "C", "")
	}
	for _, function := range cG.LA.FunctionStorage {
		cG.printFunctionToken(function)
	}

	for _, line := range cG.Vars {
		cG.OCL.Println(line)
	}

	for label, line := range cG.LA.HashTable.Labels {
		cG.OCL.Printf("%v,I,I,%v,0,#,", label, line)
	}
	cG.OCL.Println("@")
	for _, line := range cG.LA.HashTable.Lines {
		cG.OCL.Println(line)
	}

	return nil
}

func (cG *CodeGenerator) printToken(token *models.Token, tokenType string, context string) {
	tokenProp := []string{}
	if tokenType != "P" {
		tokenProp = append(tokenProp, token.Key)
		if context != "" {
			tokenProp = append(tokenProp, []string{"I", "I", "0", "0"}...)
		}
	}

	tokenProp = append(tokenProp, tokenType)

	tokenProp = append(tokenProp, string(token.Type[0]))

	for i := 0; i < 2; i++ {
		dimension := "0"
		if i < len(token.Dimensions) {
			dimension = fmt.Sprintf("%v", token.Dimensions[i])
		}
		tokenProp = append(tokenProp, dimension)
	}

	if context != "" {
		tokenProp = append(tokenProp, context)
	}
	tokenProp = append(tokenProp, "#,")

	if cG.Vars[token.Key] != "" {
		cG.Vars[token.Key] = strings.ReplaceAll(cG.Vars[token.Key], "#,", strings.Join(tokenProp, ","))
	} else {
		cG.Vars[token.Key] = strings.Join(tokenProp, ",")
	}
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

	tokenDefinition := token.HashTableLineIndex
	tokenProp = append(tokenProp, tokenDefinition)
	tokenProp = append(tokenProp, "0,#,")

	cG.OCL.Println(strings.Join(tokenProp, ","))

	for _, funcParam := range token.Params {
		cG.printToken(funcParam, "P", token.Key)
	}
	for _, funcVar := range token.Vars {
		cG.printToken(funcVar, "V", token.Key)
	}

	return
}
