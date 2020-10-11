package lexyc

import (
	"bufio"
	"fmt"
)

type lexicalAnalyzer struct {
	File *bufio.Reader
}

//NewLexicalAnalyzer ...
func NewLexicalAnalyzer(file *bufio.Reader) (*lexicalAnalyzer, error) {
	if file == nil {
		return nil, fmt.Errorf("[ERROR] file is not present")
	}
	return &lexicalAnalyzer{
		File: file,
	}, nil
}
