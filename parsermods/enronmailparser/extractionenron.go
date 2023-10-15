package enronmailparser

import (
	"fmt"
	"path/filepath"
	"runtime"
	"parsermods/email"
)

func LoadEnronEmailsIntoMemory(dirPath string, emails []*email.Email) {

	folderPaths, _ := getSubdirectories(dirPath)

	taskRanges := make(map[int]taskRange)
	taskFiles := make(map[int][]string)

	var start int
	for id, folderPath := range folderPaths {
		getFilesInDir(filepath.Join(dirPath, folderPath), taskFiles, id)
		taskRanges[id] = taskRange{Start: start, End: start + len(taskFiles[id]) - 1}
		start = start + len(taskFiles[id]) + 1
	}

	numWorkers := runtime.NumCPU()
	tasks := make(chan task, len(folderPaths))
	progress := make(chan int, len(folderPaths))


	for i := 0; i < numWorkers; i++ {
		go worker(tasks, taskRanges, emails, progress, taskFiles)
	}

	for id, folderPath := range folderPaths {
		tasks <- task{ID: id, UserID: folderPath}
	}

	close(tasks)

	for i := 0; i < len(folderPaths); i++ {
		<-progress
		fmt.Printf("Progress: %d/%d tasks completed\n", i+1, len(folderPaths))
	}
}