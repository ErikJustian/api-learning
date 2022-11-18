package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	getSupportedLanguage()
	// getTranslate("Erik mencari ikan", "id", "en")
	// randomUserData := getRandomUser()
	// fmt.Println(randomUserData)
}

func getSupportedLanguage() {

	url := "https://google-translate1.p.rapidapi.com/language/translate/v2/languages"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", "b2dfc50505msh8094628bf21487bp19c135jsn07151c73e707")
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func getTranslate(text string, source string, target string) {
	endpoint := "https://google-translate1.p.rapidapi.com/language/translate/v2"

	var param = url.Values{}
	param.Set("q", text)
	param.Set("source", source)
	param.Set("target", target)

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

func getRandomUser() string {
	url := "https://random-user.p.rapidapi.com/getuser"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "b2dfc50505msh8094628bf21487bp19c135jsn07151c73e707")
	req.Header.Add("X-RapidAPI-Host", "random-user.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
