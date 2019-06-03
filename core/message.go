package core

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/husk"

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

func (m Message) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}

func GetMessages(page, size int) husk.Collection {
	return ctx.Messages.Find(page, size, husk.Everything())
}

func GetMessage(key husk.Key) (*Message, error) {
	rec, err := ctx.Messages.FindByKey(key)

	if err != nil {
		return nil, err
	}

	return rec.Data().(*Message), nil
}

func (m Message) SendMessage() error {
	body, err := populatTemplate(m)

	if err != nil {
		return err
	}

	err = sendEmail(body, m.Name, m.To)

	if err != nil {
		m.Sent = false
		m.Error = err.Error()
	} else {
		m.Sent = true
	}

	set := ctx.Messages.Create(m)

	return set.Error
}

func sendEmail(body, name, to string) error {
	smtpUser := beego.AppConfig.String("smtpUsername")
	smtpPass := beego.AppConfig.String("smtpPassword")
	smtpAddress := beego.AppConfig.String("smtpAddress")
	smtpPort, _ := beego.AppConfig.Int("smtpPort")

	gm := gomail.NewMessage()
	gm.SetHeader("From", smtpUser)
	gm.SetHeader("To", to)
	gm.SetHeader("Subject", "Contact Us - "+name)
	gm.SetBody("text/html", body)

	d := gomail.NewDialer(smtpAddress, smtpPort, smtpUser, smtpPass)

	err := d.DialAndSend(gm)

	if err != nil {
		log.Println("sendMail:", err)
	}

	return err
}
