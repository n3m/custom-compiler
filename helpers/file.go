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
		return fmt.Errorf("[ERROR] %+v", err)
	}

	if info.IsDir() {
		return fmt.Errorf("[ERROR] Path is directory")
	}

	return nil
}

//GetReaderFromFile ...
func GetReaderFromFile(file *os.File) *bufio.Reader {
	return bufio.NewReader(file)
}
