package main

import (
	"os"

	"go-custom-compiler/helpers"
	"go-custom-compiler/lexyc"
	"go-custom-compiler/recreate"
)

var version string = "0.0.12"

func main() {

	/*
		Runtime Params: app.exe pathToSourceCodeFile
	*/

	/* Get parameters*/
	if len(os.Args) < 2 {
		panic("Not enough arguments!\nUsage: 'app.exe pathToSourceCodeFile'")
	}
	path := os.Args[1]

	/* Create Loggers */
	errLogger, errFile, err := helpers.CreateLogger("error_data.err", false)
	defer errFile.Close()

	lexLogger, lexFile, err := helpers.CreateLogger("lex_data.lex", false)
	defer lexFile.Close()

	testLogger, testFile, err := helpers.CreateLogger("test_data.test", false)
	defer testFile.Close()

	generalLogger, logFile, err := helpers.CreateLogger("process.log", true)
	defer logFile.Close()
	generalLogger.Printf("<=== Compiler has started on V%+v ===>", version)

	/* Validation V1 and Creator*/
	if err := helpers.CheckIfFileExists(path); err != nil {
		panic(err)
	} else {
		generalLogger.Printf("File exists")
	}

	file, err := os.Open(path)
	if err != nil {
		generalLogger.Printf("Error while opening file! (%+v)", err.Error())
		panic(err)
	} else {
		generalLogger.Printf("File opened")
	}
	defer file.Close()

	/* Recreate the file without inline instructions*/
	rec := recreate.NewRecreate()
	newFile, err := rec.RecreateScan(file)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	/*Analyzers*/
	reader := helpers.GetScannerFromFile(newFile)
	generalLogger.Printf("Created Scanner from to File")

	lex, err := lexyc.NewLexicalAnalyzer(reader, errLogger, lexLogger, generalLogger, testLogger)
	if err != nil {
		generalLogger.Printf("Error while creating a new Lexical Analyzer! (%+v)", err.Error())
		panic(err)
	} else {
		generalLogger.Printf("Created Lexical Analyzer")
	}

	debugMode := false
	err = lex.Analyze(debugMode)
	if err != nil {
		generalLogger.Printf("Error while analyzing! (%+v)", err.Error())
		panic(err)
	} else {
		generalLogger.Printf("File analyzed correctly")
	}

	generalLogger.Printf("Compiler has finished with Status [0]")
	os.Remove(newFile.Name())
}
