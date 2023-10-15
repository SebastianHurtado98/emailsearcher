package enronmailparser

import (
	//"os"
	//"fmt"
	//"path/filepath"
	"parsermods/email"
)

func extractEmailsFromUser(basePath, userID string, userFiles []string, userTaskRange taskRange, emails []*email.Email) error {
	for i, filePath := range userFiles {
		e, err := createEmailFromEnronFile(filePath, "", userID)
		if err != nil {
			return err
		}
		emails[userTaskRange.Start+i] = e
	}
	return nil
}