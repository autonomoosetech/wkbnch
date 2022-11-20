package main

import (
	"os"
	"strings"
)

func validFileExt(name string) bool {
	return strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml")
}

func filesInDir(dir string) (files []string, err error) {
	// ensure trailing slash this may cause a double slash at the end but we don't care
	dir += "/"

	fileList, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range fileList {
		name := file.Name()
		if validFileExt(name) {
			files = append(files, dir+name)
		}
	}

	return
}

func filesFromFlag(input string) (files []string, err error) {
	if validFileExt(input) {
		// treat as file
		files = append(files, input)
	} else {
		// treat as directory
		files, err = filesInDir(input)
	}

	return
}
