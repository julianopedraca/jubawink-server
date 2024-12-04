package mail

import (
	"fmt"
	"os"
	"strconv"

	gomail "gopkg.in/mail.v2"
)

func SendSignupConfirmation(email string, uuid string) error {
	message := gomail.NewMessage()

	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	password := os.Getenv("EMAIL_PASSWORD")
	if err != nil {
		return err
	}
	message.SetHeader("From", "juliano.pedraca@gmail.com")
	message.SetHeader("To", "juliano.pedraca@gmail.com")
	message.SetHeader("Subject", "Email Confirmation")

	message.SetBody("text/plain", "Thanks for trying my app, to confirm your email please click here")
	message.AddAlternative("text/html", fmt.Sprintf(
		`
			<html>
				<body>
					<p>Thanks for trying my app, to confirm your email please click <a href="%s/email/confirmation/%s">here</a>.</p>
				</body>
			</html>
		`, os.Getenv("BASE_URL"), uuid))

	dialer := gomail.NewDialer(smtpServer, smtpPort, "juliano.pedraca@gmail.com", password)
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
