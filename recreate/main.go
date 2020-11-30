package recreate

import (
	"fmt"
	"go-custom-compiler/helpers"
	"os"
	"strings"
)

//Recreate ...
type Recreate struct {
}

//NewRecreate ...
func NewRecreate() *Recreate {
	return &Recreate{}
}

//RecreateScan ...
func (r *Recreate) RecreateScan(file *os.File) (*os.File, error) {
	fileExt := strings.Split(file.Name(), ".")[1]
	logger, logFile, err := helpers.CreateLogger(strings.Replace(file.Name(), ".", "_", -1)+"_tmp."+fileExt, false)
	if err != nil {
		return nil, fmt.Errorf("[Recreate()] > %+v", err.Error())
	}

	FileReader := helpers.GetScannerFromFile(file)

	for FileReader.Scan() {
		currentLine := FileReader.Text()
		if len(currentLine) == 0 {
			logger.Printf("")
			continue
		}
		if strings.HasPrefix(currentLine, "//") {
			continue
		}
		currentLine = strings.TrimSpace(currentLine)

		//Do stuff

		newStrs, err := r._ScanForReservedInitializers(currentLine)
		if err == nil {
			logger.Println(currentLine)
			continue
		}

		for _, each := range newStrs {
			logger.Println(each)
		}
	}

	return logFile, nil
}

type tokPos struct {
	text string
	pos  int
}

func (r *Recreate) _ScanForReservedInitializers(str string) ([]string, error) {
	tester := NewReg()
	counter := 0
	toLookFor := []tokPos{}
	isFin := false

	if tester.Fin.MatchString(str) && tester.FinTEST.MatchString(str) {
		match := tester.Fin.FindString(str)
		pos := tester.Fin.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
		isFin = true
	}
	if tester.Constantes.MatchString(str) && tester.ConstantesTEST.MatchString(str) {
		match := tester.Constantes.FindString(str)
		pos := tester.Constantes.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Continua.MatchString(str) && tester.ContinuaTEST.MatchString(str) {
		match := tester.Continua.FindString(str)
		pos := tester.Continua.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Cuando.MatchString(str) && tester.CuandoTEST.MatchString(str) {
		match := tester.Cuando.FindString(str)
		pos := tester.Cuando.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Desde.MatchString(str) && tester.DesdeTEST.MatchString(str) {
		match := tester.Desde.FindString(str)
		pos := tester.Desde.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Funcion.MatchString(str) && tester.FuncionTEST.MatchString(str) {
		if !isFin {
			match := tester.Funcion.FindString(str)
			pos := tester.Funcion.FindStringIndex(str)
			toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
			counter++
		}
	}
	if tester.Procedimiento.MatchString(str) && tester.ProcedimientoTEST.MatchString(str) {
		if !isFin {
			match := tester.Procedimiento.FindString(str)
			pos := tester.Procedimiento.FindStringIndex(str)
			toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
			counter++
		}
	}
	if tester.Imprime.MatchString(str) && tester.ImprimeTEST.MatchString(str) {
		match := tester.Imprime.FindString(str)
		pos := tester.Imprime.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.ImprimeNl.MatchString(str) && tester.ImprimeNlTEST.MatchString(str) {
		match := tester.ImprimeNl.FindString(str)
		pos := tester.ImprimeNl.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Inicio.MatchString(str) && tester.InicioTEST.MatchString(str) {
		match := tester.Inicio.FindString(str)
		pos := tester.Inicio.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Interrumpe.MatchString(str) && tester.InterrumpeTEST.MatchString(str) {
		match := tester.Interrumpe.FindString(str)
		pos := tester.Interrumpe.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Lee.MatchString(str) && tester.LeeTEST.MatchString(str) {
		match := tester.Lee.FindString(str)
		pos := tester.Lee.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Limpia.MatchString(str) && tester.LimpiaTEST.MatchString(str) {
		match := tester.Limpia.FindString(str)
		pos := tester.Limpia.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Mientras.MatchString(str) && tester.MientrasTEST.MatchString(str) {
		match := tester.Mientras.FindString(str)
		pos := tester.Mientras.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Otro.MatchString(str) && tester.OtroTEST.MatchString(str) {
		match := tester.Otro.FindString(str)
		pos := tester.Otro.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Regresa.MatchString(str) && tester.RegresaTEST.MatchString(str) {
		match := tester.Regresa.FindString(str)
		pos := tester.Regresa.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Repetir.MatchString(str) && tester.RepetirTEST.MatchString(str) {
		match := tester.Repetir.FindString(str)
		pos := tester.Repetir.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Sea.MatchString(str) && tester.SeaTEST.MatchString(str) {
		match := tester.Sea.FindString(str)
		pos := tester.Sea.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}

	if tester.Si.MatchString(str) && tester.SiTEST.MatchString(str) {
		match := tester.Si.FindString(str)
		pos := tester.Si.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Sino.MatchString(str) && tester.SinoTEST.MatchString(str) {
		match := tester.Sino.FindString(str)
		pos := tester.Sino.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}
	if tester.Variables.MatchString(str) && tester.VariablesTEST.MatchString(str) {
		match := tester.Variables.FindString(str)
		pos := tester.Variables.FindStringIndex(str)
		toLookFor = append(toLookFor, tokPos{text: match, pos: pos[0]})
		counter++
	}

	newStrs := []string{}
	if counter > 1 {
		// //log.Printf("==========\n\tCounter: %v\n\tLine: '%+v'\n\t%+v", counter, str, toLookFor)
		toSplit := toLookFor[0]
		for _, each := range toLookFor {
			if each.pos > toSplit.pos {
				toSplit = each
			}
		}

		strData := strings.Split(str, toSplit.text)
		ln1 := strData[0]
		ln1 = strings.TrimSpace(ln1)
		ln2 := toSplit.text + strData[1]
		ln2 = strings.TrimSpace(ln2)

		newStrs = append(newStrs, ln1)
		newStrs = append(newStrs, ln2)

		// //log.Printf("===\n\tLine: '%+v'\n\tNew1: '%+v'\n\tNew2: '%+v'", str, ln1, ln2)
		return newStrs, fmt.Errorf("Newln")
	}

	return newStrs, nil
}
