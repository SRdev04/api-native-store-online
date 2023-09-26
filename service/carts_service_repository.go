package service

import (
	"context"
	"store-online-restfulapi/web"
)

type ServiceRepository interface {
	GetAllCarts(ctx context.Context, id int) []web.WriterResponse
	InputCarts(ctx context.Context, request web.RequestCarts) web.WebCarts
	EditQtyId(ctx context.Context, request web.RequestEditCarts) web.WebCarts
	FindIdCarts(ctx context.Context, id int, userId int) web.WriterResponse
	DeleteCarts(ctx context.Context, id int, userId int)
}
