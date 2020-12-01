package email

import (
	"ginn/config"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendText(from, toEmail, subject,text string) (err error) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = from
	// 设置接收方的邮箱
	e.To = []string{toEmail}
	//设置主题
	e.Subject = subject
	//设置文件发送的内容
	e.Text = []byte(text)
	//设置服务器相关的配置
	err = e.Send(config.Gconfig.Email.Addr, smtp.PlainAuth("", config.Gconfig.Email.Username, config.Gconfig.Email.Password, config.Gconfig.Email.Host))
	return
}

func SendHTML(from, toEmail, subject,html string) (err error) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = from
	// 设置接收方的邮箱
	e.To = []string{toEmail}
	//设置主题
	e.Subject = subject
	//设置文件发送的内容
	e.HTML = []byte(html)
	//设置服务器相关的配置
	err = e.Send(config.Gconfig.Email.Addr, smtp.PlainAuth("", config.Gconfig.Email.Username, config.Gconfig.Email.Password, config.Gconfig.Email.Host))
	return err
}

func SendHTMLAndFile(from, toEmail, subject,html, filename string) (err error) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = from
	// 设置接收方的邮箱
	e.To = []string{toEmail}
	//设置主题
	e.Subject = subject
	//设置文件发送的内容
	e.HTML = []byte(html)
	_, err = e.AttachFile(filename)
	if nil != err{
		return
	}
	//设置服务器相关的配置
	err = e.Send(config.Gconfig.Email.Addr, smtp.PlainAuth("", config.Gconfig.Email.Username, config.Gconfig.Email.Password, config.Gconfig.Email.Host))
	return
}