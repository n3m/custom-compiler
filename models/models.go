package models

//Token ...
type Token struct {
	Type  TokenType
	Key   string
	Value interface{}
}
