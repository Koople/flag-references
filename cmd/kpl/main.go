package main

import (
	"os"

	fs "flags-searcher"
)

func main() {
	projectPath := os.Args[2]
	apiKey := os.Args[1]
	err := fs.Run(projectPath, apiKey)
	if err == 1 {
		os.Stderr.Write([]byte("Error"))
	}
}
