package service

import (
	"context"
	"store-online-restfulapi/web"
)

type ServiceShipRepository interface {
	WebShipInsert(ctx context.Context, request web.RequestShipping) web.ShipingRespon
	GetAllShip(ctx context.Context, userId int) []web.ShipingRespon
}
