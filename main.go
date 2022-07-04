package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	cleanTrash()
}

func cleanTrash() (status string, err error) {
	x, err := readFolder("/Volumes")
	if err != nil {
		return "", err
	}
	for _, dirs := range x {
		fmt.Println("💿 - Found Volume: " + dirs)
	}

	return "", err
}

func readFolder(dir string) (fileArray []string, err error) {

	var fileArry []string
	f, err := os.Open(dir)

	if err != nil {
		return fileArry, errors.New("error opening dir")
	}

	files, err := f.Readdir(-1)
	f.Close()

	if err != nil {
		return fileArry, errors.New("error reading dir")
	}

	for _, file := range files {
		fileArry = append(fileArry, file.Name())
	}
	return fileArry, nil
}
