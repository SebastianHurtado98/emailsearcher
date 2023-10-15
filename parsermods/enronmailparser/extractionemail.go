package enronmailparser

import (
	"parsermods/email"
)


func createEmailFromEnronFile(content, sourceCategory, sourceParentID string) (*email.Email, error) {
	emailDate, err := parseField(content, "Date")
	if err != nil {
		return nil, err
	}

	from, err := parseField(content, "From")
	if err != nil {
		return nil, err
	}

	to, err := parseField(content, "To")
	if err != nil {
		return nil, err
	}

	cc, err := parseField(content, "Cc")
	if err != nil {
		return nil, err
	}

	bcc, err := parseField(content, "Bcc")
	if err != nil {
		return nil, err
	}

	subject, err := parseField(content, "Subject")
	if err != nil {
		return nil, err
	}

	message, err := parseMessage(content)
	if err != nil {
		return nil, err
	}

	e := new(email.Email)
	e.ParseDateField(emailDate)
	e.ParseFromField(from)
	e.ParseToField(to)
	e.ParseCcField(cc)
	e.ParseBccField(bcc)
	e.ParseSubjectField(subject)
	e.ParseMessageField(message)
	e.ParseSourceCategoryField(sourceCategory)
	e.ParseSourceParentIDField(sourceParentID)

	return e, nil
}