package mailer

import (
	"fmt"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
	"gopkg.in/gomail.v2"
)

type MailSender interface {
	SendMail(adverts []entities.Advertisement, Mail string) error
}

type mailer struct {
	Port     int
	Host     string
	From     string
	Password string
}

func NewMailer(port int, host, mail, password string) *mailer {
	return &mailer{
		Host:     host,
		Port:     port,
		Password: password,
		From:     mail,
	}
}

func (m *mailer) SendMail(adverts []entities.Advertisement, Mail string) error {
	mess := gomail.NewMessage()
	mess.SetHeader("From", m.From)
	mess.SetHeader("To", Mail)
	mess.SetHeader("Subject", "Hi! New advertisements into OLX")

	mess.SetBody("text/html", m.GenerateBody(adverts))

	d := gomail.NewDialer(m.Host, m.Port, m.From, m.Password)
	d.SSL = true

	if err := d.DialAndSend(mess); err != nil {
		return err
	}
	return nil
}

func (m *mailer) GenerateBody(adverts []entities.Advertisement) string {
	var message string
	for k, v := range adverts {
		message = message + fmt.Sprintf("%d - <b>%v</b>  <a href=\"%v\" style=\"color: #7f8fa4; text-decoration: underline;\">[click here]</a>!</br>", k, v.Title, v.URL)
	}
	return message
}
