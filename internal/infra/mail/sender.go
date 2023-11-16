package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddresGmail    = "smtp.gmail.com"
	smtpServerAddressGmail = "smtp.gmail.com:587"
)

type EmailSenderInterface interface {
	SendEmail(subject, content, to string, attachfiles []string) error
}

// envio de emails pelo gmail
type GmailSender struct {
	//nome do email que ira enviar
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAddress, fromEmailPassword string) EmailSenderInterface {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(subject, content, to string, attachfiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)

	for _, at := range attachfiles {
		_, err := e.AttachFile(at)
		if err != nil {
			return fmt.Errorf("failed to atach file %s: %w", at, err)
		}
	}

	smtAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddresGmail)
	return e.Send(smtpServerAddressGmail,smtAuth)
}
