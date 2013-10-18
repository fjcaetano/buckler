package main

import (
	"log"
	"os"
	"path/filepath"

	"bitbucket.org/kardianos/osext"
)

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func resourcePaths() (staticPath string, dataPath string) {
	base, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal("Could not read base dir")
	}

	staticPath = filepath.Join(base, "static")
	dataPath = filepath.Join(base, "data")
	if exists(dataPath) && exists(staticPath) {
		return
	}

	staticPath = filepath.Join(".", "static")
	dataPath = filepath.Join(".", "data")
	return
}
