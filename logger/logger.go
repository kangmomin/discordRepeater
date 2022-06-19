package logger

import (
	"log"
	"os"
	"sync"
)

type logger struct {
	*log.Logger
}

var once sync.Once
var Logger = getLogger()

func getLogger() *logger {
	var logger *logger
	once.Do(func() {
		logger = setLogger()
	})
	return logger
}

func setLogger() *logger {
	return &logger{
		Logger: log.New(os.Stderr, "repeater: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
