package lexyc

import (
	"bufio"
	"fmt"
	"go-custom-compiler/models"
	"go-custom-compiler/regex"
	"log"

	"github.com/DrN3MESiS/pprnt"
)

//LexicalAnalyzer ...
type LexicalAnalyzer struct {
	File *bufio.Scanner     //File
	R    *regex.CustomRegex //Regex Handler
	EL   *log.Logger        //Error Logger
	LL   *log.Logger        //Lex Logger
	GL   *log.Logger        //General Logger

	//TEST
	CurrentBlockType models.BlockType
	ParentBlockType  models.BlockType
	ConstantStorage  []models.Token
	VariableStorage  []models.Token
}

//NewLexicalAnalyzer ...
func NewLexicalAnalyzer(file *bufio.Scanner, ErrorLogger, LexLogger, GeneralLogger *log.Logger) (*LexicalAnalyzer, error) {
	var moduleName string = "[Lexyc][NewLexicalAnalyzer()]"

	if file == nil {
		GeneralLogger.Printf("[ERR]%+v file is not present", moduleName)
		return nil, fmt.Errorf("[ERR]%+v file is not present", moduleName)
	}
	if ErrorLogger == nil || LexLogger == nil || GeneralLogger == nil {
		GeneralLogger.Printf("[ERR]%+v Loggers are not present", moduleName)
		return nil, fmt.Errorf("[ERR]%+v Loggers are not present", moduleName)
	}
	R, err := regex.NewRegex(ErrorLogger, LexLogger, GeneralLogger)
	if err != nil {
		GeneralLogger.Printf("[ERR]%+v %+v", moduleName, err.Error())
		return nil, fmt.Errorf("[ERR]%+v %+v", moduleName, err.Error())
	}

	ErrorLogger.Printf("=============================================================")
	ErrorLogger.Printf("# Linea | # Columna | Error | Descripcion | Linea del Error")
	ErrorLogger.Printf("=============================================================")

	return &LexicalAnalyzer{
		File: file,
		R:    R,
		EL:   ErrorLogger,
		LL:   LexLogger,
		GL:   GeneralLogger,

		ParentBlockType:  models.NULLBLOCK,
		CurrentBlockType: models.NULLBLOCK,
		ConstantStorage:  []models.Token{},
		VariableStorage:  []models.Token{},
	}, nil
}

//Analyze ...
func (l *LexicalAnalyzer) Analyze(debug bool) error {
	funcName := "[Analyze()]"
	var lineIndex int64 = 1
	for l.File.Scan() {
		currentLine := l.File.Text()
		l.GL.Printf("%+v Analyzing line: %+v", funcName, lineIndex)

		if len(currentLine) == 0 {
			l.GL.Printf("%+v Skipped line: %+v; Reason: Empty", funcName, lineIndex)
			lineIndex++

			continue
		}
		var LastBlockState models.BlockType
		LastBlockState = l.CurrentBlockType
		/* Type Validation */
		if isComment, err := l.R.StartsWith("//", currentLine); err != nil {
			l.GL.Printf("%+v[APP_ERR] %+v", funcName, err.Error())
			return fmt.Errorf("%+v[APP_ERR] %+v", funcName, err.Error())
		} else {
			if isComment {
				l.GL.Printf("%+vSkipping Comment at line %+v", funcName, lineIndex)
				if debug {
					log.Printf("Skipping Comment at line %+v", lineIndex)
				}
				lineIndex++

				continue
			}
		}

		/* StartsWith */

		if l.R.RegexConstante.StartsWithConstante(currentLine, lineIndex) {
			l.CurrentBlockType = models.CONSTANTBLOCK
		}

		if l.R.RegexVariable.StartsWithVariable(currentLine, lineIndex) {
			l.CurrentBlockType = models.VARIABLEBLOCK
		}

		if l.R.RegexFuncionProto.StartsWithFuncionProto(currentLine, lineIndex) && l.ParentBlockType == models.NULLBLOCK {
			l.CurrentBlockType = models.FUNCTIONPROTOBLOCK
		}

		if l.R.RegexProcedureProto.StartsWithProcedureProto(currentLine, lineIndex) && l.ParentBlockType == models.NULLBLOCK {
			l.CurrentBlockType = models.PROCEDUREPROTOBLOCK
		}

		if l.R.RegexProcedure.StartsWithProcedure(currentLine, lineIndex) {
			if l.ParentBlockType == models.NULLBLOCK {
				l.ParentBlockType = models.PROCEDUREBLOCK
				l.CurrentBlockType = models.PROCEDUREBLOCK
			} else {
				log.Printf("[ERR] Attempted to create new procedure without finalizing the last Function or Procedure [Line: %+v]", lineIndex)
				l.GL.Printf("[ERR] Attempted to create new procedure without finalizing the last Function or Procedure [Line: %+v]", lineIndex)
				//"# Linea | # Columna | Error | Descripcion | Linea del Error"
				l.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, "N/A", "N/A", "Attempted to create new procedure without finalizing the last Function or Procedure", currentLine)
			}
		}

		if l.R.RegexFunction.StartsWithFunction(currentLine, lineIndex) {
			if l.ParentBlockType == models.NULLBLOCK {
				l.ParentBlockType = models.FUNCTIONBLOCK
				l.CurrentBlockType = models.FUNCTIONBLOCK
			} else {
				log.Printf("[ERR] Attempted to create new function without finalizing the last Function or Procedure [Line: %+v]", lineIndex)
				l.GL.Printf("[ERR] Attempted to create new function without finalizing the last Function or Procedure [Line: %+v]", lineIndex)
				//"# Linea | # Columna | Error | Descripcion | Linea del Error"
				l.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, "N/A", "N/A", "Attempted to create new function without finalizing the last Function or Procedure", currentLine)
			}
		}

		if l.R.RegexInicio.StartsWithInicio(currentLine, lineIndex) {
			if l.ParentBlockType == models.PROCEDUREBLOCK {
				l.GL.Printf("%+v Initialized a PROCEDUREBLOCK [Line: %+v]", funcName, lineIndex)
			} else if l.ParentBlockType == models.FUNCTIONBLOCK {
				l.GL.Printf("%+v Initialized a FUNCTIONBLOCK [Line: %+v]", funcName, lineIndex)
			} else {
				log.Printf("[ERR] Attempted to initialize something outside of a Block [Line: %+v]", lineIndex)
				l.GL.Printf("[ERR] Attempted to initialize something outside of a Block [Line: %+v]", lineIndex)
				//"# Linea | # Columna | Error | Descripcion | Linea del Error"
				l.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, "N/A", "N/A", "Attempted to initialize something outside of a Block", currentLine)
			}
		}

		if l.R.RegexFinFunction.StartsWithFinDeFuncion(currentLine, lineIndex) {
			if l.ParentBlockType == models.FUNCTIONBLOCK {
				l.ParentBlockType = models.NULLBLOCK
				l.GL.Printf("%+v Finished a FUNCTIONBLOCK [Line: %+v]", funcName, lineIndex)
			} else if l.ParentBlockType == models.PROCEDUREBLOCK {

				log.Printf("[ERR] Attempted to end a FUNCTIONBLOCK:Inicio with a PROCEDUREBLOCK as parent [Line: %+v]", lineIndex)
				l.GL.Printf("[ERR] Attempted to end a FUNCTIONBLOCK:Inicio with a PROCEDUREBLOCK as parent [Line: %+v]", lineIndex)
				//"# Linea | # Columna | Error | Descripcion | Linea del Error"
				l.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, "N/A", "N/A", "Attempted to end a FUNCTIONBLOCK:Inicio with a PROCEDUREBLOCK as parent", currentLine)
			} else {
				log.Printf("[ERR] Attempted to end a FUNCTIONBLOCK outside of a FUNCTIONBLOCK [Line: %+v]", lineIndex)
				l.GL.Printf("[ERR] Attempted to end a FUNCTIONBLOCK outside of a FUNCTIONBLOCK [Line: %+v]", lineIndex)
				//"# Linea | # Columna | Error | Descripcion | Linea del Error"
				l.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, "N/A", "N/A", "Attempted to end a FUNCTIONBLOCK outside of a FUNCTIONBLOCK", currentLine)
			}
		}
		if l.R.RegexFinProcedure.StartsWithFinDeProcedimiento(currentLine, lineIndex) {
			if l.ParentBlockType == models.PROCEDUREBLOCK {
				l.ParentBlockType = models.NULLBLOCK
				l.GL.Printf("%+v Finished a PROCEDUREBLOCK [Line: %+v]", funcName, lineIndex)
			} else if l.ParentBlockType == models.FUNCTIONBLOCK {

				log.Printf("[ERR] Attempted to end a PROCEDUREBLOCK:Inicio with a FUNCTIONBLOCK as parent [Line: %+v]", lineIndex)
				l.GL.Printf("[ERR] Attempted to end a PROCEDUREBLOCK:Inicio with a FUNCTIONBLOCK as parent [Line: %+v]", lineIndex)
				//"# Linea | # Columna | Error | Descripcion | Linea del Error"
				l.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, "N/A", "N/A", "Attempted to end a PROCEDUREBLOCK:Inicio with a FUNCTIONBLOCK as parent", currentLine)
			} else {
				log.Printf("[ERR] Attempted to end a PROCEDUREBLOCK outside of a PROCEDUREBLOCK [Line: %+v]", lineIndex)
				l.GL.Printf("[ERR] Attempted to end a PROCEDUREBLOCK outside of a PROCEDUREBLOCK [Line: %+v]", lineIndex)
				//"# Linea | # Columna | Error | Descripcion | Linea del Error"
				l.EL.Printf("%+v|%+v|%+v|%+v|%+v", lineIndex, "N/A", "N/A", "Attempted to end a PROCEDUREBLOCK outside of a PROCEDUREBLOCK", currentLine)
			}
		}

		//Logger
		l.RegisterBlockChange(LastBlockState, debug, funcName, lineIndex)

		/* Data Segregator */
		if l.CurrentBlockType == models.CONSTANTBLOCK {
			l.NextConstant(currentLine, lineIndex, debug)
		}

		if l.CurrentBlockType == models.VARIABLEBLOCK {
			l.NextVariable(currentLine, lineIndex, debug)
		}

		if l.CurrentBlockType == models.FUNCTIONPROTOBLOCK {
			l.NextFuncionProto(currentLine, lineIndex, debug)
		}

		if l.CurrentBlockType == models.PROCEDUREPROTOBLOCK {
			l.NextProcedureProto(currentLine, lineIndex, debug)
		}

		if l.CurrentBlockType == models.PROCEDUREBLOCK {
		}

		if l.CurrentBlockType == models.FUNCTIONBLOCK {
		}
		lineIndex++
	}

	return nil
}

//RegisterBlockChange ...
func (l *LexicalAnalyzer) RegisterBlockChange(LastBlockState models.BlockType, debug bool, funcName string, lineIndex int64) {
	if LastBlockState != l.CurrentBlockType {
		l.GL.Printf("%+vSwitched to %+v [%+v]", funcName, l.CurrentBlockType, lineIndex)
		if debug {
			log.Printf("Switched to %+v [%+v]", l.CurrentBlockType, lineIndex)
		}
	}
}

//Print ...
func (l *LexicalAnalyzer) Print() {
	log.SetFlags(0)
	if len(l.ConstantStorage) > 0 {
		log.Print("Constants: ")
		pprnt.Print(l.ConstantStorage)
		log.Print("\n")
	} else {
		log.Println("Constants: 0")
	}

	if len(l.VariableStorage) > 0 {
		log.Print("Variables: ")
		pprnt.Print(l.VariableStorage)
		log.Print("\n")
	} else {
		log.Println("Variables: 0")
	}

	log.SetFlags(log.LstdFlags)
}
