package main

import (
	"bytes"
	"fmt"
	"net"
	"net/mail"
	"smtpdog/check"
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
	msg, _ := mail.ReadMessage(bytes.NewReader(*data))
	subject := destr.Destr((msg.Header.Get("Subject")))
	var info string
	spfck := check.Spfck(origin, from)
	dkimck := check.Dkimck(data)
	if !(spfck && dkimck) {
		info = fmt.Sprintf("邮箱 %s@%s 收到了邮件:\nFrom: %s (%s) \nSubject: %s\nspf dkim检查未通过 不得通知用户", username, site.Name, from, origin.String(), subject)
		fmt.Println(info)
		return
	}
	uname := strings.Split(username, "+")[0]
	id, res := userinfo.Getuserid(site.Name, site.Key, uname)
	if res {
		fileid, ok := postnote.Uploadfile(site.Name, site.Key, data)
		if ok {
			fmt.Println(fileid)
		} else {
			fmt.Println("error")
		}
		var ckinfo string
		if spfck && dkimck {
			ckinfo = "邮件可能有问题，注意防范钓鱼虚假欺诈邮件。"
		} else {
			if spfck {
				ckinfo = "DKIM检查未通过，邮件风险较高，注意防范钓鱼虚假欺诈邮件。"
			} else {
				ckinfo = "SPF检查未通过，邮件风险极高，注意防范钓鱼虚假欺诈邮件。"
			}
		}

		cont := fmt.Sprintf("@%s \n您的邮箱 %s@%s 收到了邮件:\nFrom: %s (%s) \nSubject: **%s**\n邮件内容在附件中,请下载查看。 \n附件14日内有效，如有问题请联系管理员。\n%s\n", username, username, site.Name, from, origin.String(), subject, ckinfo)
		status := postnote.Postnote(id, cont, fileid, site.Name, site.Key)
		info = fmt.Sprintf("邮箱 %s@%s 收到了邮件:\nFrom: %s (%s) \nSubject: %s\n应当通知用户: %s\n附件id: %s 状态：%v dkim:%v sfp:%v", username, site.Name, from, origin.String(), subject, id, fileid, status, dkimck, spfck)

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
