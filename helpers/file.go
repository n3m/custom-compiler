package helpers

import (
	"bufio"
	"fmt"
	"os"
)

//CheckIfFileExists ...
func CheckIfFileExists(path string) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("[ERROR] %+v", err.Error())
	}

	if info.IsDir() {
		return fmt.Errorf("[ERROR] Path is directory")
	}

	return nil
}

//GetScannerFromFile ...
func GetScannerFromFile(file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	return scanner
}
