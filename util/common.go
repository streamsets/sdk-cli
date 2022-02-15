package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/markbates/pkger"
	"github.com/markbates/pkger/pkging"
)

func CreateProjectDirectory(root string, project string) {
	path := filepath.Join(root, string(os.PathSeparator), project)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func GetFile(path string) (pkging.File, int64) {
	f, err := pkger.Open(path)

	if err != nil {
		panic(err)
	}

	info, err := f.Stat()
	if err != nil {
		panic(err)
	}

	return f, info.Size()
}

func WriteStatus(msg string) {
	fmt.Println(msg)
}
