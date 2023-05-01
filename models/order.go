package models

import (
	"time"
)

type Order struct {
  Id int `json:"id"`
  Product_id int `json:"product_id"`
  User_id int `json:"user_id"`
  Product_price uint `json:"product_price"`
  Product_quantity int `json:"product_quantity"`
  Order_date time.Time `json:"order_date"`
}