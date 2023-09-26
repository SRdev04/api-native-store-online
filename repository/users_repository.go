package repository

import (
	"context"
	"database/sql"
	"store-online-restfulapi/entity"
)

type UsersRepository interface {
	Register(ctx context.Context, tx *sql.Tx, users entity.Users) entity.Users
	Login(ctx context.Context, tx *sql.Tx, users string) (entity.Users, error)
}
