package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
)

func SendNewArticleEmail(toEmail, title, body string) error {
	from := os.Getenv("EMAIL_SENDER")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")

	basePath, err := os.Getwd()
	if err != nil {
		log.Println("Failed to get working directory:", err)
		return err
	}
	templateFile := filepath.Join(basePath, "app", "utils", "mail", "new-article.html")

	templateContent, err := os.ReadFile(templateFile)
	if err != nil {
		log.Println("Failed to read email template:", err)
		return err
	}

	tmpl, err := template.New("email").Parse(string(templateContent))
	if err != nil {
		log.Println("Failed to parse email template:", err)
		return err
	}

	var bodyBuffer bytes.Buffer
	err = tmpl.Execute(&bodyBuffer, struct {
		Title string
		Body  string
	}{
		Title: title,
		Body:  body,
	})
	if err != nil {
		log.Println("Failed to execute email template:", err)
		return err
	}

	msg := fmt.Sprintf("Subject: New Article Published - %s\nMIME-Version: 1.0\nContent-Type: text/html; charset=UTF-8\n\n%s",
		title, bodyBuffer.String())

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, []byte(msg))
	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	log.Printf("Email sent to %s about new article: %s\n", toEmail, title)
	return nil
}
