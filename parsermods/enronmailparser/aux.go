package enronmailparser

import (
	"os"
	"path/filepath"
	"io"
)

func getSubdirectories(dirPath string) ([]string, error) {
	file, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	subdirs, err := file.ReadDir(-1)
	if err != nil {
		return nil, err
	}

	var dirNames []string
	for _, subdir := range subdirs {
		if subdir.IsDir() {
			dirNames = append(dirNames, subdir.Name())
		}
	}

	return dirNames, nil
}

func getFilesInDir(dirPath string, taskFiles map[int][]string, id int) error {
	if _, exists := taskFiles[id]; !exists {
		taskFiles[id] = []string{}
	}

	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, readErr := readFileContent(path, 64*1024)
			if readErr != nil {
				return readErr
			}
			taskFiles[id] = append(taskFiles[id], content)
		}
		return nil
	})
}

func readFileContent(filePath string, maxBytes int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := make([]byte, maxBytes)
	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	return string(buf[:n]), nil
}
