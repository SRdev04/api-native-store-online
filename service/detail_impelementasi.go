package service

import (
	"context"
	"database/sql"

	"github.com/SRdev04/api-native-store-online/entity"
	"github.com/SRdev04/api-native-store-online/helper"
	"github.com/SRdev04/api-native-store-online/repository"
	"github.com/SRdev04/api-native-store-online/web"

	"github.com/go-playground/validator"
)

type DetailServImpl struct {
	DetailRepository repository.DetailRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewDetailService(detailRepo repository.DetailRepository, DB *sql.DB, validate *validator.Validate) DetailService {
	return &DetailServImpl{
		DetailRepository: detailRepo,
		DB:               DB,
		Validate:         validate,
	}
}
func (service *DetailServImpl) DetailIns(ctx context.Context, request web.DetailRequest) web.DetailRespon {

	tx, err := service.DB.Begin()
	helper.IfError(err)

	defer helper.CommitOrRollback(tx)

	orderId := service.DetailRepository.InsertOrder(ctx, tx, request.Id_User)

	result := helper.ConvToSlice(request)

	order := entity.Orders_Detail{
		Id_Orders: orderId,
		Id_User:   request.Id_User,
		Product:   result.Product,
	}

	hasil := service.DetailRepository.InsDetail(ctx, tx, order)

	respon := helper.ConvToRes(hasil)

	return respon

}

func (service *DetailServImpl) GetOrders(ctx context.Context, request int) []web.ResponDetail {
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	rows := service.DetailRepository.GetOrders(ctx, tx, request)

	result := helper.ConvertAllSlice(rows)

	return result

}
