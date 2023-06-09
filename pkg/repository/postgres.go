package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usertable = "users"
	productsTable = "products"
	comment_table = "comments"
	order_detail = "order_details"
)

type Config struct {
	Host string 
	Port string 
	Username string 
	Password string 
    DBname string 
    SSLMode string 
}


func NewPostgresConfig(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",cfg.Host,cfg.Port,cfg.Username,cfg.Password,cfg.DBname,cfg.SSLMode))

	if err!= nil {
        return nil, err
    }
	err = db.Ping()
	if err!= nil {
        return nil, err
    }

	return db, nil
}