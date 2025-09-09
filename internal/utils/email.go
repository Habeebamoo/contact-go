package utils

import (
	"fmt"
	"os"

	"github.com/Habeebamoo/contact-go/internal/models"
	"gopkg.in/gomail.v2"
)

func NotifyMe(contactReq models.Contact) error {
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress("habeebfrommaildrop@gmail.com", "Porfolio"))
	m.SetHeader("Subject", "Contact Request")
	m.SetHeader("To", "habeebamoo08@gmail.com")

	//Email body (HTML)
	body := fmt.Sprintf(`
		<html>
			<body>
				<h1>Porfolio Contact</h1>
				<p>New contact request from your portfolio</p>
				<hr>

				<p>Name: %s</p>
				<p>Email: %s</p>
				<p>Message: %s</p>
			</body>
		</html>
	`, contactReq.Name, contactReq.Email, contactReq.Message)

	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 464, "habeebfrommaildrop@gmail.com", os.Getenv("PASSWORD"))
	d.SSL = true

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending message")
	}

	return nil
}