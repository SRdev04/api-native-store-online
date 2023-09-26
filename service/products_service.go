package service

import (
	"context"
	"store-online-restfulapi/web"
)

type ServiceProducts interface {
	Create(ctx context.Context, request web.RequestCreateWeb) web.WritertWeb
	Update(ctx context.Context, request web.RequestUpdateWeb) web.WritertWeb
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.WritertWeb
	FindAll(ctx context.Context) []web.WritertWeb
}
