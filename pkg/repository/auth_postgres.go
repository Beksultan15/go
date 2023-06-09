package repository

import (
	"fmt"

	"github.com/Beksultan15/project_go/models"
	"github.com/jmoiron/sqlx"
)


type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db:db}
}

func (a *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name,last_name,username,email,password) VALUES ($1,$2,$3,$4,$5) RETURNING id", usertable)
	row := a.db.QueryRow(query,user.First_name,user.Last_name,user.First_name,user.Email,user.Password)
	
	if err:= row.Scan(&id); err!= nil {
		return 0,err
	}
	return id,nil
}

func (a *AuthPostgres) GetUser(username,password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usertable)
	err := a.db.Get(&user,query,username,password)
	return user,err
}