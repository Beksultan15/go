package models

type Comment struct {
  Id int `json:"id"`
  Product_id int `json:"product_id"`
  User_id int `json:"user_id"`
  Rating int `json:"rating"`
  Comment string `json:"comment_body"` 
}