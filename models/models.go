package models

//Token ...
type Token struct {
	Type       TokenType
	Key        string
	Value      interface{}
	Dimensions []int
}

//Line ...
type Line struct {
	CurrentLine string
	LineIndex   int64
}

//TokenFunc ...
type TokenFunc struct {
	Type      TokenType
	Key       string
	Params    []Token
	Vars      []Token
	IsDefined bool
	Calls     []Line
}
