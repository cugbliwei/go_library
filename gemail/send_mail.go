package gemail

import (
	"log"
	"net/smtp"
	"strings"
)

//smtp服务发送邮件, mailType表示邮件格式是普通文件还是html或其他
func SendToMail(host, user, password, to, mailType, subject, body string) {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	content_type := "Content-Type: text/" + mailType + "; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	if err != nil {
		log.Println("SendMail error: ", err)
	}
}
