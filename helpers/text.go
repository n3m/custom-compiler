package helpers

import (
	"math"
	"strings"
)

const (
	// LEXINDENT ...
	LEXINDENT int = 5
)

const (
	//IDENTIFICADOR ...
	IDENTIFICADOR string = "<Ident>"
	//OPERADORARITMETICO ...
	OPERADORARITMETICO string = "<OpArit>"
	//OPERADORASIGNACION ...
	OPERADORASIGNACION string = "<OpAsig>"
	//CONSTANTEREAL ...
	CONSTANTEREAL string = "<CteReal>"
	//CONSTANTEENTERA ...
	CONSTANTEENTERA string = "<CteEnt>"
	//CONSTANTEALFABETICA ...
	CONSTANTEALFABETICA string = "<CteAlfa>"
	//CONSTANTELOGICA ...
	CONSTANTELOGICA string = "<CteLog>"
	//PALABRARESERVADA ...
	PALABRARESERVADA string = "<PalRes>"
	//DELIMITADOR ...
	DELIMITADOR string = "<Delim>"
)

// IndentString ...
func IndentString(noIndents int, elements []string) string {
	indented := ""

	for _, element := range elements[:len(elements)-1] {
		indented += element
		noIndentations := int(math.Ceil((4.0*float64(noIndents) - float64(len(element))) / 4.0))
		indented += strings.Repeat("\t", noIndentations)
	}
	indented += elements[len(elements)-1]

	return indented
}
