package helpers

import (
	"fmt"
	"log"
	"os"
)

//CreateLogger ...
func CreateLogger(filename string, flags bool) (*log.Logger, *os.File, error) {
	if filename == "" {
		return nil, nil, fmt.Errorf("[ERROR][CREATELOGGER]: Filename is empty")
	}

	f, err := os.OpenFile(filename,
		os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("[ERROR] %+v", err)
	}

	logger := log.New(f, "", log.LstdFlags)

	if !flags {
		logger.SetFlags(0)
	}

	return logger, f, nil
}
