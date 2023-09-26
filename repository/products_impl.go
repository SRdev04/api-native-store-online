package repository

import (
	"context"
	"database/sql"
	"log"
	"store-online-restfulapi/entity"
	"store-online-restfulapi/helper"
)

const (
	I = "insert into"
	S = "select * from"
)

type ProductsImpl struct {
}

func NewProductsRepository() ProductsRepository {
	return &ProductsImpl{}
}

func (repo *ProductsImpl) Save(ctx context.Context, tx *sql.Tx, product entity.ProductsTabel) entity.ProductsTabel {
	SQL := "insert into products(name,price,description,quantity)values(?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Description, product.Quantity)
	helper.IfError(err)
	id, err := result.LastInsertId()
	helper.IfError(err)
	product.Id = int(id)
	log.Println("connected Save")

	return product
}

func (repo *ProductsImpl) Update(ctx context.Context, tx *sql.Tx, product entity.ProductsTabel) entity.ProductsTabel {
	SQL := "update products set name=?,price=?,description=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Description, product.Id)
	helper.IfError(err)

	log.Println("connected Update")
	return product

}

func (repo *ProductsImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "delete from products where id=? "
	log.Println("connected Delete")
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.IfError(err)
}

func (repo *ProductsImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.ProductsTabel, error) {
	SQL := "select *from products where id=?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.IfError(err)
	defer rows.Close()
	product := entity.ProductsTabel{}

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.Quantity)
		helper.IfError(err)

		return product, nil
	} else {

		return product, err

	}

}

func (repo *ProductsImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.ProductsTabel {
	SQL := "select *from products"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.IfError(err)

	var product []entity.ProductsTabel
	defer rows.Close()
	for rows.Next() {
		products := entity.ProductsTabel{}
		err := rows.Scan(&products.Id, &products.Name, &products.Price, &products.Description, &products.Quantity)
		helper.IfError(err)
		product = append(product, products)

	}
	log.Println("connected FindAll")

	return product

}
