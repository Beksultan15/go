package models

type Products struct {
  Id int `json:"id"`
  Product_title string `json:"product_title"`
  Category_id int  `json:"category_id"`
  Price uint `json:"price"`
  Quantity int `json:"quantity"`
  Product_description string `json:"product_description"`
}