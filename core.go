package main

import (
	"bytes"
	"fmt"
	"net"
	"net/mail"
	"smtpdog/destr"
	"smtpdog/postnote"
	"smtpdog/userinfo"
	"strings"
)

func dmail(origin net.Addr, from string, to string, data *[]byte) {
	domain := getdomain(to)
	for i := 0; i < sitenum; i++ {
		if domain == config.Site[i].Name {
			address, err := mail.ParseAddress(to)
			if err != nil {
				return
			}
			name := strings.Split(address.Address, "@")[0]
			email(origin, from, name, data, &config.Site[i])
			return
		}
	}
}

func email(origin net.Addr, from string, username string, data *[]byte, site *Site) {
	uname := strings.Split(username, "+")[0]
	id, res := userinfo.Getuserid(site.Name, site.Key, uname)
	msg, _ := mail.ReadMessage(bytes.NewReader(*data))
	subject := destr.Destr((msg.Header.Get("Subject")))
	var info string
	if res {
		fileid, ok := postnote.Uploadfile(site.Name, site.Key, data)
		if ok {
			fmt.Println(fileid)
		} else {
			fmt.Println("error")
		}
		cont := fmt.Sprintf("@%s \n您的邮箱 %s@%s 收到了邮件:\nFrom: %s (%s) \nSubject: **%s**\n邮件内容在附件中,请下载查看。 \n附件14日内有效，如有问题请联系管理员。\n", username, username, site.Name, from, origin.String(), subject)
		status := postnote.Postnote(id, cont, fileid, site.Name, site.Key)
		info = fmt.Sprintf("邮箱 %s@%s 收到了邮件:\nFrom: %s (%s) \nSubject: %s\n应当通知用户: %s\n 附件id: %s 状态：%v", username, site.Name, from, origin.String(), subject, id, fileid, status)

	} else {
		info = fmt.Sprintf("邮箱 %s@%s 收到了邮件:\nFrom: %s (%s) \nSubject: %s\n查无此人", username, site.Name, from, origin.String(), subject)
	}
	fmt.Println(info)
}
func getdomain(adress string) string {
	address, err := mail.ParseAddress(adress)
	if err != nil {
		return "error"
	}
	return strings.Split(address.Address, "@")[1]
}
