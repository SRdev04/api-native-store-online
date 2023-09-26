package controller

import (
	"net/http"

	"github.com/SRdev04/api-native-store-online/helper"
	"github.com/SRdev04/api-native-store-online/service"
	"github.com/SRdev04/api-native-store-online/web"

	"github.com/julienschmidt/httprouter"
)

type CtrlDetailImpl struct {
	ControllerRepo service.DetailService
}

func NewCtrlDetail(serviceDetail service.DetailService) DetailController {
	return &CtrlDetailImpl{
		ControllerRepo: serviceDetail,
	}

}

func (controller *CtrlDetailImpl) CtrlInsDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var detailRequest web.DetailRequest

	helper.DecodeRequest(r, &detailRequest)

	result := controller.ControllerRepo.DetailIns(r.Context(), detailRequest)

	respon := web.ResponsesCode{
		Code:   200,
		Status: "succes insert",
		Data:   result,
	}

	helper.EncodeWriteResponse(w, respon)
}

func (controller *CtrlDetailImpl) CtrlAllDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var requestUser web.DetailRequest

	helper.DecodeRequest(r, &requestUser)

	result := controller.ControllerRepo.GetOrders(r.Context(), requestUser.Id_User)

	respon := web.ResponsesCode{
		Code:   200,
		Status: "succes get all orders",
		Data:   result,
	}

	helper.EncodeWriteResponse(w, respon)
}
