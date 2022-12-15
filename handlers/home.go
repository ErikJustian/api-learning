package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseObj struct {
	Status  int               `json:"status"` // Fungsi `json:"status"` adalah untuk mendeklarasikan bahwa Status menjadi "status" saat di-marshal jd json.
	Message string            `json:"message"`
	Val     map[string]string `json:"val"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "halo!")

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var resObj = ResponseObj{}
		resObj.Status = 1
		resObj.Message = "Request berhasil."
		resObj.Val = nil

		var result, err = json.Marshal(resObj)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
