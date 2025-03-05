package mail

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
)

func SendOTPEmail(toEmail string, otpCode string) error {
	from := os.Getenv("EMAIL_SENDER")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")

	// get path absolut from file
	basePath, err := os.Getwd()
	if err != nil {
		log.Println("Failed to get working directory:", err)
		return err
	}
	templateFile := filepath.Join(basePath, "app", "utils", "mail", "email-verification.html")

	// read template HTML from file
	templateContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		log.Println("Failed to read email template:", err)
		return err
	}

	// Parsing template
	tmpl, err := template.New("email").Parse(string(templateContent))
	if err != nil {
		log.Println("Failed to parse email template:", err)
		return err
	}

	// Render template with OTP
	var body bytes.Buffer
	err = tmpl.Execute(&body, struct{ OTP string }{OTP: otpCode})
	if err != nil {
		log.Println("Failed to execute email template:", err)
		return err
	}

	msg := "Subject: Your OTP Code\nMIME-Version: 1.0\nContent-Type: text/html; charset=UTF-8\n\n" + body.String()

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, []byte(msg))
	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}
	return nil
}

