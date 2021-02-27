package main

import (
	log "github.com/sirupsen/logrus"
	"os"

	fs "flags-searcher"
)

func main() {
	projectPath := os.Args[2]
	apiKey := os.Args[1]
	err := fs.Run(projectPath, apiKey)
	if err != nil {
		log.Error(err)
	}
}
