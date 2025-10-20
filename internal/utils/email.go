package utils

import (
	"fmt"
	"os"

	"github.com/Habeebamoo/contact-go/internal/models"
	"gopkg.in/gomail.v2"
)

func NotifyMe(contactReq *models.Contact) error {
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress("habeebfrommaildrop@gmail.com", contactReq.Subject))
	m.SetHeader("Subject", contactReq.Subject)
	m.SetHeader("To", contactReq.ReceiverEmail)

	//Email body (HTML)
	body := fmt.Sprintf(`
		<html>
			<body>
				<h1>Message from %s</h1>
				<hr>

				<p>Name: %s</p>
				<p>Email: %s</p>
				<p>Message: %s</p>
			</body>
		</html>
	`, contactReq.Subject, contactReq.SenderName, contactReq.SenderEmail, contactReq.Message)

	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 465, "habeebfrommaildrop@gmail.com", os.Getenv("PASSWORD"))
	d.SSL = true

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}