package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	endpoint := "https://google-translate1.p.rapidapi.com/language/translate/v2"

	if r.Method == "POST" {
		var q string = r.FormValue("text")

		var param = url.Values{}
		param.Set("q", q)
		param.Set("source", "en")
		param.Set("target", "id")

		var payload = bytes.NewBufferString(param.Encode())

		req, _ := http.NewRequest("POST", endpoint, payload)

		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("Accept-Encoding", "application/gzip")
		req.Header.Add("X-RapidAPI-Key", "b2dfc50505msh8094628bf21487bp19c135jsn07151c73e707")
		req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Println(string(body))
	}
}
