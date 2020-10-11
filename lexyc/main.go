package lexyc

import (
	"bufio"
	"fmt"
	"go-custom-compiler/regex"
	"log"
)

//LexicalAnalyzer ...
type LexicalAnalyzer struct {
	File *bufio.Scanner
	R    *regex.CustomRegex
}

//NewLexicalAnalyzer ...
func NewLexicalAnalyzer(file *bufio.Scanner) (*LexicalAnalyzer, error) {
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
	for l.File.Scan() {
		currentLine := l.File.Text()
		if l.R.RegexConstante.StartsWithConstante(currentLine) {
			log.Printf("Inicia con constante! > %+v", currentLine)

			continue
		}

		if l.R.RegexVariable.StartsWithVariable(currentLine) {
			log.Printf("Inicia con constante! > %+v", currentLine)

			continue
		}

	}

	return nil
}
