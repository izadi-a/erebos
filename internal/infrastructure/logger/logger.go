package logger

import (
	"log"
	"os"
)

func New(file string) (*log.Logger, func(), error) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() { _ = f.Close() }
	return log.New(f, "", log.LstdFlags|log.Lshortfile), cleanup, nil
}
