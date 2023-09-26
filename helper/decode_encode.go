package helper

import (
	"encoding/json"
	"net/http"
)

func DecodeRequest(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(result)
	IfError(err)
}

func EncodeWriteResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	IfError(err)
}

func EncodeTokenResponse(w http.ResponseWriter, token string, response interface{}) {
	w.Header().Add("content-type", "application/json")
	w.Header().Add("acces_token", token)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	IfError(err)
}
