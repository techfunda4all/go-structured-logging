package library

import (
	"log"
)

func HelloWithLogger(log *log.Logger) {
	if log == nil {
		return
	}
	log.Println("With Logger: hello from library")
}

func HelloWithoutLogger() {
	log.Println("Without Logger: hello from library")
}
