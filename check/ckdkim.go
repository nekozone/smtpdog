package check

import (
	"bytes"

	"github.com/emersion/go-msgauth/dkim"
)

func Dkimck(data *[]byte) bool {
	r := bytes.NewReader(*data)
	verifications, err := dkim.Verify(r)
	if err != nil {
		return false
	}
	if len(verifications) == 0 {
		return false
	}
	if verifications[0].Err == nil {
		return true
	}
	return false
}
