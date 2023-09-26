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

type ServiceShipImpl struct {
	Shippingrepo repository.ShippingRepository

	DB *sql.DB

	validate *validator.Validate
}

func NewShipService(ship repository.ShippingRepository, DB *sql.DB, validate *validator.Validate) ServiceShipRepository {
	return &ServiceShipImpl{
		Shippingrepo: ship,
		DB:           DB,
		validate:     validate,
	}

}

func (service *ServiceShipImpl) WebShipInsert(ctx context.Context, request web.RequestShipping) web.ShipingRespon {
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	ship := entity.Shipping{
		User_Id:   request.User_Id,
		Kota:      request.Kota,
		Kecamatan: request.Kecamatan,
		Kelurahan: request.Kelurahan,
		Jalan:     request.Jalan,
		Wa:        request.Wa,
	}

	ship = service.Shippingrepo.ShipInsert(ctx, tx, ship)

	return helper.ConvertShipping(ship)

}
func (service *ServiceShipImpl) GetAllShip(ctx context.Context, userId int) []web.ShipingRespon {
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	result := service.Shippingrepo.ShipAll(ctx, tx, userId)

	return helper.ConvertShipSlice(result)

}
