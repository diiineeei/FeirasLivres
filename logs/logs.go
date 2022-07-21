package logs

import (
	"log"
	"os"
)

const LogFile = "/tmp/app.log"

func Print(v ...any) {
	logFile, err := os.OpenFile(LogFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println(v...)
}
