package zincsearchparser

import (
	"os"
	"fmt"
	"encoding/json"
	"parsermods/email"
)

func truncateEmailMessage(email *email.Email) error {
	emailJSON, err := json.Marshal(email)
	if err != nil {
		return fmt.Errorf("error marshalling email: %v", err)
	}
	if len(emailJSON) > 64*1024 {
		email.Message = email.Message[:len(email.Message)/2]
		return truncateEmailMessage(email)
	}
	return nil
}

func WriteEmailsToIndexerFile(emails *[]*email.Email, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	indexLine := `{ "index" : { "_index" : "enron" } }` + "\n"

	for i, email := range (*emails) {
		if email == nil {
			continue
		}
		err := truncateEmailMessage(email)
		if err != nil {
			fmt.Printf("error truncating message: %v", err)
			continue
		}
		emailJSON, err := email.ToJSON()
		if err != nil {
			fmt.Printf("error converting email to JSON: %v", err)
			continue
		}
		_, err = file.WriteString(indexLine + emailJSON + "\n")
		if i%100 == 0 {
			fmt.Printf("Progress: Written %d/%d emails to file\n", i, len(*emails))
		}
		if err != nil {
			fmt.Printf("error writing email to file: %v", err)
			continue
		}
	}
	return nil
}