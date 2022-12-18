package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"learning/api/models"
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
		var gtTranslateRes models.GtTranslateRes
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		// fmt.Println(string(body))

		var unmarshalErr error = json.Unmarshal(body, &gtTranslateRes)

		if unmarshalErr != nil {
			http.Error(w, unmarshalErr.Error(), http.StatusInternalServerError)
			return
		}

		// Create our own response object to return.
		var response models.TranslateRes

		if len(gtTranslateRes.Data.Translations) > 0 {
			for i := 0; i < len(gtTranslateRes.Data.Translations); i++ {
				var translatedTxt string = gtTranslateRes.Data.Translations[i].TranslatedText

				if translatedTxt != "" {
					response.Status = 1
					response.Message = "Text translated successfully."
					response.Val = translatedTxt
				} else {
					response.Status = 0
					response.Message = "No word to be translated."
					response.Val = ""
				}
			}
		} else {
			response.Status = 0
			response.Message = "No results from Google Translate API."
			response.Val = ""
		}

		returnValue, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// fmt.Println(res)
		// fmt.Println(string(body))
		// respone, err := json.Marshal(string(body))
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// translation := body[];

		w.Header().Set("Content-Type", "application/json")
		w.Write(returnValue)
	}
}
