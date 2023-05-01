package repository


import (
	"fmt"
	"github.com/Beksultan15/project_go/models"
	"github.com/jmoiron/sqlx"
)



type ProductPostgres struct {
	db *sqlx.DB
}


func NewProductPostgres(db *sqlx.DB) *ProductPostgres{
	return &ProductPostgres{db:db}
}

func (a *ProductPostgres) CreateProduct(product models.Products) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (product_title,price,quantity,product_description,category_id) VALUES ($1,$2,$3,$4,$5) RETURNING id",productsTable)
	row := a.db.QueryRow(query,product.Product_title,product.Price,product.Quantity,product.Product_description,product.Category_id)	
	
	if err:= row.Scan(&id); err!= nil {
		return 0,err
	}
	return id,nil
}