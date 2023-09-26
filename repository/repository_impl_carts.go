package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"store-online-restfulapi/entity"
	"store-online-restfulapi/helper"
)

type CartsImplementasi struct {
}

func NewCartsRepository() RepositoryCarts {
	return &CartsImplementasi{}

}

func (repo *CartsImplementasi) GetAllCarts(ctx context.Context, tx *sql.Tx, user_id int) ([]entity.ResultCarts, error) {

	SQL := "select carts.user_id,products.name ,carts.id_products,carts.price_carts, products.description,carts.quantity from carts join products on (products.id = carts.id_products)where user_id=?"

	rows, err := tx.QueryContext(ctx, SQL, user_id)
	helper.IfError(err)

	var Cart []entity.ResultCarts
	defer rows.Close()
	for rows.Next() {
		carts := entity.ResultCarts{}

		rows.Scan(&carts.User_Id, &carts.Name, &carts.Id_Products, &carts.Price_Carts, &carts.Description, &carts.Quantity)
		Cart = append(Cart, carts)

	}
	log.Println("Connected FindAllCarts")

	return Cart, nil
}
func (repo *CartsImplementasi) InsertCarts(ctx context.Context, tx *sql.Tx, carts entity.ResultCarts) entity.ResultCarts {
	SQL := "insert into carts (user_id,id_products,price_carts,quantity)values(?,?,?,?)"

	result, err := tx.ExecContext(ctx, SQL, carts.User_Id, carts.Id_Products, carts.Price_Carts, carts.Quantity)
	helper.IfError(err)

	id, err := result.LastInsertId()
	helper.IfError(err)

	carts.Id = int(id)
	log.Println("connect carts insert")
	return carts

}
func (repo *CartsImplementasi) EditQtyId(ctx context.Context, tx *sql.Tx, carts entity.ResultCarts) entity.ResultCarts {
	SQL := "update carts set quantity=?,price_carts=?  where id_products=? and user_id=?"

	_, err := tx.ExecContext(ctx, SQL, carts.Quantity, carts.Price_Carts, carts.Id_Products, carts.User_Id)
	fmt.Println("connecteditqty")
	helper.IfError(err)
	return carts

}
func (repo *CartsImplementasi) CartsId(ctx context.Context, tx *sql.Tx, id int, userId int) (entity.ResultCarts, error) {
	SQL := "select carts.user_id, products.name ,carts.id_products,carts.price_carts, products.description,carts.quantity from carts join products on (products.id = carts.id_products)where id_products=? and user_id=?"

	rows, err := tx.QueryContext(ctx, SQL, id, userId)
	helper.IfError(err)
	defer rows.Close()
	carts := entity.ResultCarts{}

	if rows.Next() {

		err := rows.Scan(&carts.User_Id, &carts.Name, &carts.Id_Products, &carts.Price_Carts, &carts.Description, &carts.Quantity)
		helper.IfError(err)

		return carts, nil
	} else {
		return carts, errors.New("item not found")
	}

}
func (repo *CartsImplementasi) CartsDelete(ctx context.Context, tx *sql.Tx, id int, userId int) {
	SQL := "delete from carts where id_products=?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.IfError(err)
}
func (repo *CartsImplementasi) Update1(ctx context.Context, tx *sql.Tx, id int, userId int) {
	SQL := "update carts set quantity=quantity+1 where id_products=? and user_id=?"

	_, err := tx.ExecContext(ctx, SQL, id, userId)

	helper.IfError(err)

}
