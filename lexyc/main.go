package lexyc

import (
	"bufio"
	"fmt"
	"go-custom-compiler/regex"
	"log"
)

//LexicalAnalyzer ...
type LexicalAnalyzer struct {
	File *bufio.Reader
	R    *regex.CustomRegex
}

//NewLexicalAnalyzer ...
func NewLexicalAnalyzer(file *bufio.Reader) (*LexicalAnalyzer, error) {
	if file == nil {
		return nil, fmt.Errorf("[ERROR] file is not present")
	}
	R, err := regex.NewRegex()
	if err != nil {
		return nil, fmt.Errorf("[ERROR] %+v", err)
	}

	return &LexicalAnalyzer{
		File: file,
		R:    R,
	}, nil
}

//Analyze ...
func (l *LexicalAnalyzer) Analyze() error {
	for range []int{1, 2, 3} {
		str, err := l.File.ReadString('\n')
		if err != nil {
			return err
		}

		log.Printf("%+v", str)
	}
	return nil
}
