package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DetailController interface {
	CtrlInsDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	CtrlAllDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
