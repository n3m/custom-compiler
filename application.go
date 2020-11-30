package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go-custom-compiler/helpers"
	"go-custom-compiler/lexyc"
	"go-custom-compiler/object"
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

	if !strings.HasSuffix(path, ".up") && !strings.HasSuffix(path, ".UP") {
		panic("File is not a SOURCE CODE File")
	}

	name := path
	namewoext := strings.Split(name, ".")[0]
	name = strings.Replace(name, ".", "_", -1)

	/*Remove previous logs*/
	os.Remove(name + "_" + "error_data.err")
	os.Remove(name + "_" + "lex_data.lex")
	os.Remove(name + "_" + "test_data.test")
	os.Remove(name + "_" + "process.log")
	os.Remove(name + "_" + "tmp.up")
	log.Printf("> %+v", namewoext+".eje")
	os.Remove(namewoext + ".eje")

	/* Create Loggers */
	errLogger, errFile, err := helpers.CreateLogger(name+"_"+"error_data.err", false)
	defer errFile.Close()

	lexLogger, lexFile, err := helpers.CreateLogger(name+"_"+"lex_data.lex", false)
	defer lexFile.Close()

	testLogger, testFile, err := helpers.CreateLogger(name+"_"+"test_data.test", false)
	defer testFile.Close()

	generalLogger, logFile, err := helpers.CreateLogger(name+"_"+"process.log", true)
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

	tempFile, err := os.Open(newFile.Name())
	if err != nil {
		generalLogger.Printf("Error while opening second file! (%+v)", err.Error())
		panic(err)
	} else {
		generalLogger.Printf("File opened")
	}
	defer tempFile.Close()

	/*Analyzers*/
	reader := helpers.GetScannerFromFile(tempFile)
	generalLogger.Printf("Created Scanner from to File")

	lex, err := lexyc.NewLexicalAnalyzer(reader, errLogger, lexLogger, generalLogger, testLogger, errFile)
	if err != nil {
		generalLogger.Printf("Error while creating a new Lexical Analyzer! (%+v)", err.Error())
		panic(err)
	} else {
		generalLogger.Printf("Created Lexical Analyzer")
	}

	debugMode := true
	err = lex.Analyze(debugMode)
	if err != nil {
		generalLogger.Printf("Error while analyzing! (%+v)", err.Error())
		panic(err)
	} else {
		generalLogger.Printf("File analyzed correctly")
	}

	/*CHECK FOR ANY ERRORS */
	f, _ := os.Open(errFile.Name())
	errors, err := lineCounter(f)
	if err != nil {
		log.Printf("[CHECK FOR ANY ERRORS] > %+v", err.Error())
	}
	if (errors - 3 - lex.WarningsCount) > 0 {
		generalLogger.Printf("The source code has errors! Fix them for object code generation")
		panic("The source code has errors! Fix them for object code generation")
	}

	/*Object Code*/
	fileName := filepath.Base(path)
	ejeFile := strings.ReplaceAll(fileName, filepath.Ext(fileName), ".eje")
	objectCodeLogger, objectCodeFile, err := helpers.CreateLogger(ejeFile, false)
	defer objectCodeFile.Close()

	objectCode, err := object.NewCodeGenerator(lex, objectCodeLogger)
	if err != nil {
		generalLogger.Printf("Error while creating a new Object Code Generator! (%+v)", err.Error())
		panic(err)
	} else {
		generalLogger.Printf("Created Object Code Generator")
	}

	err = objectCode.Generate()
	if err != nil {
		generalLogger.Printf("Error while generating object code! (%+v)", err.Error())
		panic(err)
	} else {
		generalLogger.Printf("Object Code generated correctly")
	}

	generalLogger.Printf("Compiler has finished with Status [%v]", lex.Status)
	os.Remove(tempFile.Name())
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
