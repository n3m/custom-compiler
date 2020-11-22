package recreate

import "regexp"

//Reg ...
type Reg struct {
	Constantes *regexp.Regexp
	Variables  *regexp.Regexp
	// Real       *regexp.Regexp
	// Alfabetico *regexp.Regexp
	// Logico     *regexp.Regexp
	// Entero     *regexp.Regexp
	Funcion       *regexp.Regexp
	Procedimiento *regexp.Regexp
	Inicio        *regexp.Regexp
	Fin           *regexp.Regexp
	Regresa       *regexp.Regexp
	Si            *regexp.Regexp
	Sino          *regexp.Regexp
	Cuando        *regexp.Regexp
	Sea           *regexp.Regexp
	Otro          *regexp.Regexp
	Repetir       *regexp.Regexp
	Mientras      *regexp.Regexp
	Desde         *regexp.Regexp
	Continua      *regexp.Regexp
	Interrumpe    *regexp.Regexp
	Limpia        *regexp.Regexp
	Lee           *regexp.Regexp
	Imprime       *regexp.Regexp
	ImprimeNl     *regexp.Regexp
}

//NewReg ...
func NewReg() *Reg {
	return &Reg{
		Constantes:    regexp.MustCompile(`((?i)constantes)`),
		Continua:      regexp.MustCompile(`((?i)continua)`),
		Cuando:        regexp.MustCompile(`((?i)cuando)`),
		Desde:         regexp.MustCompile(`((?i)desde)`),
		Fin:           regexp.MustCompile(`((?i)fin)`),
		Funcion:       regexp.MustCompile(`((?i)funcion)`),
		Procedimiento: regexp.MustCompile(`((?i)procedimiento)`),
		Imprime:       regexp.MustCompile(`((?i)imprime(\s)*\()`),
		ImprimeNl:     regexp.MustCompile(`((?i)imprimenl)`),
		Inicio:        regexp.MustCompile(`((?i)inicio)`),
		Interrumpe:    regexp.MustCompile(`((?i)interrumpe)`),
		Lee:           regexp.MustCompile(`((?i)lee(\s)*\()`),
		Limpia:        regexp.MustCompile(`((?i)limpia)`),
		Mientras:      regexp.MustCompile(`((?i)Mientras)`),
		Otro:          regexp.MustCompile(`((?i)Otro)`),
		Regresa:       regexp.MustCompile(`((?i)Regresa)`),
		Repetir:       regexp.MustCompile(`((?i)Repetir)`),
		Sea:           regexp.MustCompile(`((?i)Sea)`),
		Si:            regexp.MustCompile(`((?i)Si )`),
		Sino:          regexp.MustCompile(`((?i)Sino)`),
		Variables:     regexp.MustCompile(`((?i)Variables)`),
	}
}
