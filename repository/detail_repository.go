package repository

import (
	"context"
	"database/sql"
	"store-online-restfulapi/entity"
)

type DetailRepository interface {
	InsDetail(ctx context.Context, tx *sql.Tx, detail entity.Orders_Detail) entity.Orders_Detail
	InsertOrder(ctx context.Context, tx *sql.Tx, userId int) int
	GetOrders(ctx context.Context, tx *sql.Tx, userId int) []entity.AllOrders
}
