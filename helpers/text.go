package helpers

import (
	"math"
	"strings"
)

const (
	// LEXINDENT ...
	LEXINDENT int = 20
)

const (
	//IDENTIFICADOR ...
	IDENTIFICADOR string = "<Ident>"
	//OPERADORARITMETICO ...
	OPERADORARITMETICO string = "<OpArit>"
	//OPERADORASIGNACION ...
	OPERADORASIGNACION string = "<OpAsig>"
	//OPERADORRELACIONAL ...
	OPERADORRELACIONAL string = "<OpRel>"
	//OPERADORLOGICO ...
	OPERADORLOGICO string = "<OpLog>"
	//OPERADOR ...
	OPERADOR string = "<Op>"
	//CONSTANTEREAL ...
	CONSTANTEREAL string = "<CteReal>"
	//CONSTANTEENTERA ...
	CONSTANTEENTERA string = "<CteEnt>"
	//CONSTANTEALFABETICA ...
	CONSTANTEALFABETICA string = "<CteAlfa>"
	//CONSTANTELOGICA ...
	CONSTANTELOGICA string = "<CteLog>"
	//CONSTANTE ...
	CONSTANTE string = "<Cte>"
	//PALABRARESERVADA ...
	PALABRARESERVADA string = "<PalRes>"
	//DELIMITADOR ...
	DELIMITADOR string = "<Delim>"
)

// IndentString returns the giving elements indented
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

// IndentStringInLines returns the giving elements grouped in lines indented
func IndentStringInLines(noIndents, noElementsInLine int, elements []string) string {
	indented := ""

	for i := 0; i < len(elements); i += noElementsInLine {
		indented += IndentString(noIndents, elements[i:i+noElementsInLine])
		indented += "\n"
	}

	return indented
}
