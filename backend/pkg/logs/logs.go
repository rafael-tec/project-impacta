package logs

import (
	"log"
	"os"
)

var Info *log.Logger
var Warning *log.Logger
var Error *log.Logger

func Init() {
	logfile, _ := os.Create("admin.log")
	Info = log.New(logfile, "[INFO] ", log.Ltime)
	Warning = log.New(logfile, "[WARNING] ", log.Ltime)
	Error = log.New(logfile, "[ERROR] ", log.Ltime)
}
