package enronmailparser

import (
	"fmt"
	"parsermods/email"
)

type task struct {
	ID   int
	UserID string
}

type taskRange struct {
	Start int
	End   int
}


func worker(tasks <-chan task, taskRanges map[int]taskRange, emails []*email.Email, progress chan<- int, taskFiles map[int][]string) {
	basePath := "../data/maildir"

	for task := range tasks {
		err := extractEmailsFromUser(basePath, task.UserID, taskFiles[task.ID], taskRanges[task.ID], emails)
		if err != nil {
			fmt.Println("Error extracting emails:", err, task.UserID)
			progress <- 1
			continue
		}
		progress <- 1
	}
}