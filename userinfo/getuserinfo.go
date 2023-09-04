package userinfo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type postdata struct {
	Username string `json:"username"`
	Host     string `json:"host"`
	Token    string `json:"i"`
}

type recdata struct {
	Id string `json:"id"`
}

func Getuserid(sitename string, key string, username string) (string, bool) {
	postdata := postdata{
		Username: username,
		Host:     sitename,
		Token:    key,
	}
	data, err := json.Marshal(postdata)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(fmt.Sprintf("https://%s/api/users/show", sitename), "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))
	var rec recdata
	err = json.Unmarshal(body, &rec)
	if err != nil {
		return "", false
	}
	if rec.Id == "" {
		return "", false
	}

	return string(rec.Id), true
}
