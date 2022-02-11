package util

import (
	"os"
	"path/filepath"
)

func CreateProjectDirectory(project string) {
	path := filepath.Join(".", string(os.PathSeparator), project)
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}