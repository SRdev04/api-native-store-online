package controller

import (
	"log"
	"net/http"
	"store-online-restfulapi/helper"
	"store-online-restfulapi/service"
	"store-online-restfulapi/web"

	"github.com/julienschmidt/httprouter"
)

type ControllerShipImpl struct {
	ControlShipping service.ServiceShipRepository
}

func NewControllerShipping(control service.ServiceShipRepository) ShippingControlRepo {
	return &ControllerShipImpl{
		ControlShipping: control,
	}

}

func (controller *ControllerShipImpl) ShipInsert(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	requestDecode := web.RequestShipping{}
	helper.DecodeRequest(r, &requestDecode)

	requestShip := controller.ControlShipping.WebShipInsert(r.Context(), requestDecode)

	respon := web.ResponsesCode{
		Code:   200,
		Status: "succes insert shipping",
		Data:   requestShip,
	}
	log.Println("connected")
	helper.EncodeWriteResponse(w, respon)
}
func (controller *ControllerShipImpl) ShipAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	requestDecode := web.RequestShipping{}
	helper.DecodeRequest(r, &requestDecode)
	requestShip := controller.ControlShipping.GetAllShip(r.Context(), requestDecode.User_Id)

	respon := web.ResponsesCode{
		Code:   200,
		Status: "succes get all shipping",
		Data:   &requestShip,
	}
	helper.EncodeWriteResponse(w, respon)
}
