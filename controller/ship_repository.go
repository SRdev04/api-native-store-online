package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ShippingControlRepo interface {
	ShipInsert(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	ShipAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
