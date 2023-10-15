package email

import (
	"strings"
	"time"
	"fmt"
	"encoding/json"
)

type Email struct {
	EmailDate      time.Time `json:"date"`
	From           string    `json:"from"`
	To             []string  `json:"to"`
	Subject        string    `json:"subject"`
	Cc             []string  `json:"cc"`
	Bcc            []string  `json:"bcc"`
	Message        string    `json:"message"`
	SourceCategory string    `json:"sourceCategory"`
	SourceParentID string    `json:"sourceParentId"`
}

func (e *Email) ParseDateField(dateStr string) error {
	const layout = "Mon, 2 Jan 2006 15:04:05 -0700 (MST)"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return fmt.Errorf("error parsing date: %v", err)
	}
	e.EmailDate = date
	return nil
}

func (e *Email) ParseFromField(line string) error {
	e.From = strings.TrimSpace(line)
	return nil
}

func (e *Email) ParseToField(line string) error {
	e.To = parseEmailList(line)
	return nil
}

func (e *Email) ParseCcField(line string) error {
	e.Cc = parseEmailList(line)
	return nil
}

func (e *Email) ParseBccField(line string) error {
	e.Bcc = parseEmailList(line)
	return nil
}

func (e *Email) ParseSubjectField(line string) error {
	e.Subject = strings.TrimSpace(line)
	return nil
}

func (e *Email) ParseMessageField(line string) error {
	e.Message = strings.TrimSpace(line)
	return nil
}

func (e *Email) ParseSourceCategoryField(line string) error {
	e.SourceCategory = strings.TrimSpace(line)
	return nil
}

func (e *Email) ParseSourceParentIDField(line string) error {
	e.SourceParentID = strings.TrimSpace(line)
	return nil
}

func (e *Email) Print() {
	fmt.Println("Date:", e.EmailDate)
	fmt.Println("From:", e.From)
	fmt.Println("To:", e.To)
	fmt.Println("Subject:", e.Subject)
	fmt.Println("Cc:", e.Cc)
	fmt.Println("Bcc:", e.Bcc)
	fmt.Println("Message:", e.Message)
	fmt.Println("SourceCategory:", e.SourceCategory)
	fmt.Println("SourceParentID:", e.SourceParentID)
}

func (e *Email) ToJSON() (string, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return "", fmt.Errorf("error marshalling email to JSON: %v", err)
	}
	return string(bytes), nil
}

func parseEmailList(line string) []string {
	emails := strings.Split(line, ",")
	for i, email := range emails {
		emails[i] = strings.TrimSpace(email)
	}
	return emails
}