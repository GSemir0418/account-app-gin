package email

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func newMessage(to, subject, body string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", "845217811@qq.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return m
}

func newDialer() *gomail.Dialer {
	// 从环境变量中获取数据库连接字符串
	host := os.Getenv("EMAIL_SMTP_HOST")
	portStr := os.Getenv("EMAIL_SMTP_PORT")
	user := os.Getenv("EMAIL_SMTP_USER")
	pwd := os.Getenv("EMAIL_SMTP_PASSWORD")
	if host == "" || portStr == "" || user == "" || pwd == "" {
		log.Fatal("EMAIL_SMTP_ENV is not set")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("EMAIL_SMTP_PORT is not a valid port number")
	}

	return gomail.NewDialer(host, port, user, pwd)
}

func SendValidationCode(email, code string) error {
	m := newMessage(
		email,
		fmt.Sprintf("[%s] 记账验证码", code),
		fmt.Sprintf(`
					你正在登录或注册记账网站，你的验证码是 %s 。
					<br/>
					如果你没有进行相关的操作，请直接忽略本邮件即可`, code),
	)
	d := newDialer()
	return d.DialAndSend(m)
}
