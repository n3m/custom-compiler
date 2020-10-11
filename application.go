package main

import (
	"go-custom-compiler/helpers"
	"go-custom-compiler/lexyc"
	"os"
)

func main() {

	/*
		Runtime Params: app.exe pathToSourceCodeFile
	*/

	if len(os.Args) < 2 {
		panic("Not enough arguments!\nUsage: 'app.exe pathToSourceCodeFile'")
	}

	path := os.Args[1]

	if err := helpers.CheckIfFileExists(path); err != nil {
		panic(err)
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := helpers.GetScannerFromFile(file)
	lex, err := lexyc.NewLexicalAnalyzer(reader)
	if err != nil {
		panic(err)
	}

	err = lex.Analyze()
	if err != nil {
		panic(err)
	}

	/*
		= Expresion Regular para Asignacion a Strings (sin variables): ^\w[^\s]*\s*:=\s*\"[^"]*\"(\s*\+\s*\"[^"]*\")*;
		= Expresion Regular para Asignacion a Enteros (sin variables): ^\w[^\s]*\s*:=\s*\d+\s*(\s*\+\s*\d+)*;
		= Expresion Regular para Asignacion a Floats (sin variables): ^\w[^\s]*\s*:=\s*\d+\.\d+\s*(\s*\+\s*\d+\.\d+)*;
		= Expresion Regular para Identificador: ^\w[^\s]*
		= Expresion Regular para Palabra Reservada (constantes): ^constantes
			= Expresion Regular para Palabra Reservada (variables): ^variables
		= Expresion Regular para Palabra Reservada (real): ^real
			= Expresion Regular para Palabra Reservada (alfabetico): ^alfabetico
		= Expresion Regular para Palabra Reservada (logico): ^logico
			= Expresion Regular para Palabra Reservada (entero): ^entero
		= Expresion Regular para Palabra Reservada (funcion): ^funcion
		= Expresion Regular para Palabra Reservada (inicio): ^inicio
		= Expresion Regular para Palabra Reservada (fin): ^fin
			= Expresion Regular para Palabra Reservada (de): ^de
		= Expresion Regular para Palabra Reservada (procedimiento): ^procedimiento
			= Expresion Regular para Palabra Reservada (regresa): ^regresa
		= Expresion Regular para Palabra Reservada (si): ^si
			= Expresion Regular para Palabra Reservada (hacer): ^hacer
		= Expresion Regular para Palabra Reservada (sino): ^sino
		= Expresion Regular para Palabra Reservada (cuando): ^cuando
		= Expresion Regular para Palabra Reservada (el): ^el
			= Expresion Regular para Palabra Reservada (valor): ^valor
			= Expresion Regular para Palabra Reservada (sea): ^sea
			= Expresion Regular para Palabra Reservada (otro): ^otro
		= Expresion Regular para Palabra Reservada (desde): ^desde
		= Expresion Regular para Palabra Reservada (hasta): ^hasta
		= Expresion Regular para Palabra Reservada (incr): ^incr
		= Expresion Regular para Palabra Reservada (decr): ^decr
		= Expresion Regular para Palabra Reservada (repetir): ^repetir
			= Expresion Regular para Palabra Reservada (que): ^que
		= Expresion Regular para Palabra Reservada (mientras): ^mientras
			= Expresion Regular para Palabra Reservada (se): ^se
		= Expresion Regular para Palabra Reservada (cumpla): ^cumpla
		= Expresion Regular para Palabra Reservada (continua): ^continua
		= Expresion Regular para Palabra Reservada (interrumpe): ^interrumpe
		= Expresion Regular para Palabra Reservada (limpia): ^limpia
		= Expresion Regular para Palabra Reservada (lee): ^lee
		= Expresion Regular para Palabra Reservada (imprime): ^imprime
		= Expresion Regular para Palabra Reservada (imprimenl): ^imprimenl
	*/

}
