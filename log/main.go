package main

import (
	"io"
	"log"
	"os"
)

func loggingSettings(logFile string) {
	testLog, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, testLog)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func main() {
	loggingSettings("./log/test.log")
	_, err := os.Open("the path which not exists")
	if err != nil {
		log.Fatalln("Exit", err)
	}
}
