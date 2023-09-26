package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/SRdev04/api-native-store-online/entity"
	"github.com/SRdev04/api-native-store-online/execption"
	"github.com/SRdev04/api-native-store-online/helper"
	"github.com/SRdev04/api-native-store-online/repository"
	"github.com/SRdev04/api-native-store-online/web"

	"github.com/go-playground/validator"
)

type ProductsServImpl struct {
	ProductsRepository repository.ProductsRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewProductsService(productService repository.ProductsRepository, DB *sql.DB, validate *validator.Validate) ServiceProducts {
	return &ProductsServImpl{
		ProductsRepository: productService,
		DB:                 DB,
		Validate:           validate,
	}

}

func (service *ProductsServImpl) Create(ctx context.Context, request web.RequestCreateWeb) web.WritertWeb {

	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	product := entity.ProductsTabel{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Quantity:    request.Quantity,
	}
	product = service.ProductsRepository.Save(ctx, tx, product)

	return helper.ConvertProduct(product)

}

func (service *ProductsServImpl) Update(ctx context.Context, request web.RequestUpdateWeb) web.WritertWeb {
	err := service.Validate.Struct(request)
	helper.IfError(err)
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductsRepository.FindById(ctx, tx, request.Id)

	if err != nil {
		panic(execption.NewNotFound(err.Error()))
	}

	product = service.ProductsRepository.Update(ctx, tx, product)

	return helper.ConvertProduct(product)

}

func (service *ProductsServImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(execption.NewNotFound(err.Error()))
	}
	defer helper.CommitOrRollback(tx)
	service.ProductsRepository.Delete(ctx, tx, id)

}

func (service *ProductsServImpl) FindById(ctx context.Context, id int) web.WritertWeb {

	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)
	result, _ := service.ProductsRepository.FindById(ctx, tx, id)

	return helper.ConvertProduct(result)

}

func (service *ProductsServImpl) FindAll(ctx context.Context) []web.WritertWeb {
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)
	log.Println("connected Service")
	products := service.ProductsRepository.FindAll(ctx, tx)

	return helper.ForFindAll(products)

}
