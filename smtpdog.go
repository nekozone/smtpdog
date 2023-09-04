package main

import (
	"encoding/json"
	"net"
	"os"

	"github.com/mhale/smtpd"
)

type Site struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Config struct {
	Appname  string `json:"appname"`
	Hostname string `json:"hostname"`
	Listen   string `json:"listen"`
	Site     []Site `json:"site"`
}

var config Config
var sitenum int

func init() {
	configfile, _ := os.ReadFile("config.json")
	err := json.Unmarshal(configfile, &config)
	if err != nil {
		panic(err)
	}
	sitenum = len(config.Site)
}

func rcptHandler(remoteAddr net.Addr, from string, to string) bool {
	dom := getdomain(to)
	for i := 0; i < sitenum; i++ {
		if dom == config.Site[i].Name {
			return true
		}
	}
	return false
}

func mailHandler(remoteAddr net.Addr, from string, to []string, data []byte) error {
	for i := 0; i < len(to); i++ {
		dmail(remoteAddr, from, to[i], &data)
	}
	return nil
}

func ListenAndServe(addr string, handler smtpd.Handler, rcpt smtpd.HandlerRcpt) error {
	srv := &smtpd.Server{
		Addr:        addr,
		Handler:     handler,
		HandlerRcpt: rcpt,
		Appname:     config.Appname,
		Hostname:    config.Hostname,
	}
	return srv.ListenAndServe()
}

// func rcptHandler(remoteAddr net.Addr, from string, to string) bool {
// 	domain, err := mail.ParseAddress(to)
// 	if err != nil {
// 		return false
// 	}
// 	// return domain == "mail.example.com"
// }

func main() {
	// addlist := []string{"dog@cat.moe", "dogs@cat.moe", "neko@gmail.com"}
	// kkl := getDomain(addlist)
	// for i := 0; i < len(kkl); i++ {
	// 	fmt.Println(kkl[i])
	// }
	// smtpd.ListenAndServe("127.0.0.1:25", mailHandler, "Mao", "SMTPDOG")
	ListenAndServe(config.Listen, mailHandler, rcptHandler)
}
