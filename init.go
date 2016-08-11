package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func loadConfig() *AppConfig {
	appConfig = new(AppConfig)

	fmt.Println("Loading application config")

	fileName := "config.json"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		panic("Error parsing config file, Perhaps it doesn't exist?")
	}
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading" + fileName)
		os.Exit(0)
	}
	err = json.Unmarshal(fileData, &appConfig)
	if err != nil {
		fmt.Println("Error parsing " + fileName)
		os.Exit(0)
	}

	// Add our current Hostname
	if appConfig.ThisHostname, err = os.Hostname(); err != nil {
		fmt.Println("Failed to get hostname")
		os.Exit(0)
	}
	fmt.Println("This hostname determined to be " + appConfig.ThisHostname)
	return appConfig
}
