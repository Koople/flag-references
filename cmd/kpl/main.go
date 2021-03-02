package main

import (
	log "github.com/sirupsen/logrus"
	"os"

	fs "flags-searcher"
)

func main() {
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		panic("api key is mandatory")
	}

	projectPath := os.Args[1]

	err := fs.Run(projectPath, apiKey)
	if err != nil {
		log.Error(err)
	}
}
