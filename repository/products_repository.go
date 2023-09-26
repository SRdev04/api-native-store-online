package repository

import (
	"context"
	"database/sql"
	"store-online-restfulapi/entity"
)

type ProductsRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product entity.ProductsTabel) entity.ProductsTabel
	Update(ctx context.Context, tx *sql.Tx, product entity.ProductsTabel) entity.ProductsTabel
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindById(ctx context.Context, tx *sql.Tx, id int) (entity.ProductsTabel, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.ProductsTabel
}
