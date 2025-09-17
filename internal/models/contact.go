package models

import "fmt"

type Contact struct {
	SenderName     string  `json:"senderName"`
	SenderEmail    string  `json:"senderEmail"`
	Message        string  `json:"message"`
	ReceiverEmail  string  `json:"receiverEmail"`
	Subject        string  `json:"subject"`
}

func (c *Contact) Validate() error {
	if c.SenderName == "" {
		return fmt.Errorf("missing field: name")
	} else if c.SenderEmail == "" {
		return fmt.Errorf("missing field: email")
	} else if c.Message == "" {
		return fmt.Errorf("missing field: message")
	} else if c.ReceiverEmail == "" {
		return fmt.Errorf("missing field: receiver email")
	} else if c.Subject == "" {
		return fmt.Errorf("missing field: subject")
	}
	return nil
}
