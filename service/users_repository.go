package service

import (
	"context"

	"github.com/SRdev04/api-native-store-online/web"
)

type UsersService interface {
	Regis(ctx context.Context, request web.UsersRegis) web.UsersRespon
	Login(ctx context.Context, request web.UsersLogin) web.UsersToken
}
