package logger

import (
	"log"
	"os"
)

// func New(file string) (*log.Logger, func(), error) {
// 	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	cleanup := func() { _ = f.Close() }
// 	return log.New(f, "", log.LstdFlags|log.Lshortfile), cleanup, nil
// }

func NewFileLogger(path string) *log.Logger {
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return log.New(file, "", log.LstdFlags)
}
