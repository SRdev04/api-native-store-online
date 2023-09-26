package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/SRdev04/api-native-store-online/entity"
	"github.com/SRdev04/api-native-store-online/execption"
	"github.com/SRdev04/api-native-store-online/helper"
	"github.com/SRdev04/api-native-store-online/repository"
	"github.com/SRdev04/api-native-store-online/web"

	"github.com/go-playground/validator"
)

type ServiceCartsImpl struct {
	RepositoryCarts repository.RepositoryCarts
	DB              *sql.DB
	validate        *validator.Validate
}

func NewServiceCarts(repositoryCarts repository.RepositoryCarts, DB *sql.DB, validate *validator.Validate) ServiceRepository {
	return &ServiceCartsImpl{
		RepositoryCarts: repositoryCarts,
		DB:              DB,
		validate:        validate,
	}

}

func (service *ServiceCartsImpl) GetAllCarts(ctx context.Context, id int) []web.WriterResponse {
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)
	log.Println("service connect")

	hasil, err := service.RepositoryCarts.GetAllCarts(ctx, tx, id)
	if err != nil {
		panic(execption.NewNotFound(err.Error()))
	}

	result := helper.ForGetAll(hasil)

	return result
}
func (service *ServiceCartsImpl) InputCarts(ctx context.Context, request web.RequestCarts) web.WebCarts {

	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	helper.IfError(err)
	carts := entity.ResultCarts{
		User_Id:     request.User_Id,
		Id_Products: request.Id_Products,
		Price_Carts: request.Price_Carts,
		Quantity:    request.Quantity,
	}
	result, err := service.RepositoryCarts.CartsId(ctx, tx, carts.Id_Products, carts.User_Id)
	if err != nil {
		panic(execption.NewNotFound(err.Error()))
	}
	if result.Id_Products == carts.Id_Products {
		if carts.Quantity > 0 {
			carts = service.RepositoryCarts.EditQtyId(ctx, tx, carts)
		} else {
			service.RepositoryCarts.Update1(ctx, tx, carts.Id_Products, carts.User_Id)
		}
	} else {

		carts = service.RepositoryCarts.InsertCarts(ctx, tx, carts)

	}

	return helper.ConvertInsertCart(carts)

}
func (service *ServiceCartsImpl) EditQtyId(ctx context.Context, request web.RequestEditCarts) web.WebCarts {
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	carts := entity.ResultCarts{

		User_Id:     request.User_Id,
		Id_Products: request.Id_Productsc,
		Price_Carts: request.Price_Carts,
		Quantity:    request.Quantity,
	}
	carts = service.RepositoryCarts.EditQtyId(ctx, tx, carts)
	fmt.Println("connect editqty")

	return helper.ConvertInsertCart(carts)

}
func (service *ServiceCartsImpl) FindIdCarts(ctx context.Context, id int, userId int) web.WriterResponse {
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.RepositoryCarts.CartsId(ctx, tx, id, userId)
	if err != nil {
		panic(execption.NewNotFound(err.Error()))
	}
	convert := helper.ConvertTo(result)

	return convert

}
func (service *ServiceCartsImpl) DeleteCarts(ctx context.Context, id int, userId int) {
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	service.RepositoryCarts.CartsDelete(ctx, tx, id, userId)

}
