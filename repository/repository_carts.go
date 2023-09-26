package repository

import (
	"context"
	"database/sql"
	"store-online-restfulapi/entity"
)

type RepositoryCarts interface {
	GetAllCarts(ctx context.Context, tx *sql.Tx, id int) ([]entity.ResultCarts, error)
	InsertCarts(ctx context.Context, tx *sql.Tx, carts entity.ResultCarts) entity.ResultCarts
	EditQtyId(ctx context.Context, tx *sql.Tx, carts entity.ResultCarts) entity.ResultCarts
	CartsId(ctx context.Context, tx *sql.Tx, id int, user_id int) (entity.ResultCarts, error)
	CartsDelete(ctx context.Context, tx *sql.Tx, id int, userId int)
	Update1(ctx context.Context, tx *sql.Tx, id int, userId int)
}
