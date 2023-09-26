package helper

import (
	"net/http"

	"github.com/SRdev04/api-native-store-online/web"
)

func Unauthorize(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	webRespon := web.ResponsesCode{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}
	EncodeWriteResponse(w, webRespon)
}
