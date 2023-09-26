package service

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/SRdev04/api-native-store-online/entity"
	"github.com/SRdev04/api-native-store-online/execption"
	"github.com/SRdev04/api-native-store-online/helper"
	"github.com/SRdev04/api-native-store-online/repository"
	"github.com/SRdev04/api-native-store-online/restcallapi"
	"github.com/SRdev04/api-native-store-online/web"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type UsersImpl struct {
	UsersRepo repository.UsersRepository
	DB        *sql.DB
	validate  *validator.Validate
}

func NewUsersService(repository repository.UsersRepository, DB *sql.DB, validate *validator.Validate) UsersService {
	return &UsersImpl{
		UsersRepo: repository,
		DB:        DB,
		validate:  validate,
	}
}

func (service *UsersImpl) Regis(ctx context.Context, request web.UsersRegis) web.UsersRespon {
	err := service.validate.Struct(request)
	helper.IfError(err)
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := helper.HashPassword(request.Password)

	helper.IfError(err)

	users := entity.Users{
		Username: request.Username,
		Email:    request.Email,
		Password: result,
	}

	succes := service.UsersRepo.Register(ctx, tx, users)

	return helper.UserConv(succes)
}
func (service *UsersImpl) Login(ctx context.Context, request web.UsersLogin) web.UsersToken {
	err := service.validate.Struct(request)
	helper.IfError(err)
	tx, err := service.DB.Begin()
	helper.IfError(err)

	defer helper.CommitOrRollback(tx)

	resultDb, err := service.UsersRepo.Login(ctx, tx, request.Email)
	helper.IfError(err)

	if request.Email == resultDb.Email {

		err := bcrypt.CompareHashAndPassword([]byte(resultDb.Password), []byte(request.Password))

		if err != nil {
			// password tidak cocok

			panic(execption.NewNotFound(err.Error()))

		}
		// disini akan restcall api yang berfungsi untuk menyimpan token yang sudah dibuat
		// saya mencoba aplikasikan yang mirip oauth2
		//jadi saya gunakan hash saja
		url := "http://localhost:9090/settoken"

		id_user := strconv.Itoa(resultDb.Id)
		accessToken := helper.HashToken(resultDb.Email)

		respon := web.UsersToken{

			Access_Token: accessToken,
		}

		errs := restcallapi.RestCallApiSet(url, id_user, accessToken)
		if errs != nil {
			panic(errs)
		}
		return respon

	} else {
		respon := web.UsersToken{}
		err := service.validate.Struct(respon)
		helper.IfError(err)

		return respon
	}

}
