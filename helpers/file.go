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

//OpenFile ...
func OpenFile(path string) (*bufio.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] %+v", err)
	}

	return bufio.NewReader(file), nil
}
