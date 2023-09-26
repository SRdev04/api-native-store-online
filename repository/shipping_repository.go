package repository

import (
	"context"
	"database/sql"
	"store-online-restfulapi/entity"
)

type ShippingRepository interface {
	ShipInsert(ctx context.Context, tx *sql.Tx, ship entity.Shipping) entity.Shipping
	ShipAll(ctx context.Context, tx *sql.Tx, userId int) []entity.Shipping
}
