package repository

import (
	"context"
	"database/sql"
	"store-online-restfulapi/entity"
	"store-online-restfulapi/helper"
)

type ShipImpl struct {
}

func NewShippingRepository() ShippingRepository {
	return &ShipImpl{}

}

func (repo *ShipImpl) ShipInsert(ctx context.Context, tx *sql.Tx, ship entity.Shipping) entity.Shipping {
	SQL := "insert into shipping(kota,kecamatan,kelurahan,jalan,wa,user_id)values(?,?,?,?,?,?)"

	result, err := tx.ExecContext(ctx, SQL, ship.Kota, ship.Kecamatan, ship.Kelurahan, ship.Jalan, ship.Wa, ship.User_Id)
	helper.IfError(err)
	id, err := result.LastInsertId()
	helper.IfError(err)
	ship.Id = int(id)

	return ship

}
func (repo *ShipImpl) ShipAll(ctx context.Context, tx *sql.Tx, userId int) []entity.Shipping {
	SQL := "select user_id,kota,kecamatan,kelurahan,jalan,wa from shipping where user_id=?"

	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.IfError(err)
	defer rows.Close()
	var shipRows []entity.Shipping

	for rows.Next() {
		ship := entity.Shipping{}

		rows.Scan(&ship.User_Id, &ship.Kota, &ship.Kecamatan, &ship.Kelurahan, &ship.Jalan, &ship.Wa)
		shipRows = append(shipRows, ship)

	}
	return shipRows

}
