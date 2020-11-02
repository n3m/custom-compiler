package models

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
	//NONE ...
	NONE TokenComp = "NONE"
)
