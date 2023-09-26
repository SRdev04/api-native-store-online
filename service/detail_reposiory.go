package service

import (
	"context"
	"store-online-restfulapi/web"
)

type DetailService interface {
	DetailIns(ctx context.Context, request web.DetailRequest) web.DetailRespon
	GetOrders(ctx context.Context, requets int) []web.ResponDetail
}
