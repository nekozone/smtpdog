package postnote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type resd struct {
	Id string `json:"id"`
}

func Uploadfile(sitename string, key string, data *[]byte) (fileid string, errs bool) {
	timeint := time.Now().Unix()
	filename := fmt.Sprintf("%d.eml", timeint)
	params := map[string]string{
		"i":     key,
		"force": "true",
		"name":  filename,
	}
	bodybuf := &bytes.Buffer{}
	bodywrite := multipart.NewWriter(bodybuf)

	for k, v := range params {
		bodywrite.WriteField(k, v)
	}
	formfile, err := bodywrite.CreateFormFile("file", filename)
	if err != nil {
		return "", false
	}
	formfile.Write(*data)
	bodywrite.Close()
	conttype := bodywrite.FormDataContentType()
	resp, err := http.Post(fmt.Sprintf("https://%s/api/drive/files/create", sitename), conttype, bodybuf)
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var rec resd
	err = json.Unmarshal(body, &rec)
	if err != nil {
		return "", false
	}
	logfile(fmt.Sprintf("%d", timeint), rec.Id, sitename)
	fileid = rec.Id
	errs = true
	return
	// return string(body), true
}

func logfile(uploadtime string, fileid string, sitename string) {
	f, _ := os.OpenFile("files.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	f.WriteString(uploadtime + " " + fileid + " " + sitename + "\n")
	f.Close()
}
