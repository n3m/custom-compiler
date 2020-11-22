package models

//Token ...
type Token struct {
	Type       TokenType
	Key        string
	Value      interface{}
	Dimensions []int
}

//TokenFunc ...
type TokenFunc struct {
	Type   TokenType
	Key    string
	Params []Token
	Vars   []Token
}
