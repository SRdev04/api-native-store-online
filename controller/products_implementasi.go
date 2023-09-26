package controller

import (
	"log"
	"net/http"
	"store-online-restfulapi/execption"
	"store-online-restfulapi/helper"
	"store-online-restfulapi/service"
	"store-online-restfulapi/web"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ControllerImplen struct {
	Controllererpo service.ServiceProducts
}

func NewProductsController(productServic service.ServiceProducts) ControllerRepository {
	return &ControllerImplen{
		Controllererpo: productServic,
	}

}
func (controller *ControllerImplen) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("controller connected")
	responses := controller.Controllererpo.FindAll(r.Context())

	webResponse := web.ResponsesCode{
		Code:   200,
		Status: "Succes",
		Data:   &responses,
	}

	helper.EncodeWriteResponse(w, webResponse)

}

func (controller *ControllerImplen) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("connected To Create")
	productCreateRequest := web.RequestCreateWeb{}

	helper.DecodeRequest(r, &productCreateRequest)

	productResponse := controller.Controllererpo.Create(r.Context(), productCreateRequest)
	webResponse := web.ResponsesCode{
		Code:   200,
		Status: "Succes",
		Data:   productResponse,
	}
	helper.EncodeWriteResponse(w, webResponse)

}

func (controller *ControllerImplen) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	productUpdateRequest := web.RequestUpdateWeb{}
	helper.DecodeRequest(r, &productUpdateRequest)

	stringId := p.ByName("productsId")
	convertId, err := strconv.Atoi(stringId)
	helper.IfError(err)
	productUpdateRequest.Id = convertId

	productResponse := controller.Controllererpo.Update(r.Context(), productUpdateRequest)
	webResponse := web.ResponsesCode{
		Code:   200,
		Status: "Succes",
		Data:   &productResponse,
	}
	helper.EncodeWriteResponse(w, webResponse)

}

func (controller *ControllerImplen) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("connected To Delete")
	stringId := p.ByName("productsId")
	convertId, err := strconv.Atoi(stringId)
	helper.IfError(err)

	controller.Controllererpo.Delete(r.Context(), convertId)
	webResponse := web.ResponsesCode{
		Code:   200,
		Status: "Succes",
	}
	helper.EncodeWriteResponse(w, webResponse)

}

func (controller *ControllerImplen) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	stringId := p.ByName("productsId")

	convertId, err := strconv.Atoi(stringId)
	helper.IfError(err)

	responses := controller.Controllererpo.FindById(r.Context(), convertId)

	if responses.Name == "" {

		panic(execption.NewNotFound(err.Error()))
	}
	webResponse := web.ResponsesCode{
		Code:   200,
		Status: "Succes",
		Data:   &responses,
	}
	helper.EncodeWriteResponse(w, webResponse)
}
