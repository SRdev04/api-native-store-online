package controller

import (
	"net/http"
	"store-online-restfulapi/helper"
	"store-online-restfulapi/service"
	"store-online-restfulapi/web"

	"github.com/julienschmidt/httprouter"
)

type UsersControlImpl struct {
	FieldServiceUser service.UsersService
}

func NewUserController(service service.UsersService) UsersController {
	return &UsersControlImpl{
		FieldServiceUser: service,
	}

}

func (control *UsersControlImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	usersRequest := web.UsersRegis{}
	helper.DecodeRequest(r, &usersRequest)

	result := control.FieldServiceUser.Regis(r.Context(), usersRequest)

	respon := web.ResponsesCode{
		Code:   200,
		Status: "Succes Register User",
		Data:   result,
	}

	helper.EncodeWriteResponse(w, respon)

}
func (control *UsersControlImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	userRequest := web.UsersLogin{}

	helper.DecodeRequest(r, &userRequest)

	result := control.FieldServiceUser.Login(r.Context(), userRequest)

	respon := web.ResponsesCode{
		Code:   200,
		Status: "Succes Login",
		Data:   result,
	}

	helper.EncodeWriteResponse(w, respon)
}
