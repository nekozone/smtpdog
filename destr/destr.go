package destr

import (
	"fmt"
	"io"
	"mime"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Destr(in string) string {
	dec := new(mime.WordDecoder)
	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "gbk":
			return transform.NewReader(input, simplifiedchinese.GBK.NewDecoder()), nil
		case "gb18030":
			return transform.NewReader(input, simplifiedchinese.GB18030.NewDecoder()), nil
		case "gb2312":
			return transform.NewReader(input, simplifiedchinese.GBK.NewDecoder()), nil
		default:
			return nil, fmt.Errorf("unhandled charset %q", charset)
		}
	}
	kk, err := dec.DecodeHeader(in)
	if err != nil {
		return in
	} else {
		return kk
	}
}
