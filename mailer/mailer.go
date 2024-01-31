package mailer

import (
	"bytes"
	"embed"
	"text/template"
	"time"

	"gopkg.in/mail.v2"
)

//go:embed "templates"
var templates embed.FS

type MailerConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Sender   string
	Timeout  time.Duration
}

type Mailer struct {
	dialer *mail.Dialer
	config *MailerConfig
}

func NewMailer(config *MailerConfig) *Mailer {
	dialer := mail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	dialer.Timeout = config.Timeout
	return &Mailer{dialer: dialer, config: config}
}

func (m *Mailer) Send(to, templateName string, data interface{}) error {
	templatePath := "templates/" + templateName + ".html"
	content, err := template.New("email").ParseFS(templates, templatePath)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	if err := content.ExecuteTemplate(subject, "subject", data); err != nil {
		return err
	}

	plainTextBody := new(bytes.Buffer)
	if err := content.ExecuteTemplate(plainTextBody, "plainTextBody", data); err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)
	if err := content.ExecuteTemplate(htmlBody, "htmlBody", data); err != nil {
		return err
	}

	msg := mail.NewMessage()
	msg.SetHeader("From", m.config.Sender)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainTextBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	return m.dialer.DialAndSend(msg)
}
