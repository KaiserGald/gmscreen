// Package mailer
// 22 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package mailer

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"

	"github.com/KaiserGald/gmscreen/models"
	"github.com/KaiserGald/logger"
)

var log *logger.Logger

// Mailer is an instance of the mailer
type Mailer struct {
	tpl        *template.Template
	SmtpConfig *SMTPServer
}

// Init initializes the mailer
func (m *Mailer) Init(lg *logger.Logger) error {
	m.tpl = template.Must(template.ParseGlob(os.Getenv("APPROOT") + "/services/mailer/templates/*"))
	log = lg

	m.SmtpConfig = &SMTPServer{}
	f, err := ioutil.ReadFile(os.Getenv("APPROOT") + "/services/mailer/email.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(f, m.SmtpConfig)
	if err != nil {
		return err
	}

	return nil
}

// Mail contains the information contained in the email.
type Mail struct {
	senderID string
	toIDs    []string
	subject  string
	body     string
}

// writeMail writes the email message
func (m Mail) writeMail() string {
	message := ""
	message += fmt.Sprintf("MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n")
	message += fmt.Sprintf("From: %s\r\n", m.senderID)
	if len(m.toIDs) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(m.toIDs, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", m.subject)
	message += "\r\n" + m.body

	return message
}

// UserInfo contains the information about a user
type UserInfo struct {
	Username string
	Email    string
}

// SMTPServer contains info about the email server
type SMTPServer struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	Address string `json:"address"`
	Pw      string `json:"password"`
}

// ServerName returns the url of the server
func (s SMTPServer) ServerName() string {
	return s.Host + ":" + s.Port
}

// newMail returns a new email
func (m *Mailer) newMail(sender, subject, t, body string, receivers []string) (Mail, error) {
	var mail Mail
	mail.senderID = sender
	mail.toIDs = receivers
	mail.subject = subject
	mail.body = body
	return mail, nil
}

// SendVerificationMail sends an email
func (m *Mailer) SendVerificationMail(subject, t string, destinations []string, u models.User) error {
	buf := new(bytes.Buffer)
	if err := m.tpl.ExecuteTemplate(buf, t, u); err != nil {
		return err
	}

	mail, err := m.newMail(m.SmtpConfig.Address, subject, t, buf.String(), destinations)
	if err != nil {
		log.Error.Log("Error creating email.")
		return err
	}
	err = m.email(mail)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mailer) email(mail Mail) error {
	auth := smtp.PlainAuth("", mail.senderID, m.SmtpConfig.Pw, m.SmtpConfig.Host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.SmtpConfig.Host,
	}

	c, err := smtp.Dial(m.SmtpConfig.ServerName())
	if err != nil {
		return err
	}

	c.StartTLS(tlsconfig)

	if err = c.Auth(auth); err != nil {
		return err
	}

	if err = c.Mail(mail.senderID); err != nil {
		return err
	}

	for _, k := range mail.toIDs {
		if err = c.Rcpt(k); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(mail.writeMail()))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()

	return nil
}
