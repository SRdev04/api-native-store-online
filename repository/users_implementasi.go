package repository

import (
	"context"
	"database/sql"
	"store-online-restfulapi/entity"
	"store-online-restfulapi/helper"
)

type UsersRepoImpl struct {
}

func NewUsersRepository() UsersRepository {
	return &UsersRepoImpl{}
}

func (repo *UsersRepoImpl) Register(ctx context.Context, tx *sql.Tx, users entity.Users) entity.Users {
	SQL := "insert into user(username,email,password)values(?,?,?)"

	result, err := tx.ExecContext(ctx, SQL, users.Username, users.Email, users.Password)
	helper.IfError(err)
	id, err := result.LastInsertId()
	helper.IfError(err)
	users.Id = int(id)

	return users

}
func (repo *UsersRepoImpl) Login(ctx context.Context, tx *sql.Tx, users string) (entity.Users, error) {
	SQL := "select email,password from user where email=?"

	rows, err := tx.QueryContext(ctx, SQL, users)
	helper.IfError(err)
	defer rows.Close()

	user := entity.Users{}

	if rows.Next() {

		rows.Scan(&user.Email, &user.Password)

		return user, nil

	} else {

		return user, err
	}

}
