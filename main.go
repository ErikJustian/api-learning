package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://random-user.p.rapidapi.com/getuser"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "b2dfc50505msh8094628bf21487bp19c135jsn07151c73e707")
	req.Header.Add("X-RapidAPI-Host", "random-user.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
