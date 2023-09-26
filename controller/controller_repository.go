package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ControllerCarts interface {
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	InsCarts(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	EditQtyId(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindIdCarts(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeleteCarts(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
