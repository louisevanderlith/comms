package core

import (
	"github.com/louisevanderlith/husk"
	"log"

	"gopkg.in/gomail.v2"
)

type Message struct {
	Name         string `hsk:"size(50)"`
	Email        string `hsk:"size(128)"`
	Phone        string `hsk:"size(15)"`
	Body         string `hsk:"size(1024)"`
	To           string `hsk:"null;size(128)"`
	Sent         bool   `hsk:"default(false)"`
	Error        string `hsk:"null;size(2048)"`
	TemplateName string `hsk:"null;size(18)"`
}

func (m Message) Valid() error {
	return husk.ValidateStruct(&m)
}

func GetMessages(page, size int) (husk.Collection, error) {
	return ctx.Messages.Find(page, size, husk.Everything())
}

func GetMessage(key husk.Key) (husk.Dataer, error) {
	rec, err := ctx.Messages.FindByKey(key)

	if err != nil {
		return nil, err
	}

	return rec.Data(), nil
}

func (m Message) SendMessage(smtpUser, smtpPass, smtpHost string, smtpPort int) error {
	body, err := PopulatTemplate(m)

	if err != nil {
		return err
	}

	err = sendEmail(body, m.Name, m.To, smtpUser, smtpPass, smtpHost, smtpPort)

	if err != nil {
		m.Sent = false
		m.Error = err.Error()
	} else {
		m.Sent = true
	}

	set := ctx.Messages.Create(m)

	if set.Error != nil {
		return set.Error
	}

	return ctx.Messages.Save()
}

func sendEmail(body, subject, to, smtpUser, smtpPass, smtpHost string, smtpPort int) error {
	gm := gomail.NewMessage()
	gm.SetHeader("From", smtpUser)
	gm.SetHeader("To", to)
	gm.SetHeader("Subject", subject)
	gm.SetBody("text/html", body)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	err := d.DialAndSend(gm)

	if err != nil {
		log.Println("sendMail:", err)
	}

	return err
}
