package gmail

import (
	"gin-starter/internal/config"
	"gopkg.in/gomail.v2"
)

type Options struct {
	MailTo  []string // 收件人 多个用,分割
	Subject string   // 邮件主题
	Body    string   // 邮件内容
}

func Send(o *Options) error {
	m := gomail.NewMessage()
	mailConfig := config.Conf.Mail
	//设置发件人
	m.SetHeader("From", mailConfig.Username)
	//设置多个收件人
	m.SetHeader("To", o.MailTo...)
	//设置邮件主题
	m.SetHeader("Subject", o.Subject)
	//设置邮件正文（HTML 格式）
	m.SetBody("text/html", o.Body)
	// 添加附件
	//m.Attach("/path/to/attachment.txt")
	// 配置 SMTP 客户端
	d := gomail.NewDialer(mailConfig.Host, mailConfig.Port, mailConfig.Username, mailConfig.Password)
	// 发送邮件
	return d.DialAndSend(m)
}

const (
	Style = `<style>
                body {
                    font-family: Arial, sans-serif;
                    background-color: #f9f9f9;
                    color: #333;
                    margin: 0;
                    padding: 0;
                }
                .container {
                    max-width: 600px;
                    margin: 20px auto;
                    padding: 20px;
                    background-color: #fff;
                    border-radius: 8px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                }
                h1 {
                    color: #d9534f;
                    font-size: 24px;
                    margin-bottom: 20px;
                }
                p {
                    font-size: 16px;
                    line-height: 1.6;
                    margin: 10px 0;
                }
                .highlight {
                    color: #d9534f;
                    font-weight: bold;
                }
                .footer {
                    margin-top: 20px;
                    font-size: 14px;
                    color: #777;
                    text-align: center;
                }
</style>`
)
