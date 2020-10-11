package regex

import (
	"fmt"
	"regexp"
)

//CustomRegex ...
type CustomRegex struct {
	ConstantComplex *regexp.Regexp
}

//StartsWith ...
func (r CustomRegex) StartsWith(prefix, strToTest string) (bool, error) {
	compiled, err := regexp.Compile("^" + prefix)
	if err != nil {
		return false, fmt.Errorf("[ERROR] %+v", err)
	}

	return compiled.MatchString(strToTest), nil
}

//EndsWith ...
func (r CustomRegex) EndsWith(suffix, strToTest string) (bool, error) {
	compiled, err := regexp.Compile(suffix + "$")
	if err != nil {
		return false, fmt.Errorf("[ERROR] %+v", err)
	}

	return compiled.MatchString(strToTest), nil
}

func NewRegex() (*CustomRegex, error) {
	return &CustomRegex{}, nil
}
