package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// Handle the error (e.g., log it or return an error response)
		return
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		// Handle the error (e.g., log it or return an error response)
		return
	}
}
