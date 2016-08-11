package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func loadAppLogger() *log.Logger {
	file, err := os.OpenFile(appConfig.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open log file")
	}
	if appConfig.Debug {
		multiLog := io.MultiWriter(file, os.Stdout)
		return log.New(multiLog, "haproxyAgent: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return log.New(file, "haproxyAgent: ", log.Ldate|log.Ltime|log.Lshortfile)
}
