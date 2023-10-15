package enronmailparser

import (
	"bufio"
	"fmt"
	"strings"
)

func parseMessage(content string) (message string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(content))

	isParsing := false
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "X-FileName: ") {
			isParsing = true
			continue
		}

		if isParsing {
			message += line + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		if err.Error() == "bufio.Scanner: token too long" {
			fmt.Println("Message truncated due to length")
			return message, nil
		} else {
			return "", fmt.Errorf("error reading content: %v", err)
		}
	}

	return message, nil
}


func parseField(content string, fieldName string) (fieldContent string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(content))

	isParsing := false
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, fieldName+": ") {
			isParsing = true
			fieldContent = strings.TrimSpace(line[len(fieldName)+2:])
			continue
		}

		if isParsing && (line == "" || strings.Contains(line, ": ")) {
			break
		}

		if isParsing {
			fieldContent += " " + line
		}
	}

	if err := scanner.Err(); err != nil {
		if err.Error() == "bufio.Scanner: token too long" {
			fmt.Println("Field content truncated due to length")
			return fieldContent, nil
		} else {
			return "", fmt.Errorf("error reading file: %v", err)
		}
	}

	return fieldContent, nil
}