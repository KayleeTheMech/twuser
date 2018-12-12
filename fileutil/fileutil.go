package fileutil

import (
	"bufio"
	"os"
)

func ReadLinesFromFile(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	checkError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func ListFilesInDir(directory string) (fileNames []string, err error) {
	path, err2 := os.Open(directory)
	checkError(err2)

	files, err := path.Readdir(-1)
	path.Close()

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return
}
