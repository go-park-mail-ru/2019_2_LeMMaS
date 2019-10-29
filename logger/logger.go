package logger

import (
	"log"
)

func initLogger() *log.Logger {
	Error := log.New(errorHandle, 
		"ERROR: ",
		log.Ldate | log.Ltime | log.Llongfile)
	return Error
}