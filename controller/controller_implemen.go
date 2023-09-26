package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/SRdev04/api-native-store-online/helper"
	"github.com/SRdev04/api-native-store-online/service"
	"github.com/SRdev04/api-native-store-online/web"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ControllerImpl struct {
	ControllerRepo service.ServiceRepository
}

func NewControllerCarts(serviceCarts service.ServiceRepository) ControllerCarts {
	return &ControllerImpl{
		ControllerRepo: serviceCarts,
	}
}

func (controller *ControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	requestCarts := web.RequestCarts{}
	helper.DecodeRequest(r, &requestCarts)
	log.Println("connected controller")

	response := controller.ControllerRepo.GetAllCarts(r.Context(), requestCarts.User_Id)

	webResponse := web.ResponsesCode{
		Code:   http.StatusAccepted,
		Status: "succes",
		Data:   &response,
	}
	fmt.Println("Line")

	helper.EncodeWriteResponse(w, webResponse)

}

func (controller *ControllerImpl) InsCarts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	requestCarts := web.RequestCarts{}
	helper.DecodeRequest(r, &requestCarts)

	stringId := p.ByName("productsId")
	convertId, _ := strconv.Atoi(stringId)

	if convertId == 0 && requestCarts.User_Id == 0 {
		cartsRequest := controller.ControllerRepo.InputCarts(r.Context(), requestCarts)

		cartsResponse := web.ResponsesCode{
			Code:   200,
			Status: "Succes Insert Carts",
			Data:   &cartsRequest,
		}
		helper.EncodeWriteResponse(w, cartsResponse)

	} else if requestCarts.User_Id != 0 {
		requestCarts.Id_Products = convertId
		cartsRequest := controller.ControllerRepo.InputCarts(r.Context(), requestCarts)

		cartsResponse := web.ResponsesCode{
			Code:   200,
			Status: "Succes Insert Carts",
			Data:   &cartsRequest,
		}
		helper.EncodeWriteResponse(w, cartsResponse)

	}

}
func (controller *ControllerImpl) EditQtyId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	requestCarts := web.RequestEditCarts{}
	helper.DecodeRequest(r, &requestCarts)

	stringId := p.ByName("productsId")
	convertId, err := strconv.Atoi(stringId)
	helper.IfError(err)
	requestCarts.Id_Productsc = convertId
	fmt.Println(requestCarts.User_Id)

	cartsRequest := controller.ControllerRepo.EditQtyId(r.Context(), requestCarts)

	webCode := web.ResponsesCode{
		Code:   200,
		Status: "Succes update quantity",
		Data:   &cartsRequest,
	}

	helper.EncodeWriteResponse(w, webCode)

}

func (controller *ControllerImpl) FindIdCarts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	requestCarts := web.RequestEditCarts{}
	helper.DecodeRequest(r, &requestCarts)

	stringId := p.ByName("productsId")
	convertId, err := strconv.Atoi(stringId)

	helper.IfError(err)

	respopnse := controller.ControllerRepo.FindIdCarts(r.Context(), convertId, requestCarts.User_Id)

	webCode := web.ResponsesCode{
		Code:   200,
		Status: "succes get carts id",
		Data:   &respopnse,
	}
	helper.EncodeWriteResponse(w, webCode)

}
func (controller *ControllerImpl) DeleteCarts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	requestCarts := web.RequestEditCarts{}
	helper.DecodeRequest(r, &requestCarts)
	stringId := p.ByName("productsId")
	convertId, err := strconv.Atoi(stringId)

	helper.IfError(err)

	controller.ControllerRepo.DeleteCarts(r.Context(), convertId, requestCarts.User_Id)

	webCode := web.ResponsesCode{
		Code:   200,
		Status: "succes Delete",
	}
	helper.EncodeWriteResponse(w, webCode)

}
