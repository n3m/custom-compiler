package recreate

import "regexp"

//Reg ...
//(Si)(?:([^"]*"[^"]*")*[^"]*$)
type Reg struct {
	Constantes     *regexp.Regexp
	ConstantesTEST *regexp.Regexp
	Variables      *regexp.Regexp
	VariablesTEST  *regexp.Regexp
	// Real       *regexp.Regexp
	// Alfabetico *regexp.Regexp
	// Logico     *regexp.Regexp
	// Entero     *regexp.Regexp
	Funcion           *regexp.Regexp
	FuncionTEST       *regexp.Regexp
	Procedimiento     *regexp.Regexp
	ProcedimientoTEST *regexp.Regexp
	Inicio            *regexp.Regexp
	InicioTEST        *regexp.Regexp
	Fin               *regexp.Regexp
	FinTEST           *regexp.Regexp
	Regresa           *regexp.Regexp
	RegresaTEST       *regexp.Regexp
	Si                *regexp.Regexp
	SiTEST            *regexp.Regexp
	Sino              *regexp.Regexp
	SinoTEST          *regexp.Regexp
	Cuando            *regexp.Regexp
	CuandoTEST        *regexp.Regexp
	Sea               *regexp.Regexp
	SeaTEST           *regexp.Regexp
	Otro              *regexp.Regexp
	OtroTEST          *regexp.Regexp
	Repetir           *regexp.Regexp
	RepetirTEST       *regexp.Regexp
	Mientras          *regexp.Regexp
	MientrasTEST      *regexp.Regexp
	Desde             *regexp.Regexp
	DesdeTEST         *regexp.Regexp
	Continua          *regexp.Regexp
	ContinuaTEST      *regexp.Regexp
	Interrumpe        *regexp.Regexp
	InterrumpeTEST    *regexp.Regexp
	Limpia            *regexp.Regexp
	LimpiaTEST        *regexp.Regexp
	Lee               *regexp.Regexp
	LeeTEST           *regexp.Regexp
	Imprime           *regexp.Regexp
	ImprimeTEST       *regexp.Regexp
	ImprimeNl         *regexp.Regexp
	ImprimeNlTEST     *regexp.Regexp
}

//NewReg ...
func NewReg() *Reg {
	return &Reg{
		Constantes:        regexp.MustCompile(`((?i)constantes)`),
		ConstantesTEST:    regexp.MustCompile(`((?i)constantes)(?:([^"]*"[^"]*")*[^"]*$)`),
		Continua:          regexp.MustCompile(`((?i)continua)`),
		ContinuaTEST:      regexp.MustCompile(`((?i)continua)(?:([^"]*"[^"]*")*[^"]*$)`),
		Cuando:            regexp.MustCompile(`((?i)cuando)`),
		CuandoTEST:        regexp.MustCompile(`((?i)cuando)(?:([^"]*"[^"]*")*[^"]*$)`),
		Desde:             regexp.MustCompile(`((?i)desde)`),
		DesdeTEST:         regexp.MustCompile(`((?i)desde)(?:([^"]*"[^"]*")*[^"]*$)`),
		Fin:               regexp.MustCompile(`((?i)fin)`),
		FinTEST:           regexp.MustCompile(`((?i)fin)(?:([^"]*"[^"]*")*[^"]*$)`),
		Funcion:           regexp.MustCompile(`((?i)funcion)`),
		FuncionTEST:       regexp.MustCompile(`((?i)funcion)(?:([^"]*"[^"]*")*[^"]*$)`),
		Procedimiento:     regexp.MustCompile(`((?i)procedimiento)`),
		ProcedimientoTEST: regexp.MustCompile(`((?i)procedimiento)(?:([^"]*"[^"]*")*[^"]*$)`),
		Imprime:           regexp.MustCompile(`((?i)imprime(\s)*\()`),
		ImprimeTEST:       regexp.MustCompile(`((?i)imprime(\s)*\()(?:([^"]*"[^"]*")*[^"]*$)`),
		ImprimeNl:         regexp.MustCompile(`((?i)imprimenl)`),
		ImprimeNlTEST:     regexp.MustCompile(`((?i)imprimenl)(?:([^"]*"[^"]*")*[^"]*$)`),
		Inicio:            regexp.MustCompile(`((?i)inicio)`),
		InicioTEST:        regexp.MustCompile(`((?i)inicio)(?:([^"]*"[^"]*")*[^"]*$)`),
		Interrumpe:        regexp.MustCompile(`((?i)interrumpe)`),
		InterrumpeTEST:    regexp.MustCompile(`((?i)interrumpe)(?:([^"]*"[^"]*")*[^"]*$)`),
		Lee:               regexp.MustCompile(`((?i)lee(\s)*\()`),
		LeeTEST:           regexp.MustCompile(`((?i)lee(\s)*\()(?:([^"]*"[^"]*")*[^"]*$)`),
		Limpia:            regexp.MustCompile(`((?i)limpia)`),
		LimpiaTEST:        regexp.MustCompile(`((?i)limpia)(?:([^"]*"[^"]*")*[^"]*$)`),
		Mientras:          regexp.MustCompile(`((?i)Mientras)`),
		MientrasTEST:      regexp.MustCompile(`((?i)Mientras)(?:([^"]*"[^"]*")*[^"]*$)`),
		Otro:              regexp.MustCompile(`((?i)Otro)`),
		OtroTEST:          regexp.MustCompile(`((?i)Otro)(?:([^"]*"[^"]*")*[^"]*$)`),
		Regresa:           regexp.MustCompile(`((?i)Regresa)`),
		RegresaTEST:       regexp.MustCompile(`((?i)Regresa)(?:([^"]*"[^"]*")*[^"]*$)`),
		Repetir:           regexp.MustCompile(`((?i)Repetir)`),
		RepetirTEST:       regexp.MustCompile(`((?i)Repetir)(?:([^"]*"[^"]*")*[^"]*$)`),
		Sea:               regexp.MustCompile(`((?i)Sea)`),
		SeaTEST:           regexp.MustCompile(`((?i)Sea)(?:([^"]*"[^"]*")*[^"]*$)`),
		Si:                regexp.MustCompile(`((?i)Si )`),
		SiTEST:            regexp.MustCompile(`((?i)Si )(?:([^"]*"[^"]*")*[^"]*$)`),
		Sino:              regexp.MustCompile(`((?i)Sino)`),
		SinoTEST:          regexp.MustCompile(`((?i)Sino)(?:([^"]*"[^"]*")*[^"]*$)`),
		Variables:         regexp.MustCompile(`((?i)Variables)`),
		VariablesTEST:     regexp.MustCompile(`((?i)Variables)(?:([^"]*"[^"]*")*[^"]*$)`),
	}
}
