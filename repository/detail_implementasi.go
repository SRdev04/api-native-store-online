package repository

import (
	"store-online-restfulapi/entity"
	"store-online-restfulapi/helper"

	"context"
	"database/sql"
)

type DetailImpl struct {
}

func NewDetailRepository() DetailRepository {
	return &DetailImpl{}

}

func (repo *DetailImpl) InsDetail(ctx context.Context, tx *sql.Tx, detail entity.Orders_Detail) entity.Orders_Detail {
	SQL := "insert into orders_detail(id_orders,id_user,id_products,price,quantity)values(?,?,?,?,?)"

	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.IfError(err)

	for _, result := range detail.Product {

		_, err := stmt.ExecContext(ctx, detail.Id_Orders, detail.Id_User, result.Id_Products, result.Price, result.Quantity)
		helper.IfError(err)
	}

	return detail
}

func (repo *DetailImpl) InsertOrder(ctx context.Context, tx *sql.Tx, userId int) int {
	SQL := "insert into orders(id_user)values(?)"

	result, err := tx.ExecContext(ctx, SQL, userId)
	helper.IfError(err)

	orderId, err := result.LastInsertId()
	helper.IfError(err)

	return int(orderId)

}

func (repo *DetailImpl) GetOrders(ctx context.Context, tx *sql.Tx, userId int) []entity.AllOrders {
	SQL := "select orders_detail.id_user,orders_detail.id_orders,products.id,products.name,products.description,orders_detail.quantity,orders_detail.price from orders_detail join products on(products.id = orders_detail.id_products)join orders on(orders_detail.id_orders =orders.id)where orders_detail.id_user=?"

	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.IfError(err)
	defer rows.Close()

	var results []entity.AllOrders

	for rows.Next() {
		result := entity.AllOrders{}

		rows.Scan(&result.Id_User, &result.Id_Orders, &result.Id_Products, &result.Name, &result.Description, &result.Quantity, &result.Price)
		results = append(results, result)

	}

	return results
}
