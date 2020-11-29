package models

import "strings"

//BlockType ...
type BlockType string

const (
	//NULLBLOCK ...
	NULLBLOCK BlockType = "NULLBLOCK"
	//CONSTANTBLOCK ...
	CONSTANTBLOCK = "CONSTANTBLOCK"
	//VARIABLEBLOCK ...
	VARIABLEBLOCK = "VARIABLEBLOCK"
	//FUNCTIONPROTOBLOCK ...
	FUNCTIONPROTOBLOCK = "FUNCTIONPROTOBLOCK"
	//FUNCTIONBLOCK ...
	FUNCTIONBLOCK = "FUNCTIONBLOCK"
	//PROCEDUREPROTOBLOCK ...
	PROCEDUREPROTOBLOCK = "PROCEDUREPROTOBLOCK"
	//PROCEDUREBLOCK ...
	PROCEDUREBLOCK = "PROCEDUREBLOCK"
	//INITBLOCK ...
	INITBLOCK = "INITBLOCK"
	//REPEATBLOCK ...
	REPEATBLOCK = "REPEATBLOCK"
	//CUANDOBLOCK ...
	CUANDOBLOCK = "CUANDOBLOCK"
	//PROGRAMBLOCK ...
	PROGRAMBLOCK = "PROGRAMBLOCK"
)

//TokenType ...
type TokenType string

const (
	//ENTERO ...
	ENTERO TokenType = "ENTERO"
	//ALFABETICO ...
	ALFABETICO = "ALFABETICO"
	//LOGICO ...
	LOGICO = "LOGICO"
	//REAL ...
	REAL = "REAL"
	// INDEFINIDO ...
	INDEFINIDO = "INDEFINIDO"
)

//TokenComp ...
type TokenComp string

const (
	//CTEENT ...
	CTEENT TokenComp = "CTEENT"
	//CTEALFA ...
	CTEALFA TokenComp = "CTEALFA"
	//CTELOG ...
	CTELOG TokenComp = "CTELOG"
	//CTEREAL ...
	CTEREAL TokenComp = "CTEREAL"
	//DELIM ...
	DELIM TokenComp = "DELIM"
	//BRACK ...
	BRACK TokenComp = "BRACK"
	//OPARIT ...
	OPARIT TokenComp = "OPARIT"
	//OPREL ...
	OPREL TokenComp = "OPREL"
	//OPLOG ...
	OPLOG TokenComp = "OPLOG"
	//PALRES ...
	PALRES TokenComp = "PALRES"
	//OPASIG ...
	OPASIG TokenComp = "OPASIG"
	//ID ...
	ID TokenComp = "ID"
	//CALL ...
	CALL TokenComp = "CALL"
	//PARAM ...
	PARAM TokenComp = "PARAM"
	//NONE ...
	NONE TokenComp = "NONE"
)

//ObjectCodeOperations ...
var ObjectCodeOperations map[string]string = map[string]string{
	"+":    "2",
	"-":    "3",
	"*":    "4",
	"/":    "5",
	"%":    "6",
	"^":    "7",
	"decr": "8",
	"<":    "9",
	">":    "10",
	"<=":   "11",
	">=":   "12",
	"<>":   "13",
	"=":    "14",
	"o":    "15",
	"y":    "16",
	"no":   "17",
}

//VarTypeToTokenType ...
func VarTypeToTokenType(varType string) TokenType {
	var funcType TokenType
	switch strings.TrimSpace(varType) {
	case "Entero", "entero":
		funcType = ENTERO
		break
	case "Real", "real":
		funcType = REAL
		break
	case "Alfabetico", "alfabetico":
		funcType = ALFABETICO
		break
	case "Logico", "logico":
		funcType = LOGICO
		break
	default:
		funcType = INDEFINIDO
		break
	}
	return funcType
}

//ConstTypeToTokenType ...
func ConstTypeToTokenType(constType TokenComp) TokenType {
	var funcType TokenType
	switch constType {
	case CTEENT:
		funcType = ENTERO
		break
	case CTEREAL:
		funcType = REAL
		break
	case CTEALFA:
		funcType = ALFABETICO
		break
	case CTELOG:
		funcType = LOGICO
		break
	default:
		funcType = INDEFINIDO
		break
	}
	return funcType
}
