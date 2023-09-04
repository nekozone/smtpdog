package postnote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Postnote(uid string, cont string, fileid string, sitename string, key string) bool {
	postdata := map[string]interface{}{
		"i":              key,
		"visibility":     "specified",
		"text":           cont,
		"visibleUserIds": []string{uid},
		"fileIds":        []string{fileid},
	}
	data, err := json.Marshal(postdata)
	if err != nil {
		return false
	}
	resp, err := http.Post(fmt.Sprintf("https://%s/api/notes/create", sitename), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return false
	}
	body, err := io.ReadAll(resp.Body)
	// fmt.Println(string(body))
	defer resp.Body.Close()
	if err != nil {
		return false
	}
	var rech map[string]interface{}
	err = json.Unmarshal(body, &rech)
	if err == nil {
		return true
	}
	return true
}
