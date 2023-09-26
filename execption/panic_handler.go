package execption

import (
	"log"
	"net/http"

	"github.com/SRdev04/api-native-store-online/helper"
	"github.com/SRdev04/api-native-store-online/web"

	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	if badRequest(w, r, err) {
		return
	}
	if notfound(w, r, err) {
		return
	}
	internalServerError(w, r, err)

}

func notfound(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		log.Println("not found")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		result := web.UsersToken{}

		webRespon := web.ResponsesCode{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}
		helper.EncodeTokenResponse(w, result.Access_Token, webRespon)

		return true
	} else {
		return false
	}

}

func badRequest(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	execption, ok := err.(validator.ValidationErrors)
	if ok {
		log.Println("bad request")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webRespon := web.ResponsesCode{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   execption.Error(),
		}
		helper.EncodeWriteResponse(w, webRespon)
		return true

	} else {
		return false
	}

}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println("internal error")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webRespon := web.ResponsesCode{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVICE ERROR",
		Data:   err,
	}
	helper.EncodeWriteResponse(w, webRespon)

}
