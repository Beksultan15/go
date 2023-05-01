package repository

import (
	"time"

	"github.com/Beksultan15/project_go/models"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PurchasingProduct struct {
	db *sqlx.DB
}

func NewPurchasingProduct(db *sqlx.DB) * PurchasingProduct{
	return &PurchasingProduct{db:db}
}

func (a *PurchasingProduct) PurchasingItems(id,userId int) (string,error){
	var success string = "success"
	var fail string = "fail"
	dsn := "host=localhost user=docker password=docker dbname=go port=5434 sslmode=disable"
	
	dba,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
	
	var sourceItems []models.Products
	if err := dba.Table("products").Find(&sourceItems,id).Error; err != nil {
		return fail,err
    }

	for _, sourceItem := range sourceItems {
        destinationItem := models.Order{
			Product_id: sourceItem.Id,
			User_id: 2,
			Product_price: sourceItem.Price,
			Product_quantity: sourceItem.Quantity,
			Order_date: time.Now(),
        }
        if err := dba.Table("order_details").Create(&destinationItem).Error; err != nil {
            return fail,err
        }
    }
	return success,nil
}