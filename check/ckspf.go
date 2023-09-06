package check

import (
	"net"
	"net/mail"
	"strings"

	"github.com/mileusna/spf"
)

func Spfck(ip net.Addr, from string) bool {
	faddr, err := mail.ParseAddress(from)
	if err != nil {
		return false
	}
	servername := strings.Split(faddr.Address, "@")[1]
	ipstr := ip.String()
	ipadress, _, err := net.SplitHostPort(ipstr)
	if err != nil {
		return false
	}
	ipadre := net.ParseIP(ipadress)
	r := spf.CheckHost(ipadre, servername, faddr.Address, "")
	if r == "PASS" {
		return true
	} else {
		return false
	}

}
